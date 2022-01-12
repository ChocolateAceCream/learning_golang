package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const Opentsdb = "http://172.200.96.82:4242"

type Tags struct {
	Area   string `json:"area"`
	IP     string `json:"ip"`
	Combo  string `json:"combo"`
	Target string `json:"target"`
}

type TsdbFloat struct {
	Metric    string  `json:"metric"`
	Tags      Tags    `json:"tags"`
	Value     float64 `json:"value"`
	Timestamp int64   `json:"timestamp"`
}

//上传浮点数型数据到opentsdb
func httpOpenbsdDoFloat(node, metric string, value float64, ip, combo, target string) {
	tag := Tags{
		Area:   node,
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
	req, err := http.NewRequest("POST", Opentsdb+"/api/put?details", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("can not connect to opentsdb.")
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("can not connect to opentsdb.")
	}

	log.Println(string(body) + metric)

}

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func main() {
	tiker := time.NewTicker(time.Minute)
	nodeList := [6]string{"73302001", "73302002", "33010001", "990102", "yyy-yyy1", "99sip-down"}
	ipList := [6]string{"172.53.30.12", "172.53.30.13", "172.53.30.14", "172.53.30.15", "172.53.30.16", "172.53.30.17"}

	for {
		fmt.Println("------------------------", <-tiker.C)
		for _, value := range nodeList {
			httpOpenbsdDoFloat(value, "video.home.load.ipc", randomFloat(100, 200), "none", "none", "online")    //在线数
			httpOpenbsdDoFloat(value, "video.home.ratio.compute", randomFloat(10, 90), "none", "none", "online") //在线率
			httpOpenbsdDoFloat(value, "video.home.ratio.compute", randomFloat(10, 90), "none", "none", "stream") //直播成功率
			httpOpenbsdDoFloat(value, "video.home.load.ipc", randomFloat(200, 300), "none", "nru", "online")     //全天在线数量
			httpOpenbsdDoFloat(value, "video.home.load.ipc", randomFloat(400, 600), "none", "event", "online")   //事件在线数量
			httpOpenbsdDoFloat(value, "video.home.ratio.compute", randomFloat(10, 90), "none", "none", "nru")    //全天存储正常率
			httpOpenbsdDoFloat(value, "video.home.ratio.store", randomFloat(10, 90), "none", "event", "success") //事件存储成功率

			//新增设备
			httpOpenbsdDoFloat(value, "video.home.device.add", randomFloat(4, 20), "none", "none", "home")    //天翼看家
			httpOpenbsdDoFloat(value, "video.home.device.add", randomFloat(1, 20), "none", "none", "gb")      //国标
			httpOpenbsdDoFloat(value, "video.home.device.add", randomFloat(2, 20), "none", "none", "21cn")    //集团
			httpOpenbsdDoFloat(value, "video.home.device.add", randomFloat(3, 20), "none", "none", "cascade") //联级

			for _, ip := range ipList {
				httpOpenbsdDoFloat(value, "video.home.load.ipc", randomFloat(10, 440), ip, "none", "stream") //直播服务负载
				httpOpenbsdDoFloat(value, "video.home.load.ipc", randomFloat(10, 140), ip, "nru", "store")   //全天存储服务负载
				httpOpenbsdDoFloat(value, "video.home.load.ipc", randomFloat(10, 140), ip, "event", "store") //全天存储服务负载
			}
		}

	}
}
