package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeatherBody struct {
	WeatherInfo `json:"weatherinfo"`
}

type WeatherInfo struct {
	// 城市
	City string `json:"city"`
	// 最低温度
	Low string `json:"temp1"`
	// 最高温度
	High string `json:"temp2"`
}

// 获取天气情况的函数
func GetWeatherByCity(cityCode string) WeatherInfo {
	resp, err := http.Get("http://www.weather.com.cn/data/cityinfo/" + cityCode + ".html")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	weatherBody := WeatherBody{}

	json.Unmarshal(body, &weatherBody)

	return weatherBody.WeatherInfo
}

func main() {
	// 待宰的羔羊
	cityIDS := []string{"101010100", "101270101"}
	result := []WeatherInfo{}
	// 记录任务完成的计数器
	taskCount := 0
	// 创建一个channal, 用了接受完成信号
	taskChan := make(chan bool)
	for _, cityID := range cityIDS {
		// 此处是不是很熟悉？
		id := cityID
		go func(ch chan bool) {
			weatherInfo := GetWeatherByCity(id)
			result = append(result, weatherInfo)
			// 当任务完成，向channel内发送一个信号
			taskChan <- true
		}(taskChan)
	}

	for {
		// 在死循环种，这一步会阻塞，直到从channel中接收到数据
		<-taskChan
		// 每次收到信号，计数器加一，直到计数器等于任务个数，跳出循环
		taskCount++
		if taskCount == len(cityIDS) {
			// 全部完成
			fmt.Printf("%+v\n", result)
			break
		}
	}
}
