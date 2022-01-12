package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/robfig/cron/v3"
)

type DeviceRep struct {
	Time        int64
	DeviceNum   string
	NodeCode    string
	Reason      string
	CorpCode    string
	ModelCode   string
	FirstRepair int64
}

type Resp struct {
	Code int    `json:"errorCode"`
	Msg  string `json:"errorMsg"`
	Data *Data  `json:"data,omitempty"`
}

type Data struct {
	Status int `json:"status"`
}

type urlData struct {
	DeviceNum string `json:"deviceNum"`
	UrlType   int    `json:"urlType"`
	PlayType  int    `json:"playType"`
	Force     string `json:"force"`
	Type      int    `json:"type"`
	NodeCode  string `json:"nodeCode"`
	SipFlag   string `json:"sipFlag"`
}

type urlResponse struct {
	ErrorCode string `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}

type logPostResponse struct {
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
	Data      string `json:"data"`
}

type Tags struct {
	Area       string `json:"area"`
	ReturnCode string `json:"returncode"`
	Combo      string `json:"combo"`
	Target     string `json:"target"`
}

type TsdbItem struct {
	Metric    string `json:"metric"`
	Tags      Tags   `json:"tags"`
	Value     int    `json:"value"`
	Timestamp int64  `json:"timestamp"`
}

type RepairLog struct {
	Time      int64  `json:"timestamp"`
	DeviceNum string `json:"deviceNum"`
	NodeCode  string `json:"nodeCode"`
	CorpCode  string `json:"corp"`
	ModelCode string `json:"model"`
	Operator  string `json:"dataItem"`
	Reason    string `json:"type"`
	ErrorCode string `json:"resultCode"`
	ErrorMsg  string `json:"resultMsg"`
}

type RtspData struct {
	RtspUri string `json:"rtspUri"`
}

type rtspGetResponse struct {
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
	Data      bool   `json:"data"`
}

//var redisIpMap map[string]string
var reasonMap map[string]string
var repairFreq map[string]int
var repairLastTime map[string]int64
var wg sync.WaitGroup
var start = 1

var nodeRepair = flag.String("nodeRepair", "", "node code need repair")
var urlIp = flag.String("urlIp", "", "url ip")
var urlPort = flag.String("urlPort", "", "url port")
var fileUrl = flag.String("fileUrl", "", "file address")
var opentsdb = flag.String("opentsdb", "", "opentsdb address")
var redisHost = flag.String("redis_host", "", "redis host address")
var sipRedis = flag.String("sip_redis", "", "sip gateway redis host.")
var sipTable = flag.Int("sip_table", 1, "sip redis table")
var startPort = flag.String("startPort", "", "repair start port")

func LogAppend(rep DeviceRep, errorCode string, errMsg string, repairLogs *[]RepairLog) {

	repairLog := RepairLog{
		Time:      time.Now().Unix(),
		DeviceNum: rep.DeviceNum,
		NodeCode:  rep.NodeCode,
		CorpCode:  rep.CorpCode,
		ModelCode: rep.ModelCode,
		Operator:  "auto-repair",
		Reason:    rep.Reason,
		ErrorCode: errorCode,
		ErrorMsg:  errMsg,
	}
	*repairLogs = append(*repairLogs, repairLog)
}

func httpOpentsdbDo(metric string, value int, returnCode string, combo string, target string) {

	tag := Tags{
		Area:       *nodeRepair,
		Target:     target,
		Combo:      combo,
		ReturnCode: returnCode,
	}

	data := TsdbItem{
		Metric:    metric,
		Value:     value,
		Timestamp: time.Now().Unix(),
		Tags:      tag,
	}

	client := &http.Client{Timeout: 20 * time.Second}

	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", *opentsdb+"/api/put?details", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("can not connect to opentsdb.")
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("can not connect to opentsdb.")
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("can not connect to opentsdb.")
	}

	log.Println(string(body) + metric)

}

func Read(expNode string, deviceReps []DeviceRep) []DeviceRep {
	today := time.Now().Format("2006-01-02")

	f, err := os.Open(*fileUrl + "device-repair-" + today + ".csv")

	if err != nil {
		fmt.Println(err)
		return deviceReps
	}
	defer f.Close()
	w := csv.NewReader(f)
	nowOneMin := time.Now().Truncate(time.Minute).Unix()

	for {
		row, err := w.Read()
		if err != nil && err != io.EOF {
			log.Printf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		if row[0] == "Time" {
			continue
		}

		row[0] = strings.Replace(row[0], "\uFEFF", "", -1)
		timeTamp, e := strconv.ParseInt(row[0], 10, 64)
		if e != nil {
			continue
			//log.Println(e)
		}
		if len(deviceReps) > 500 {
			log.Println("Exceeded the maximum number of repaired devices.")
			break
		}

		if timeTamp > nowOneMin && timeTamp < time.Now().Unix() && strings.Contains(expNode, row[2]) {

			v, ok := repairFreq[row[1]]
			if ok {
				lastRepair, _ := repairLastTime[row[1]]
				//log.Printf("device:%s repair times: %d, last repair time: %d",row[1], v, lastRepair)
				switch v {
				case 1:
					if time.Now().Unix()-lastRepair < 180 {
						continue
					} else {
						repairFreq[row[1]]++
						repairLastTime[row[1]] = time.Now().Unix()
					}
					break
				case 2:
					if time.Now().Unix()-lastRepair < 120 {
						continue
					} else {
						repairFreq[row[1]]++
						repairLastTime[row[1]] = time.Now().Unix()
					}
					break
				case 3:
					if time.Now().Unix()-lastRepair < 1560 {
						continue
					} else {
						repairFreq[row[1]] = 1
						repairLastTime[row[1]] = time.Now().Unix()
					}
					break
				default:
					break
				}
			} else {
				repairFreq[row[1]] = 1
				repairLastTime[row[1]] = time.Now().Unix()
			}

			devRep := DeviceRep{
				Time:      timeTamp,
				DeviceNum: row[1],
				NodeCode:  row[2],
				Reason:    row[3],
				CorpCode:  row[4],
				ModelCode: row[5],
			}

			deviceReps = append(deviceReps, devRep)
		}

	}

	log.Printf("节点：%s,需修复的总设备数：%d\n", *nodeRepair, len(deviceReps))
	return deviceReps
}

func Repair(rep DeviceRep, repairLogs *[]RepairLog) {

	client := &http.Client{Timeout: 20 * time.Second}

	conn, err := redis.Dial("tcp", *redisHost,
		redis.DialDatabase(9),
		redis.DialPassword("st8evLW3"))
	if err != nil {
		log.Printf("设备：%s,节点：%s,%s修复,同步网关在线状态,失败,原因：%s",
			rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], err.Error())
		return
	}

	defer conn.Close()
	defer wg.Done()

	//判断"HASH#MEDIA#SERVER"是否存在
	isKeyExit, err := redis.Bool(conn.Do("EXISTS", "HASH#MEDIA#SERVER"))
	if !isKeyExit {
		mediaOpenurl := "http://122.229.8.245:6018/video/cnt/node/syn/data"
		mediaData := urlData{
			NodeCode: rep.NodeCode,
		}
		b, _ := json.Marshal(mediaData)
		req, _ := http.NewRequest("POST", mediaOpenurl, bytes.NewBuffer(b))
		req.Header.Set("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			log.Printf("repair media server fail: %s", err.Error())
			return
		}

		if res.StatusCode != 200 {
			LogAppend(rep, strconv.Itoa(res.StatusCode), res.Status, repairLogs)
			log.Printf("设备：%s,节点：%s,%s修复,修复HASH#MEDIA#SERVER,url post fail,%s",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], res.Status)
			return
		}

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		urlRes := urlResponse{}
		jsonErr := json.Unmarshal(body, &urlRes)
		if jsonErr != nil {
			log.Printf("json 解析失败, %s ", string(body))
		}

		errorCode := urlRes.ErrorCode
		errorMsg := urlRes.ErrorMsg

		if errorCode == "0" {
			log.Printf("repair HASH#MEDIA#SERVER success")
		} else {
			log.Printf("设备：%s,节点：%s,%s修复,修复HASH#MEDIA#SERVER,失败,原因：%s,%s",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], errorCode, errorMsg)
			LogAppend(rep, errorCode, errorMsg, repairLogs)
			return
		}

	}

	streamAddrInfo, _ := redis.String(conn.Do("GET", "DEVICE-RELATION-DELIVERY-"+rep.DeviceNum))
	streamUrl, _ := redis.String(conn.Do("HGET", "DEVICE-LIVE-PLAY-URI-"+rep.DeviceNum, rep.DeviceNum))

	connSip, err := redis.Dial("tcp", *redisHost,
		redis.DialDatabase(*sipTable),
		redis.DialPassword("st8evLW3"))
	if *sipRedis != "" {
		connSip, err = redis.Dial("tcp", *sipRedis,
			redis.DialDatabase(*sipTable),
			redis.DialPassword("st8evLW3"))
	}
	if err != nil {
		log.Printf("设备：%s,节点：%s,%s修复,同步网关在线状态,失败,原因：%s",
			rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], err)
		return
	}

	//判断设备此时是否在线
	status, _ := conn.Do("ZSCORE", "DEVICE#HEART#BEAT", "\""+rep.DeviceNum+"\"")
	if status == nil {
		log.Printf("设备：%s,节点：%s,%s修复，设备已离线",
			rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason])
		LogAppend(rep, "0", "null", repairLogs)
		return
	}

	ngxStatus, err := redis.String(connSip.Do("GET", "status_"+rep.DeviceNum))

	if ngxStatus == "" || ngxStatus == "0" {
		//log.Println("同步设备状态：sip网关离线 协同在线,修复在线状态:" + rep.DeviceNum )
		_, err = conn.Do("ZREM", "DEVICE#HEART#BEAT", "\""+rep.DeviceNum+"\"")
		if err != nil {
			log.Printf("设备：%s,节点：%s,%s修复,同步网关在线状态,失败,原因：%s",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], err)
			return
		}
		log.Println("同步设备状态：sip网关离线 协同在线,修复在线状态成功:" + rep.DeviceNum)
		LogAppend(rep, "0", "null", repairLogs)
		return
	}

	//检测设备直播信息是否存在脏数据

	if streamAddrInfo == "" && streamUrl != "" {
		_, err = conn.Do("DEL", "DEVICE-LIVE-PLAY-URI-"+rep.DeviceNum)
		if err != nil {
			log.Printf("设备：%s,节点：%s,%s修复,存在脏数据直播缺key删除脏数据,失败,原因：%s",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], err)
			return
		}
	}

	if streamAddrInfo != "" && streamUrl == "" {
		_, err = conn.Do("DEL", "DEVICE-RELATION-DELIVERY-"+rep.DeviceNum)
		if err != nil {
			log.Printf("设备：%s,节点：%s,%s修复,存在脏数据直播缺key删除脏数据,失败,原因：%s",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], err)
			return

		} else {
			log.Printf("设备：%s,节点：%s,%s修复,存在脏数据直播缺key删除脏数据,成功",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason])
			streamAddrInfo = ""
		}
	}

	if streamAddrInfo != "" && streamUrl != "" {
		if rep.Reason == "stream" {
			log.Printf("设备：%s,节点：%s,%s修复,直播流存在。",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason])
			LogAppend(rep, "0", "null", repairLogs)
			return
		}
	}

	//设备无直播,开始请求直播
	if streamAddrInfo == "" {
		result := reload(rep, 0, repairLogs)
		if result == 0 {
			LogAppend(rep, "0", "null", repairLogs)
			return
		} else {
			return
		}
	}

	if streamUrl != "" {
		//判断rtsp地址是否可用
		firstPos := strings.Index(streamUrl, "publicPlay")
		publicPlay := streamUrl[firstPos:]
		pos := strings.Index(publicPlay, "rtspUri\":\"")
		end := strings.Index(publicPlay[pos+10:], "\"")
		rtspUrl := publicPlay[pos+10 : end+pos+10]

		rtspCheckUrl := "http://apm-support-platform-provider.internal/apm/sp/device/repair/rtsp/uri/check"

		parseURL, _ := url.Parse(rtspCheckUrl)

		rtspData := RtspData{
			RtspUri: rtspUrl,
		}

		b, _ := json.Marshal(rtspData)

		params := url.Values{}
		params.Set("params", string(b))
		parseURL.RawQuery = params.Encode()
		urlPathWithParams := parseURL.String()

		req, er := http.NewRequest("GET", urlPathWithParams, nil)

		if er != nil {
			log.Printf("%s unavaliable", rtspCheckUrl)
			return
		}

		//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		//
		//log.Println(req.URL.String())

		resp, err := client.Do(req)

		if err != nil {
			log.Printf("resp err: %s", err.Error())
		} else {
			if resp != nil && resp.StatusCode == 200 {
				defer resp.Body.Close()

				body, _ := ioutil.ReadAll(resp.Body)

				urlRes := rtspGetResponse{}
				jsonErr := json.Unmarshal(body, &urlRes)
				if jsonErr != nil {
					log.Printf("json 解析失败, %s ", string(body))
				}

				errorCode := urlRes.ErrorCode
				log.Printf("rtsp check url return %s:", string(body))
				if errorCode == 0 && urlRes.Data == false {
					result := reload(rep, 1, repairLogs)
					if result != 0 {
						return
					}
				}
			} else {
				log.Printf("rtsp check url return : %s", resp.Status)
			}
		}
	} else {
		log.Printf("device:%s, stream address lost.", rep.DeviceNum)
		return
	}

	//开通全天存储
	if rep.Reason == "nru" {

		openStorageUrl := "http://" + *urlIp + ":" + *urlPort + "/videocloud2/edg/cooperative/send/storage/start"
		openData := urlData{
			DeviceNum: rep.DeviceNum,
		}
		b, _ := json.Marshal(openData)
		req, _ := http.NewRequest("POST", openStorageUrl, bytes.NewBuffer(b))
		req.Header.Set("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			log.Printf("设备：%s,节点：%s,%s修复,重新拉起全天存储,失败,原因：请求%s失败,%s",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], openStorageUrl, err)
			return
		}

		if res.StatusCode != 200 {
			log.Printf("设备：%s,节点：%s,%s修复,重新拉起全天存储,失败,原因：请求%s失败,%s",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], openStorageUrl, res.Status)
			LogAppend(rep, strconv.Itoa(res.StatusCode), res.Status, repairLogs)
			return
		}

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		urlRes := urlResponse{}
		jsonErr := json.Unmarshal(body, &urlRes)
		if jsonErr != nil {
			log.Printf("json 解析失败, %s ", string(body))
		}

		errorCode := urlRes.ErrorCode
		errorMsg := urlRes.ErrorMsg

		if errorCode == "0" || errorCode == "8014" || errorCode == "8814" || errorCode == "8811" {
			log.Printf("设备：%s,节点：%s,%s修复,开通全天存储,成功成功,返回:%s,%s",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], errorCode, errorMsg)
			LogAppend(rep, "0", "null", repairLogs)
		} else {
			log.Printf("设备：%s,节点：%s,%s修复,开通全天存储,失败,原因：%s,%s",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], errorCode, errorMsg)
			LogAppend(rep, errorCode, errorMsg, repairLogs)
			return
		}

	}

	if rep.Reason == "event" {

		openEventUrl := "http://" + *urlIp + ":" + *urlPort + "/videocloud2/edg/cooperative/send/event/start"
		openData := urlData{
			DeviceNum: rep.DeviceNum,
			Type:      1,
		}
		b, _ := json.Marshal(openData)
		req, _ := http.NewRequest("POST", openEventUrl, bytes.NewBuffer(b))
		req.Header.Set("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			log.Printf("设备：%s,节点：%s,%s修复,重新拉起事件存储,失败,原因：请求%s失败,%s",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], openEventUrl, err)

			return
		}

		if res.StatusCode != 200 {
			log.Printf("设备：%s,节点：%s,%s修复,重新拉起事件存储,失败,原因：请求%s失败,%s",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], openEventUrl, res.Status)
			LogAppend(rep, strconv.Itoa(res.StatusCode), res.Status, repairLogs)
			return
		}

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		urlRes := urlResponse{}
		jsonErr := json.Unmarshal(body, &urlRes)
		if jsonErr != nil {
			log.Printf("json 解析失败, %s ", string(body))
		}
		errorCode := urlRes.ErrorCode
		errorMsg := urlRes.ErrorMsg

		if errorCode == "0" || errorCode == "8014" || errorCode == "8814" || errorCode == "8811" {
			log.Printf("设备：%s,节点：%s,%s修复,开通事件存储,成功,返回:%s,%s",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], errorCode, errorMsg)
			LogAppend(rep, "0", "null", repairLogs)
		} else {
			log.Printf("设备：%s,节点：%s,%s修复,重新拉起事件存储,失败,原因：%s,%s",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], errorCode, errorMsg)
			LogAppend(rep, errorCode, errorMsg, repairLogs)

			return
		}

	}

}

func reload(rep DeviceRep, needStop int, repairLogs *[]RepairLog) int {

	client := &http.Client{Timeout: 20 * time.Second}

	//强制停直播
	if needStop == 1 {
		stopUrl := "http://" + *urlIp + ":" + *urlPort + "/video/stream/media/live/stop"

		stopData := urlData{
			DeviceNum: rep.DeviceNum,
			UrlType:   1,
			PlayType:  1,
			Force:     "true",
			SipFlag:   "true",
		}

		b, _ := json.Marshal(stopData)
		req, _ := http.NewRequest("DELETE", stopUrl, bytes.NewBuffer(b))
		req.Header.Set("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			log.Printf("设备：%s,节点：%s,%s修复,强制停直播,失败,原因：%s",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], err)
			return -1
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		if res.StatusCode != 200 {
			log.Printf("设备：%s,节点：%s,%s修复,强制停直播,失败,原因：请求%s失败,%s",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], stopUrl, res.Status)
			LogAppend(rep, strconv.Itoa(res.StatusCode), res.Status, repairLogs)
			return -1
		}

		log.Printf("设备：%s,节点：%s,%s修复,强制停直播,成功，返回结果：%s",
			rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], string(body))

		//重新请求直播
		requestTime := time.Now().Unix()
		startUrl := "http://" + *urlIp + ":" + *urlPort + "/videocloud2/edg/cooperative/send/live/start"
		startData := urlData{
			DeviceNum: rep.DeviceNum,
			UrlType:   1,
			PlayType:  1,
		}
		b, _ = json.Marshal(startData)
		req, err = http.NewRequest("POST", startUrl, bytes.NewBuffer(b))
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Set("Content-Type", "application/json")
		res, err = client.Do(req)
		if err != nil {
			log.Printf("设备：%s,节点：%s,%s修复,重新请求直播,失败,原因：请求%s失败,%s",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], startUrl, err)
			return -1
		}

		defer res.Body.Close()
		body, _ = ioutil.ReadAll(res.Body)

		if res.StatusCode != 200 {
			log.Printf("设备：%s,节点：%s,%s修复,重新请求直播,失败,原因：请求%s失败,%s",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], startUrl, res.Status)
			LogAppend(rep, "-1", res.Status, repairLogs)
			return -1
		}

		urlRes := urlResponse{}
		jsonErr := json.Unmarshal(body, &urlRes)
		if jsonErr != nil {
			log.Printf("json 解析失败, %s ", string(body))
			return -1
		}
		errorCode := urlRes.ErrorCode
		errorMsg := urlRes.ErrorMsg

		if errorCode == "0" || errorCode == "8014" || errorCode == "8814" || errorCode == "8811" {
			log.Printf("设备：%s,节点：%s,%s修复,重新请求直播,成功,返回：%s,%s",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], errorCode, errorMsg)
			//LogAppend(rep, "请求直播", "0","null" )
			return 0
		}

		if errorCode != "" && errorMsg != "" && errorCode != "0" && strings.Contains(errorMsg, "415") {
			//设备415, 重启设备
			rebootUrl := "http://" + *urlIp + ":" + *urlPort + "/videocloud2/edg/cooperative/sip/reboot"

			req, _ = http.NewRequest("POST", rebootUrl, bytes.NewBuffer(b))
			req.Header.Set("Content-Type", "application/json")
			res, _ = client.Do(req)

			defer res.Body.Close()
			body, _ = ioutil.ReadAll(res.Body)

			if res.StatusCode != 200 {
				log.Printf("设备：%s,节点：%s,%s修复,重新设备,失败,原因：请求%s失败,%s",
					rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], startUrl, err)
				LogAppend(rep, "-1", res.Status, repairLogs)

				return -1
			} else {
				jsonErr = json.Unmarshal(body, &urlRes)
				if jsonErr != nil {
					log.Printf("json 解析失败, %s ", string(body))
				}
				errorCode = urlRes.ErrorCode
				errorMsg = urlRes.ErrorMsg

				if errorCode == "0" {
					log.Printf("设备：%s,节点：%s,%s修复,415重启设备,成功,返回:%s,%s",
						rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], errorCode, errorMsg)
					return 0
				} else {
					log.Printf("设备：%s,节点：%s,%s修复,415重启设备,失败,原因：%s,%s",
						rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], errorCode, errorMsg)
					LogAppend(rep, "-1", errorMsg, repairLogs)
					return -1
				}
			}

		}

		//设备拉流返回497 且响应超过6秒 重启设备
		if errorCode != "" && errorMsg != "" && errorCode != "0" && strings.Contains(errorMsg, "497") && time.Now().Unix()-requestTime > 6 {

			rebootUrl := "http://" + *urlIp + ":" + *urlPort + "/videocloud2/edg/cooperative/sip/reboot"

			req, _ = http.NewRequest("POST", rebootUrl, bytes.NewBuffer(b))
			req.Header.Set("Content-Type", "application/json")
			res, _ = client.Do(req)

			defer res.Body.Close()
			body, _ = ioutil.ReadAll(res.Body)

			if res.StatusCode != 200 {
				log.Printf("设备：%s,节点：%s,%s修复,497重启设备,失败,原因：请求%s失败,%s",
					rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], startUrl, err)
				LogAppend(rep, "-1", res.Status, repairLogs)

				return -1
			} else {
				jsonErr = json.Unmarshal(body, &urlRes)
				if jsonErr != nil {
					log.Printf("json 解析失败, %s ", string(body))
				}
				errorCode = urlRes.ErrorCode
				errorMsg = urlRes.ErrorMsg

				if errorCode == "0" {
					log.Printf("设备：%s,节点：%s,%s修复,497重启设备,成功,返回:%s,%s",
						rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], errorCode, errorMsg)
					return 0
				} else {
					log.Printf("设备：%s,节点：%s,%s修复,497重启设备,失败,原因：%s,%s",
						rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], errorCode, errorMsg)
					LogAppend(rep, "-1", errorMsg, repairLogs)
					return -1
				}
			}

		}

		log.Printf("设备：%s,节点：%s,%s修复,重新请求直播,失败,原因：%s请求返回:%s,%s",
			rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], startUrl, errorCode, errorMsg)

		LogAppend(rep, "-1", errorMsg, repairLogs)

		return -1

	} else {
		//重新请求直播
		requestTime := time.Now().Unix()
		startUrl := "http://" + *urlIp + ":" + *urlPort + "/videocloud2/edg/cooperative/send/live/start"
		startData := urlData{
			DeviceNum: rep.DeviceNum,
			UrlType:   1,
			PlayType:  1,
		}
		b, _ := json.Marshal(startData)
		req, err := http.NewRequest("POST", startUrl, bytes.NewBuffer(b))
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Set("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			log.Printf("设备：%s,节点：%s,%s修复,重新请求直播,失败,原因：请求%s失败,%s",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], startUrl, err)
			return -1
		}

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		if res.StatusCode != 200 {
			log.Printf("设备：%s,节点：%s,%s修复,重新请求直播,失败,原因：请求%s失败,%s",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], startUrl, res.Status)
			LogAppend(rep, strconv.Itoa(res.StatusCode), res.Status, repairLogs)
			return -1
		}

		urlRes := urlResponse{}
		jsonErr := json.Unmarshal(body, &urlRes)
		if jsonErr != nil {
			log.Printf("json 解析失败, %s ", string(body))
			return -1
		}
		errorCode := urlRes.ErrorCode
		errorMsg := urlRes.ErrorMsg

		if errorCode == "0" || errorCode == "8014" || errorCode == "8814" || errorCode == "8811" {
			log.Printf("设备：%s,节点：%s,%s修复,重新请求直播,成功,返回：%s,%s",
				rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], errorCode, errorMsg)
			//LogAppend(rep, "请求直播", "0","null" )
			return 0
		}

		if errorCode != "" && errorMsg != "" && errorCode != "0" && strings.Contains(errorMsg, "415") {
			//设备415, 重启设备
			rebootUrl := "http://" + *urlIp + ":" + *urlPort + "/videocloud2/edg/cooperative/sip/reboot"

			req, _ = http.NewRequest("POST", rebootUrl, bytes.NewBuffer(b))
			req.Header.Set("Content-Type", "application/json")
			res, _ = client.Do(req)

			defer res.Body.Close()
			body, _ = ioutil.ReadAll(res.Body)

			if res.StatusCode != 200 {
				log.Printf("设备：%s,节点：%s,%s修复,重新设备,失败,原因：请求%s失败,%s",
					rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], startUrl, err)
				LogAppend(rep, strconv.Itoa(res.StatusCode), res.Status, repairLogs)

				return -1
			} else {
				jsonErr = json.Unmarshal(body, &urlRes)
				if jsonErr != nil {
					log.Printf("json 解析失败, %s ", string(body))
				}
				errorCode = urlRes.ErrorCode
				errorMsg = urlRes.ErrorMsg

				if errorCode == "0" {
					log.Printf("设备：%s,节点：%s,%s修复,415重启设备,成功,返回:%s,%s",
						rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], errorCode, errorMsg)
					return 0
				} else {
					log.Printf("设备：%s,节点：%s,%s修复,415重启设备,失败,原因：%s,%s",
						rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], errorCode, errorMsg)
					LogAppend(rep, errorCode, errorMsg, repairLogs)
					return -1
				}
			}

		}

		//设备拉流返回497 且响应超过6秒 重启设备
		if errorCode != "" && errorMsg != "" && errorCode != "0" && strings.Contains(errorMsg, "497") && time.Now().Unix()-requestTime > 6 {

			rebootUrl := "http://" + *urlIp + ":" + *urlPort + "/videocloud2/edg/cooperative/sip/reboot"

			req, _ = http.NewRequest("POST", rebootUrl, bytes.NewBuffer(b))
			req.Header.Set("Content-Type", "application/json")
			res, _ = client.Do(req)

			defer res.Body.Close()
			body, _ = ioutil.ReadAll(res.Body)

			if res.StatusCode != 200 {
				log.Printf("设备：%s,节点：%s,%s修复,497重新设备,失败,原因：请求%s失败,%s",
					rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], startUrl, err)
				LogAppend(rep, strconv.Itoa(res.StatusCode), res.Status, repairLogs)

				return -1
			} else {
				jsonErr = json.Unmarshal(body, &urlRes)
				if jsonErr != nil {
					log.Printf("json 解析失败, %s ", string(body))
				}
				errorCode = urlRes.ErrorCode
				errorMsg = urlRes.ErrorMsg

				if errorCode == "0" {
					log.Printf("设备：%s,节点：%s,%s修复,497重启设备,成功,返回:%s,%s",
						rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], errorCode, errorMsg)
					return 0
				} else {
					log.Printf("设备：%s,节点：%s,%s修复,497重启设备,失败,原因：%s,%s",
						rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], errorCode, errorMsg)
					LogAppend(rep, errorCode, errorMsg, repairLogs)
					return -1
				}
			}

		}

		log.Printf("设备：%s,节点：%s,%s修复,重新请求直播,失败,原因：%s请求返回:%s,%s",
			rep.DeviceNum, rep.NodeCode, reasonMap[rep.Reason], startUrl, errorCode, errorMsg)

		LogAppend(rep, errorCode, errorMsg, repairLogs)

		return -1
	}

}

func repairLogCsvFile() *os.File {
	today := time.Now().Format("2006-01-02")

	file, err := os.Open(*fileUrl + "repair-log-" + today + ".csv")
	// 如果文件不存在，创建文件
	if err != nil && os.IsNotExist(err) {
		file, err = os.Create(*fileUrl + "repair-log-" + today + ".csv")
		if err != nil {
			log.Println(err)
		}

		w := csv.NewWriter(file) //创建一个新的写入文件流
		title := []string{"Operator", "Time", "DeviceNum", "NodeCode", "CorpCode", "ModelCode",
			"Reason", "ErrorCode", "ErrorMsg"}

		// 这里必须刷新，才能将数据写入文件。
		w.Write(title)
		w.Flush()
	} else {
		file, err = os.OpenFile(*fileUrl+"repair-log-"+today+".csv", os.O_APPEND|os.O_RDWR, 0644)
		if err != nil {
			log.Println("open file is failed, err: ", err)
		}
	}
	return file
}

func goRepair(deviceReps []DeviceRep) {
	repairLogs := make([]RepairLog, 0, len(deviceReps))

	//repairLog.Printf("start repair : %d",len(deviceReps) )

	for _, rep := range deviceReps {
		wg.Add(1)
		go Repair(rep, &repairLogs)
	}

	wg.Wait()

	logfile := repairLogCsvFile()

	returnCodeStr := make(map[string]int)
	returnCodeEve := make(map[string]int)
	returnCodeNru := make(map[string]int)
	jsonLogs := "["

	for _, repairLog := range repairLogs {
		logfile.WriteString("\xEF\xBB\xBF")
		w := csv.NewWriter(logfile)
		w.Write([]string{"auto-repair", strconv.FormatInt(repairLog.Time, 10), repairLog.DeviceNum, repairLog.NodeCode,
			repairLog.CorpCode, repairLog.ModelCode, repairLog.Reason, repairLog.ErrorCode, repairLog.ErrorMsg})
		w.Flush()

		jsonStr, _ := json.Marshal(repairLog)

		newJson := strings.Replace(string(jsonStr), "null", "", -1)

		jsonLogs = jsonLogs + newJson + ","

		if repairLog.Reason == "stream" {

			_, ok := returnCodeStr[repairLog.ErrorCode]
			if ok {
				returnCodeStr[repairLog.ErrorCode]++
			} else {
				returnCodeStr[repairLog.ErrorCode] = 1
			}
		}

		if repairLog.Reason == "event" {
			_, ok := returnCodeEve[repairLog.ErrorCode]
			if ok {
				returnCodeEve[repairLog.ErrorCode]++
			} else {
				returnCodeEve[repairLog.ErrorCode] = 1
			}
		}

		if repairLog.Reason == "nru" {
			_, ok := returnCodeNru[repairLog.ErrorCode]
			if ok {
				returnCodeNru[repairLog.ErrorCode]++
			} else {
				returnCodeNru[repairLog.ErrorCode] = 1
			}
		}
	}

	jsonLogs = jsonLogs[:len(jsonLogs)-1] + "]"
	//println(jsonLogs)
	if len(repairLogs) != 0 {
		repairLogPost(jsonLogs)
	}

	for k, v := range returnCodeStr {
		httpOpentsdbDo("video.home.repair.ipc", v, k, "stream", "repair")
		log.Printf("node:%s, stream repair result :%s , %d", *nodeRepair, k, v)
	}

	for k, v := range returnCodeEve {
		httpOpentsdbDo("video.home.repair.ipc", v, k, "event", "repair")
		log.Printf("node:%s, event repair result :%s , %d", *nodeRepair, k, v)
	}

	for k, v := range returnCodeNru {
		httpOpentsdbDo("video.home.repair.ipc", v, k, "nru", "repair")
		log.Printf("node:%s, nru repair result :%s , %d", *nodeRepair, k, v)
	}
	//
	//for k,v := range repairFreq{
	//	if v > 0 {
	//		repairFreq[k]--
	//		//repairLog.Printf("%s : %d",k,repairFreq[k])
	//	}else {
	//		delete(repairFreq,k)
	//	}
	//}

	log.Println("over")
}

func selectRep() {
	deviceReps := make([]DeviceRep, 0, 501)
	if start == 1 {
		goRepair(Read(*nodeRepair, deviceReps))
	}
}

func startRepair(writer http.ResponseWriter, request *http.Request) {
	var result Resp
	if request.Method == "GET" {
		start = 1
		result.Code = 0
		result.Msg = "自动修复启动成功"
	} else {
		result.Code = 401
		result.Msg = "自动修复启动失败"
	}
	if err := json.NewEncoder(writer).Encode(result); err != nil {
		log.Println(err)
	}
	log.Printf("node: %s start repair.", *nodeRepair)
}

func stopRepair(writer http.ResponseWriter, request *http.Request) {
	var result Resp
	if request.Method == "GET" {
		start = 0
		result.Code = 0
		result.Msg = "自动修复暂停成功"
	} else {
		result.Code = 401
		result.Msg = "自动修复暂停失败"
	}
	if err := json.NewEncoder(writer).Encode(result); err != nil {
		log.Println(err)
	}
	log.Printf("node: %s stop repair.", *nodeRepair)
}

func GetRepairStatus(writer http.ResponseWriter, request *http.Request) {
	var result Resp
	var data Data
	if request.Method == "GET" {
		result.Code = 0
		result.Msg = "success"
		result.Data = &data
		data.Status = start
	} else {
		result.Code = 401
		result.Msg = "获取修复状态失败"
	}
	if err := json.NewEncoder(writer).Encode(result); err != nil {
		log.Println(err)
	}
	log.Printf("%s get repair status : %d", *nodeRepair, start)
}

func httpListen() {
	http.HandleFunc("/start", startRepair)
	http.HandleFunc("/stop", stopRepair)
	http.HandleFunc("/status", GetRepairStatus)

	http.ListenAndServe("0.0.0.0:"+*startPort, nil)
}

func repairLogPost(repairLog string) {

	client := &http.Client{Timeout: 10 * time.Second}

	//TODO
	req, err := http.NewRequest("POST", "http://apm-support-platform-provider.internal/apm/autoscript/repair/logs",
		bytes.NewBuffer([]byte(repairLog)))

	if err != nil {
		log.Println("can not connect to apm service.")
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("can not connect to apm service.")
		return
	}

	if resp.StatusCode != 200 {
		log.Printf("apm service status : %s,%s", resp.Status, resp.Body)
		return
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	urlRes := logPostResponse{}
	jsonErr := json.Unmarshal(body, &urlRes)
	if jsonErr != nil {
		log.Printf("json 解析失败, %s ", string(body))
		return
	}
	errorCode := urlRes.ErrorCode
	errorMsg := urlRes.ErrorMsg

	if errorCode != 0 {
		log.Println(errorMsg)
	}

}

func main() {

	flag.Parse()

	go httpListen()

	//redisIpMap = make(map[string]string)
	reasonMap = make(map[string]string)

	repairFreq = make(map[string]int)

	repairLastTime = make(map[string]int64)

	reasonMap["stream"] = "直播流"
	reasonMap["nru"] = "全天存储"
	reasonMap["event"] = "事件存储"

	spec := "20 * * * * * "

	c := cron.New(cron.WithSeconds())
	id, err := c.AddFunc(spec, func() {
		selectRep()
	})

	if err != nil {
		fmt.Sprintf("crontab exec error: %v with id: %v", err, id)
	}

	c.Start()

	select {}
}
