package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getData() []float64 {
	// Get data from API, return close price in the past 30 minutes
	var f interface{}
	listData := []float64{}

	// Request from API, get response
	resp, err := http.Get("https://min-api.cryptocompare.com/data/histominute?fsym=BTC&tsym=USD&limit=60")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Read content from response
	respContent, err := ioutil.ReadAll(resp.Body)

	// Parse respContent Json to interface f
	err = json.Unmarshal(respContent, &f)

	jsonData := f.(map[string]interface{})
	arrayData := jsonData["Data"].([]interface{})

	for _, object := range arrayData {
		x := object.(map[string]interface{})
		minutePrice := x["close"].(float64)
		listData = append(listData, minutePrice)

	}

	return listData
}

func checkData(listData []float64) {
	last10Data := listData[50:]
	last50Data := listData[:50]
	maxLast10 := max(last10Data)
	minLast10 := min(last10Data)
	differencelast10 := maxLast10 - minLast10

	maxLast50 := max(last50Data)
	minLast50 := min(last50Data)
	differencelast50 := maxLast50 - minLast50

	fmt.Println("Prices in the last 10 minutes:")
	fmt.Println("min: ", minLast10)
	fmt.Println("max: ", maxLast10)
	fmt.Println("Difference: ", differencelast10)

	fmt.Println("Prices from the last 60 to last 10 minutes:")
	fmt.Println("min: ", minLast50)
	fmt.Println("max: ", maxLast50)
	fmt.Println("Difference: ", differencelast50)

	if differencelast10 > differencelast50 {
		fmt.Println("It seems there is a dramatic change in Bitcoin Price")
	} else {
		fmt.Println("No big changes in prices")
	}

}

func min(array []float64) float64 {
	min := array[0]
	for _, j := range array {
		if j < min {
			min = j
		}
	}
	return min
}

func max(array []float64) float64 {
	max := array[0]
	for _, j := range array {
		if j > max {
			max = j
		}
	}
	return max
}
