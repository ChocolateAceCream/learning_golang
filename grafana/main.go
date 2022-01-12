package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
)

var node = flag.String("node", "", "input node code")
var redisHost = flag.String("redis_host", "", "redis host address")
var redisPassword = flag.String("redis_password", "", "redis password")
var database = flag.Int("db", 1, "database")
var opentsdb = flag.String("opentsdb", "", "opentsdb address")
var wg sync.WaitGroup
var today string
var preOneMin = time.Now().Add(-time.Minute * 1).Format("2006-01-02 15:04")
var mutex sync.Mutex
var redisPool *redis.Pool

var devices []Device

type Device struct {
	Time          int64  `json:"timestamp"`
	DeviceNum     string `json:"deviceNum"`
	CorpCode      string `json:"corpCode"`
	ModelCode     string `json:"modelCode"`
	NodeCode      string `json:"nodeCode"`
	Status        string `json:"onlineStatus"`
	StreamIpAddr  string `json:"streamIp"`
	StreamIpSub   string `json:"subStreamIp"`
	StreamIpBack  string `json:"playbackStreamIp"`
	Combo         string `json:"storageType"`
	StorageIpAddr string `json:"storageIp"`
	AreaCode      string `json:"areaCode"`
	LostRate      string `json:"lostRate"`
}

type JsonDevice struct {
	Time          int64  `json:"timestamp"`
	DeviceNum     string `json:"deviceNum"`
	CorpCode      string `json:"corpCode"`
	ModelCode     string `json:"modelCode"`
	NodeCode      string `json:"nodeCode"`
	Status        string `json:"onlineStatus"`
	StreamIpAddr  string `json:"streamIp"`
	StreamIpSub   string `json:"subStreamIp"`
	StreamIpBack  string `json:"playbackStreamIp"`
	Combo         string `json:"storageType"`
	StorageIpAddr string `json:"storageIp"`
	AreaCode      string `json:"areaCode"`
	LostRate      int    `json:"lostRate"`
}

type Tags struct {
	Area   string `json:"area"`
	IP     string `json:"ip"`
	Combo  string `json:"combo"`
	Target string `json:"target"`
}

type QualityTag struct {
	Area    string `json:"area"`
	IP      string `json:"ip"`
	Quality int    `json:"quality"`
	Target  string `json:"target"`
}

type TsdbItem struct {
	Metric    string `json:"metric"`
	Tags      Tags   `json:"tags"`
	Value     int    `json:"value"`
	Timestamp int64  `json:"timestamp"`
}

type TsdbQuality struct {
	Metric    string     `json:"metric"`
	Tags      QualityTag `json:"tags"`
	Value     int        `json:"value"`
	Timestamp int64      `json:"timestamp"`
}

type TsdbFloat struct {
	Metric    string  `json:"metric"`
	Tags      Tags    `json:"tags"`
	Value     float64 `json:"value"`
	Timestamp int64   `json:"timestamp"`
}

type PostResponse struct {
	ErrorCode string `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
	Data      string `json:"data,omitempty"`
}

type MysqlPostResponse struct {
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
	Data      string `json:"data,omitempty"`
}

type Job struct {
	DeviceInfo string
	DeviceCode string
}

//上传整数型数据到opentsdb
func httpOpentsdbDo(metric string, value int, ip string, combo string, target string) {

	tag := Tags{
		Area:   *node,
		Target: target,
		Combo:  combo,
		IP:     ip,
	}

	data := TsdbItem{
		Metric:    metric,
		Value:     value,
		Timestamp: time.Now().Unix(),
		Tags:      tag,
	}

	client := &http.Client{Timeout: 30 * time.Second}

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

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("can not connect to opentsdb.")
	}

	//log.Println(string(body) + metric)

}

//上传浮点数型数据到opentsdb
func httpOpentsdbDoFloat(metric string, value float64, ip string, combo string, target string) {

	tag := Tags{
		Area:   *node,
		Target: target,
		Combo:  combo,
		IP:     ip,
	}

	data := TsdbFloat{
		Metric:    metric,
		Value:     value,
		Timestamp: time.Now().Unix(),
		Tags:      tag,
	}

	client := &http.Client{Timeout: 30 * time.Second}

	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", *opentsdb+"/api/put?details", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("can not connect to opentsdb.")
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("can not connect to opentsdb.")
	}

	//log.Println(string(body) + metric)

}

//上传流质量数据到opentsdb
func httpQualityOpentsdbDo(metric string, value int, ip string, quality int, target string) {

	tag := QualityTag{
		Area:    *node,
		Target:  target,
		Quality: quality,
		IP:      ip,
	}

	data := TsdbQuality{
		Metric:    metric,
		Value:     value,
		Timestamp: time.Now().Unix(),
		Tags:      tag,
	}

	client := &http.Client{Timeout: 30 * time.Second}

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

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("can not connect to opentsdb.")
	}

	//log.Printf("value
	//:%d,ip:%s",value,ip)

}

//数据写入csv文件
func deviceCsvFile() *os.File {

	file, err := os.Open("device-" + today + ".csv")
	// 如果文件不存在，创建文件
	if err != nil && os.IsNotExist(err) {
		file, err = os.Create("device-" + today + ".csv")
		if err != nil {
			log.Println(err)
		}

		w := csv.NewWriter(file) //创建一个新的写入文件流
		title := []string{"Time", "DeviceNum", "CorpCode", "ModelCode", "NodeCode",
			"Status", "StreamIp", "StreamIpSub", "StreamIpBack", "Combo", "StorageIp", "AreaCode", "LostRate"}

		// 这里必须刷新，才能将数据写入文件。
		w.Write(title)
		w.Flush()
	} else {
		file, err = os.OpenFile("device-"+today+".csv", os.O_APPEND|os.O_RDWR, 0644)
		if err != nil {
			log.Println("open file is failed, err: ", err)
		}
	}
	return file
}

//修复数据csv文件
func repairCsvFile() *os.File {
	file, err := os.Open("device-repair-" + today + ".csv")
	// 如果文件不存在，创建文件
	if err != nil && os.IsNotExist(err) {
		file, err = os.Create("device-repair-" + today + ".csv")
		if err != nil {
			log.Println(err)
		}

		w := csv.NewWriter(file) //创建一个新的写入文件流
		title := []string{"Time", "DeviceNum", "NodeCode", "Reason", "CorpCode", "ModelCode"}

		// 这里必须刷新，才能将数据写入文件。
		w.Write(title)
		w.Flush()
	} else {
		file, err = os.OpenFile("device-repair-"+today+".csv", os.O_APPEND|os.O_RDWR, 0644)
		if err != nil {
			log.Println("open file is failed, err: ", err)
		}
	}
	return file
}

//redis数据采集主流程
func redisDataCol(conn redis.Conn) {

	defer conn.Close()

	deviceAll, err := redis.StringMap(conn.Do("HGETALL", "HASH#DEVICE#INFO"))
	if err != nil {
		log.Println("can't find device.")
		return
	}

	eventAlarm := 0
	eventAlarm, err = redis.Int(conn.Do("ZSCORE", "DEVICE#ALARM#COUNT", "\""+preOneMin+"\""))
	if err == nil {
		httpOpentsdbDo("video.home.load.ipc", eventAlarm, "none", "none", "eventalarm")
		_, err = conn.Do("ZREM", "DEVICE#ALARM#COUNT", "\""+preOneMin+"\"")
		if err != nil {
			log.Println("err while deleting alarm:", err)
		}
		log.Println("evant alarm is :", eventAlarm)
	} else {
		log.Println("err while getting alarm:", err)
	}

	eventSuccess := 0
	eventSuccess, err = redis.Int(conn.Do("ZSCORE", "DEVICE#ALARM#SUCCESS#COUNT", "\""+preOneMin+"\""))
	if err != nil {
		log.Println("err while getting event success:", err)
	}

	eventRatioStore := float64(0)
	if eventSuccess != 0 && eventAlarm != 0 {
		eventRatioStore = float64(eventSuccess*100) / float64(eventAlarm)
		log.Printf("event ratio store is : %2f", eventRatioStore)
	}
	httpOpentsdbDoFloat("video.home.ratio.store", eventRatioStore, "none", "event", "success")

	numOnline := 0

	numAll := len(deviceAll)

	log.Printf("redis device num is : %d ", len(deviceAll))

	file := deviceCsvFile()
	repairFile := repairCsvFile()
	devices = make([]Device, 0, numAll)

	wg.Add(numAll)

	for deviceNum, deviceInfo := range deviceAll {
		go deviceDataProc(deviceInfo, deviceNum)
	}

	wg.Wait()

	httpOpentsdbDo("video.home.load.ipc", len(devices), "none", "none", "device")

	log.Printf("collect device num is : %d", len(devices))
	numOnlineNru := 0
	numOnlineEvent := 0
	numOnlineNone := 0
	//全天成功直播数
	numIpNru := 0
	//事件成功直播数
	numIpEvent := 0
	onlineRate := float64(0)
	streamRate := float64(0)
	storeEveRate := float64(0)
	storeNruRate := float64(0)
	streamData := make(map[string]int)
	storageEventData := make(map[string]int)
	storageNruData := make(map[string]int)

	lostAData := make(map[string]int)
	lostBData := make(map[string]int)
	lostCData := make(map[string]int)
	lostDData := make(map[string]int)
	lostEData := make(map[string]int)
	lostFData := make(map[string]int)

	devList := "["
	lostRateNum := 0
	for _, dev := range devices {
		file.WriteString("\xEF\xBB\xBF")
		w := csv.NewWriter(file)
		w.Write([]string{strconv.FormatInt(dev.Time, 10), dev.DeviceNum, dev.CorpCode, dev.ModelCode,
			dev.NodeCode, dev.Status, dev.StreamIpAddr, dev.StreamIpSub, dev.StreamIpBack,
			dev.Combo, dev.StorageIpAddr, dev.AreaCode, dev.LostRate})
		w.Flush()

		if dev.Status == "online" {
			switch dev.Combo {
			case "nru":
				numOnlineNru++
			case "event":
				numOnlineEvent++
			case "none":
				numOnlineNone++
			}
		}

		if dev.StreamIpAddr != "null" && dev.Status == "online" {
			_, ok := streamData[dev.StreamIpAddr]
			if ok {
				streamData[dev.StreamIpAddr]++
			} else {
				streamData[dev.StreamIpAddr] = 1
			}
			if dev.Combo == "nru" {
				numIpNru++
			} else if dev.Combo == "event" {
				numIpEvent++
			}

			if dev.LostRate != "null" {
				intLost, _ := strconv.Atoi(dev.LostRate)
				switch {
				case intLost > 15:
					_, ok = lostAData[dev.StreamIpAddr]
					if ok {
						lostAData[dev.StreamIpAddr]++
					} else {
						lostAData[dev.StreamIpAddr] = 1
					}
				case intLost <= 15 && intLost > 10:
					_, ok = lostBData[dev.StreamIpAddr]
					if ok {
						lostBData[dev.StreamIpAddr]++
					} else {
						lostBData[dev.StreamIpAddr] = 1
					}
				case intLost <= 10 && intLost > 5:
					_, ok = lostCData[dev.StreamIpAddr]
					if ok {
						lostCData[dev.StreamIpAddr]++
					} else {
						lostCData[dev.StreamIpAddr] = 1
					}

				case intLost <= 5 && intLost > 0:
					_, ok = lostDData[dev.StreamIpAddr]
					if ok {
						lostDData[dev.StreamIpAddr]++
					} else {
						lostDData[dev.StreamIpAddr] = 1
					}

				case intLost == 0:
					_, ok = lostEData[dev.StreamIpAddr]
					if ok {
						lostEData[dev.StreamIpAddr]++
					} else {
						lostEData[dev.StreamIpAddr] = 1
					}
				}
				lostRateNum++
			} else {
				_, ok = lostFData[dev.StreamIpAddr]
				if ok {
					lostFData[dev.StreamIpAddr]++
				} else {
					lostFData[dev.StreamIpAddr] = 1
				}
			}

		}

		if dev.StreamIpAddr == "null" && dev.Status == "online" {
			repairFile.WriteString("\xEF\xBB\xBF")
			wr := csv.NewWriter(repairFile)
			wr.Write([]string{strconv.FormatInt(dev.Time, 10), dev.DeviceNum, dev.NodeCode, "stream", dev.CorpCode, dev.ModelCode})
			wr.Flush()
		}

		if dev.StorageIpAddr != "null" && dev.Status == "online" && dev.StreamIpAddr != "null" {
			if dev.Combo == "event" {
				_, ok := storageEventData[dev.StorageIpAddr]
				if ok {
					storageEventData[dev.StorageIpAddr]++
				} else {
					storageEventData[dev.StorageIpAddr] = 1
				}
			} else if dev.Combo == "nru" {
				_, ok := storageNruData[dev.StorageIpAddr]
				if ok {
					storageNruData[dev.StorageIpAddr]++
				} else {
					storageNruData[dev.StorageIpAddr] = 1
				}
			}
		}

		if dev.StorageIpAddr == "null" && dev.Status == "online" && dev.Combo != "none" && dev.StreamIpAddr != "null" {
			repairFile.WriteString("\xEF\xBB\xBF")
			wr := csv.NewWriter(repairFile)
			wr.Write([]string{strconv.FormatInt(dev.Time, 10), dev.DeviceNum, dev.NodeCode, dev.Combo, dev.CorpCode, dev.ModelCode})
			wr.Flush()
		}

		var lostRate int

		if dev.LostRate != "null" {
			lostRate, _ = strconv.Atoi(dev.LostRate)
		}

		devJson := JsonDevice{
			dev.Time,
			dev.DeviceNum,
			dev.CorpCode,
			dev.ModelCode,
			dev.NodeCode,
			dev.Status,
			dev.StorageIpAddr,
			dev.StreamIpSub,
			dev.StreamIpBack,
			dev.Combo,
			dev.StorageIpAddr,
			dev.AreaCode,
			lostRate,
		}

		jsonStr, _ := json.Marshal(devJson)

		newJson := strings.Replace(string(jsonStr), "null", "", -1)

		devList += newJson + ","

	}

	log.Printf("lost rate post num is : %d", lostRateNum)

	defer file.Close()
	defer repairFile.Close()
	numOnline = numOnlineNone + numOnlineNru + numOnlineEvent
	log.Printf("nru:%d , event: %d , none: %d", numOnlineNru, numOnlineEvent, numOnlineNone)

	if numOnline != 0 && len(devices) != 0 {
		onlineRate = float64(numOnline*100) / float64(len(devices))
	}

	httpOpentsdbDo("video.home.load.ipc", numOnlineNru, "none", "nru", "online")
	httpOpentsdbDo("video.home.load.ipc", numOnlineEvent, "none", "event", "online")
	httpOpentsdbDo("video.home.load.ipc", numOnlineNone, "none", "none", "online")
	httpOpentsdbDoFloat("video.home.ratio.compute", onlineRate, "none", "none", "online")

	streamAll := 0
	for k, v := range streamData {
		httpOpentsdbDo("video.home.load.ipc", v, k, "none", "stream")
		streamAll += v
	}
	if streamAll != 0 && numOnline != 0 {
		streamRate = float64(streamAll*100) / float64(numOnline)
	}
	httpOpentsdbDoFloat("video.home.ratio.compute", streamRate, "none", "none", "stream")

	//存储正常率
	storeEvent := 0
	for k, v := range storageEventData {
		httpOpentsdbDo("video.home.load.ipc", v, k, "event", "store")
		storeEvent += v
	}
	if storeEvent != 0 && numOnlineEvent != 0 {
		storeEveRate = float64(storeEvent*100) / float64(numIpEvent)
	}
	httpOpentsdbDoFloat("video.home.ratio.compute", storeEveRate, "none", "success", "event")

	storeNru := 0
	for k, v := range storageNruData {
		httpOpentsdbDo("video.home.load.ipc", v, k, "nru", "store")
		storeNru += v
	}
	if storeNru != 0 && numOnlineNru != 0 {
		storeNruRate = float64(storeNru*100) / float64(numIpNru)
	}
	httpOpentsdbDoFloat("video.home.ratio.compute", storeNruRate, "none", "success", "nru")

	//全天存储成功率 5分钟取一次

	if time.Now().Minute()%5 == 0 {
		nruSuccess := 0
		preFiveMin := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(),
			time.Now().Minute()-5, 0, 0, time.Now().Location()).Format("2006-01-02 15:04")
		log.Println(preFiveMin)
		nruSuccess, err = redis.Int(conn.Do("ZSCORE", "DEVICE#STORAGE#SUCCESS#COUNT", "\""+preFiveMin+"\""))
		if err != nil {
			log.Println("err while getting nru success:", err)
		}
		//nruStoreAll := 0
		//nruStoreAll,err = redis.Int(conn.Do("ZSCORE", "DEVICE#STORAGE#COUNT", "\""+preFiveMin+"\""))
		//if err != nil {
		//	log.Println("err while getting nru store all:", err)
		//}
		nruRatioStore := float64(0)
		if nruSuccess != 0 && numIpNru != 0 {
			nruRatioStore = float64(nruSuccess*100) / float64(numIpNru)
			log.Println("nru ratio store is s: ", nruRatioStore)
		}
		httpOpentsdbDoFloat("video.home.ratio.store", nruRatioStore, "none", "nru", "success")

	}

	//丢包率分级统计
	for k, v := range lostAData {
		httpQualityOpentsdbDo("video.home.load.ipc", v, k, 5, "lost")
	}

	for k, v := range lostBData {
		httpQualityOpentsdbDo("video.home.load.ipc", v, k, 4, "lost")
	}

	for k, v := range lostCData {
		httpQualityOpentsdbDo("video.home.load.ipc", v, k, 3, "lost")
	}

	for k, v := range lostDData {
		httpQualityOpentsdbDo("video.home.load.ipc", v, k, 2, "lost")
	}

	for k, v := range lostEData {
		httpQualityOpentsdbDo("video.home.load.ipc", v, k, 1, "lost")
	}

	for k, v := range lostFData {
		httpQualityOpentsdbDo("video.home.load.ipc", v, k, 0, "lost")
	}

	storageLoad(conn)
	eventLoad(conn)
	pushOverstack(conn)

	devList = devList[:len(devList)-1] + "]"
	//println(jsonLogs)
	if len(devices) != 0 {
		hbasePost(devList)
		mysqlPost(devList)
		log.Printf("hbase post device num: %d", len(devices))
	}
}

//多协程采集设备数据
func deviceDataProc(deviceInfo string, deviceCode string) {

	defer wg.Done()
	//ch <- true
	conn, err := redis.Dial("tcp", *redisHost,
		redis.DialDatabase(*database),
		redis.DialPassword(*redisPassword),
		redis.DialConnectTimeout(10*time.Second),
		redis.DialReadTimeout(10*time.Second),
		redis.DialWriteTimeout(10*time.Second),
	)
	if err != nil {
		log.Println("redis get error:", err)
		return
	}

	defer conn.Close()

	_, err = redis.String(conn.Do("HGET", "HASH#DEVICE#INFO", deviceCode))
	if err != nil {
		log.Printf("device:%s, err: %s ", deviceCode, err.Error())
		return
	}

	deviceStatus := "offline"
	combo := "none"
	unitId := ""
	storageAddr := "null"
	streamAddr := "null"
	streamIpSub := "null"
	streamIpBack := "null"

	if strings.Contains(deviceInfo, "nodeCode") {
		pos := strings.Index(deviceInfo, "nodeCode\\\":\\\"")
		end := strings.Index(deviceInfo[pos+13:], "\\\"")
		nodeCode := deviceInfo[pos+13 : end+pos+13]
		devStatus := "1"
		if strings.Contains(deviceInfo, "status") {
			pos = strings.Index(deviceInfo, "status")
			devStatus = deviceInfo[pos+9 : pos+10]
		}

		//设备状态为9 说明设备已迁移 超过5分钟删除设备信息
		if devStatus == "9" {
			log.Println("device status is 9: " + deviceCode)
			updateTimeStr, _ := redis.String(conn.Do("HGET", "HASH#DEVICE#INFO#UPDATETIME", deviceCode))
			if updateTimeStr != "" {
				updateTime, _ := strconv.ParseInt(updateTimeStr, 10, 64)
				if time.Now().UnixNano()-updateTime > 30000 {
					_, err = conn.Do("HDEL", "HASH#DEVICE#INFO", deviceCode)
					if err != nil {
						log.Printf("device: %s ,delete device info fail: %s ", deviceCode, err.Error())
					}
					_, err = conn.Do("ZREM", "DEVICE#HEART#BEAT", deviceCode)
					if err != nil {
						log.Printf("device: %s ,delete device heartbeat fail: %s ", deviceCode, err.Error())
					}
					_, err = redis.String(conn.Do("HDEL", "HASH#DEVICE#STORAGE", deviceCode))
					_, err = redis.String(conn.Do("HDEL", "HASH#DEVICE#EVENT#STORAGE", deviceCode))
					log.Printf("device: %s delete success", deviceCode)
				}
			}
			return
		}
		if nodeCode == *node {

			pos = strings.Index(deviceInfo, "corpCode\\\":\\\"")
			end = strings.Index(deviceInfo[pos+13:], "\\\"")
			corpCode := deviceInfo[pos+13 : end+pos+13]

			pos = strings.Index(deviceInfo, "modelCode\\\":\\\"")
			end = strings.Index(deviceInfo[pos+14:], "\\\"")
			modelCode := deviceInfo[pos+14 : end+pos+14]

			pos = strings.Index(deviceInfo, "provinceCode\\\":\\\"")
			end = strings.Index(deviceInfo[pos+17:], "\\\"")
			provinceCode := deviceInfo[pos+17 : end+pos+17]

			pos = strings.Index(deviceInfo, "cityCode\\\":\\\"")
			end = strings.Index(deviceInfo[pos+13:], "\\\"")
			cityCode := deviceInfo[pos+13 : end+pos+13]

			pos = strings.Index(deviceInfo, "countryCode\\\":\\\"")
			end = strings.Index(deviceInfo[pos+16:], "\\\"")
			countryCode := deviceInfo[pos+16 : end+pos+16]

			areaCode := provinceCode + cityCode + countryCode

			lostRate := "null"

			//丢包数统计
			qualityRes, _ := redis.String(conn.Do("HGET", "STREAM-QUALITY-RESULT-STAT#"+deviceCode, preOneMin))
			if qualityRes != "" {
				pos = strings.Index(qualityRes, "lostRate")
				end = strings.Index(qualityRes[pos+11:], ",")
				lostRate = qualityRes[pos+11 : end+pos+11]
				//log.Printf("device: %s, lost rate : %s",deviceCode,lostRate)
				//删除
				_, err := conn.Do("HDEL", "STREAM-QUALITY-RESULT-STAT#"+deviceCode, preOneMin)
				if err != nil {
					log.Printf("device: %s, delete quality err: %s.", deviceCode, err.Error())
				}

			}

			res, _ := conn.Do("ZRANK", "DEVICE#HEART#BEAT", "\""+deviceCode+"\"")
			if res != nil {
				deviceStatus = "online"
			}
			streamAddrInfo, _ := redis.String(conn.Do("GET", "DEVICE-RELATION-DELIVERY-"+deviceCode))
			if streamAddrInfo != "" {
				streamAddr = streamAddrInfo[8 : len(streamAddrInfo)-6]
			}

			streamUrl, _ := redis.String(conn.Do("HGET", "DEVICE-LIVE-PLAY-URI-"+deviceCode, deviceCode))
			if streamUrl != "" {
				streamIpSubInfo, _ := redis.String(conn.Do("GET", "DEVICE-RELATION-DELIVERY-"+deviceCode+"_sub"))
				if streamIpSubInfo != "" {
					streamIpSub = streamIpSubInfo[8 : len(streamIpSubInfo)-6]
				}
			}

			bindStatus := "2"

			//判断套餐类型
			isAllDay, _ := redis.String(conn.Do("HGET", "HASH#DEVICE#STORAGE", deviceCode))
			if isAllDay != "" {
				if strings.Contains(isAllDay, "bindStatus") {
					pos = strings.Index(isAllDay, "bindStatus")
					end = strings.Index(isAllDay[pos+13:], ",")
					bindStatus = isAllDay[pos+13 : end+pos+13]
				}
				if bindStatus == "1" {
					combo = "nru"
					unitId, _ = redis.String(conn.Do("GET", "DEVICE_STORAGE_MEDIA_RELATION_"+deviceCode))
					if unitId != "" {
						end = strings.LastIndex(unitId, "\"")
						unitId = unitId[1:end]
						mediaInfo, _ := redis.String(conn.Do("HGET", "STORAGE_MEDIA_MESSAGE", unitId))
						if mediaInfo != "" {
							pos = strings.Index(mediaInfo, "baseUrl")
							end = strings.Index(mediaInfo[pos+9:], ",")
							storageAddr = mediaInfo[pos+17 : end+pos+3]
						}
					}
				}
			}

			if combo != "nru" {
				isEvent, _ := redis.String(conn.Do("HGET", "HASH#DEVICE#EVENT#STORAGE", deviceCode))
				if isEvent != "" {
					if strings.Contains(isEvent, "eventStorageStatus") {
						pos = strings.Index(isEvent, "eventStorageStatus")
						end = strings.Index(isEvent[pos+21:], ",")
						bindStatus = isEvent[pos+21 : end+pos+21]
					}
					if bindStatus != "1" {
						combo = "none"
					} else {
						combo = "event"
						unitId, _ = redis.String(conn.Do("GET", "EVENT-RELATION-DEVICE-"+deviceCode))
						if unitId != "" {
							end = strings.LastIndex(unitId, "\"")
							unitId = unitId[1:end]
							mediaInfo, _ := redis.String(conn.Do("GET", "EVENT-MEDIA-SERVER-"+unitId))
							if mediaInfo != "" {
								pos = strings.Index(mediaInfo, "baseUrl")
								end = strings.Index(mediaInfo[pos+9:], ",")
								storageAddr = mediaInfo[pos+17 : end+pos+3]
							}
						}
					}
				}
			}

			dev := Device{
				time.Now().Unix(),
				deviceCode,
				corpCode,
				modelCode,
				nodeCode,
				deviceStatus,
				streamAddr,
				streamIpSub,
				streamIpBack,
				combo,
				storageAddr,
				areaCode,
				lostRate,
			}
			mutex.Lock()
			devices = append(devices, dev)
			mutex.Unlock()

		} else {
			log.Println("device is not belong node: " + deviceCode)
		}
	} else {
		log.Println("device do not have nodeCode: " + deviceCode)
	}
}

//全天存储数据采集
func storageLoad(conn redis.Conn) {
	storageSucc := 0
	storageAll := 0
	storageStack := 0
	//nodeSucc := 0
	//nodeAll := 0
	//newStoreNruRate:=float64(0)

	storageNruMap, _ := redis.IntMap(conn.Do("ZRANGE", "STORAGE_MEDIA_ACCESS_NUM", 0, -1, "withscores"))

	if storageNruMap != nil {
		for k, _ := range storageNruMap {
			storageInfo, _ := redis.String(conn.Do("HGET", "STORAGE-MEDIA-REPORT-"+k[1:len(k)-1], preOneMin))
			storageAddrIp, _ := redis.String(conn.Do("HGET", "STORAGE_MEDIA_MESSAGE", k[1:len(k)-1]))
			if storageAddrIp != "" {
				pos := strings.Index(storageAddrIp, "baseUrl")
				end := strings.Index(storageAddrIp[pos+9:], ",")
				storageAddrIp = storageAddrIp[pos+17 : end+pos+3]
			}
			if storageInfo != "" && strings.Contains(storageInfo, "storageTotal") {
				//log.Println(storageInfo)
				//应存储总量
				pos := strings.Index(storageInfo, "storageTotal")
				end := strings.Index(storageInfo[pos+15:], ",")
				storageAll, _ = strconv.Atoi(storageInfo[pos+15 : end+pos+15])
				//nodeAll += storageAll

				//成功量
				pos = strings.Index(storageInfo, "storageSucc")
				end = strings.Index(storageInfo[pos+14:], ",")

				storageSucc, _ = strconv.Atoi(storageInfo[pos+14 : end+pos+14])
				//nodeSucc += storageSucc

				//存储积压
				pos = strings.Index(storageInfo, "storageUploading")

				storageStack, _ = strconv.Atoi(storageInfo[pos+19 : pos+20])
				//log.Println(storageStack)
			}
			httpOpentsdbDo("video.home.load.ipc", storageAll, storageAddrIp, "nru", "allst")
			httpOpentsdbDo("video.home.load.ipc", storageSucc, storageAddrIp, "nru", "successst")
			httpOpentsdbDo("video.home.load.ipc", storageStack, storageAddrIp, "nru", "overstack")

		}
		//if nodeSucc != 0 && nodeAll != 0 {
		//	newStoreNruRate = float64(nodeSucc*100)/float64(nodeAll)
		//}
		//httpOpentsdbDoFloat("video.home.ratio.compute",newStoreNruRate,"none","nru","successmake")

	}
}

//事件存储数据采集
func eventLoad(conn redis.Conn) {
	storageSucc := 0
	storageAll := 0
	storageStack := 0
	//nodeSucc := 0
	//nodeAll := 0
	//newStoreEveRate:= float64(0)

	storageNruMap, _ := redis.IntMap(conn.Do("ZRANGE", "EVENT-MEDIA-SERVER-DEVICE-NUM", 0, -1, "withscores"))

	if storageNruMap != nil {
		for k, _ := range storageNruMap {
			storageInfo, _ := redis.String(conn.Do("HGET", "EVENT-MEDIA-REPORT-"+k[1:len(k)-1], preOneMin))
			storageAddrIp, _ := redis.String(conn.Do("GET", "EVENT-MEDIA-SERVER-"+k[1:len(k)-1]))
			if storageAddrIp != "" {
				pos := strings.Index(storageAddrIp, "baseUrl")
				end := strings.Index(storageAddrIp[pos+9:], ",")
				storageAddrIp = storageAddrIp[pos+17 : end+pos+3]
			}
			if storageInfo != "" && strings.Contains(storageInfo, "storageTotal") {
				//log.Println(storageInfo)
				//应存储总量
				pos := strings.Index(storageInfo, "storageTotal")
				end := strings.Index(storageInfo[pos+15:], ",")

				//log.Println(storageInfo[pos+15 : end+pos+15])
				storageAll, _ = strconv.Atoi(storageInfo[pos+15 : end+pos+15])
				//nodeAll += storageAll

				//成功量
				pos = strings.Index(storageInfo, "storageSucc")
				end = strings.Index(storageInfo[pos+14:], ",")

				//log.Println(storageInfo[pos+14 : end+pos+14])
				storageSucc, _ = strconv.Atoi(storageInfo[pos+14 : end+pos+14])
				//nodeSucc += storageSucc

				//存储积压
				pos = strings.Index(storageInfo, "storageUploading")

				//log.Println(storageInfo[pos+19 : pos+20])
				storageStack, _ = strconv.Atoi(storageInfo[pos+19 : pos+20])
			}

			httpOpentsdbDo("video.home.load.ipc", storageAll, storageAddrIp, "event", "allst")
			httpOpentsdbDo("video.home.load.ipc", storageSucc, storageAddrIp, "event", "successst")
			httpOpentsdbDo("video.home.load.ipc", storageStack, storageAddrIp, "event", "overstack")

		}
		//if  nodeSucc!= 0 && nodeAll != 0 {
		//	newStoreEveRate = float64(nodeSucc*100)/float64(nodeAll)
		//}
		//httpOpentsdbDoFloat("video.home.ratio.compute",newStoreEveRate,"none","event","successmake")
	}
}

//文件堆积数据采集
func pushOverstack(conn redis.Conn) {
	overStack, err := redis.Int(conn.Do("LLEN", "EDGE#STORAGE#FILE#LIST"))
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Printf("push overstack num is: %d", overStack)
	}
	httpOpentsdbDo("video.home.load.store", overStack, "none", "none", "upload")
}

//设备试图上传hbase
func hbasePost(deviceList string) {

	client := &http.Client{Timeout: 20 * time.Second}
	//log.Println(deviceList)

	req, err := http.NewRequest("POST", "http://172.20.3.23:8282/api/videocloud/report/devices", bytes.NewBuffer([]byte(deviceList)))

	if err != nil {
		log.Println("can not connect to hbase.", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("can not post to hbase.", err)
		return
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	urlRes := PostResponse{}
	jsonErr := json.Unmarshal(body, &urlRes)
	if jsonErr != nil {
		log.Printf("json 解析失败, %s ", string(body))
		return
	}
	errorCode := urlRes.ErrorCode
	errorMsg := urlRes.ErrorMsg

	if errorCode != "0" {
		log.Println(errorMsg)
	}
	//log.Printf("hbase return: %s ", string(body))

}

//设备数据上传mysql
func mysqlPost(deviceList string) {

	client := &http.Client{Timeout: 20 * time.Second}
	//log.Println(deviceList)

	req, err := http.NewRequest("POST", "http://apm-support-platform-provider.internal/apm/sp/deviceview/logs",
		bytes.NewBuffer([]byte(deviceList)))

	if err != nil {
		log.Println("can not connect to hbase.", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("can not post to hbase.", err)
		return
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	urlRes := MysqlPostResponse{}
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
	log.Printf("apm server return: %s ", string(body))

}

func main() {

	flag.Parse()
	today = time.Now().Format("2006-01-02")
	conn, err := redis.Dial("tcp", *redisHost,
		redis.DialDatabase(*database),
		redis.DialPassword(*redisPassword))
	if err != nil {
		log.Println("redis get error:", err)
		return
	}
	//redisPool = &redis.Pool{
	//
	//	MaxIdle:     1,
	//	MaxActive:   2000,
	//	IdleTimeout: 20 * time.Second,
	//	Dial: func() (redis.Conn, error) {
	//		c, err := redis.Dial("tcp",*redisHost,
	//			redis.DialDatabase(*database),
	//			redis.DialPassword(*redisPassword))
	//		if err != nil {
	//			log.Println("redis get error:", err)
	//			return nil, err
	//		}
	//		return c, nil
	//	},
	//}

	redisDataCol(conn)

	err = conn.Close()
	if err != nil {
		log.Printf("close redis connect err : %s", err.Error())
	}
	log.Println("collect over")
}
