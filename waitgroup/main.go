package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

/**
 *
 */
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
	result := []WeatherInfo{}
	// 声明一个 waitGroup 并增加两个等待信号
	var wg sync.WaitGroup
	wg.Add(2)

	// 使用协程请求北京天气情况
	go func() {
		weatherInfo := GetWeatherByCity("101010100")
		result = append(result, weatherInfo)
		// 告诉 wg 我已经完成任务
		wg.Done()
	}()

	// 使用协程请求成都天气情况
	go func() {
		weatherInfo := GetWeatherByCity("101270101")
		result = append(result, weatherInfo)
		// 告诉 wg 我已经完成任务
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("%+v\n", result)
}
