package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type WeatherInfoJson struct {
	Weatherinfo WeatherinfoObject
}

type WeatherinfoObject struct {
	City string
	Temp string
	WD   string
	WS   string
	SD   string
	Time string
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	resp, err := http.Get("http://www.weather.com.cn/data/sk/101280101.html")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	input, err := ioutil.ReadAll(resp.Body)

	var jsonWeather WeatherInfoJson
	json.Unmarshal(input, &jsonWeather)

	log.Println("城市:", jsonWeather.Weatherinfo.City)
	log.Println("温度:", jsonWeather.Weatherinfo.Temp, "℃")
	log.Println("风向:", jsonWeather.Weatherinfo.WD)
	log.Println("风力:", jsonWeather.Weatherinfo.WS)
	log.Println("湿度:", jsonWeather.Weatherinfo.SD)
	log.Println("时间:", jsonWeather.Weatherinfo.Time)
}

