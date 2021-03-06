package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/wcharczuk/go-chart"
)

var listData []float64

func main() {

	listData = getData()
	go render()

	// Loop, get data and check every 5 minutes
	for i := 1; i < 1000; i++ {
		time.Sleep(5 * time.Minute)
		fmt.Println("Check at time ", i)
		listData = getData()
		checkData(listData)
	}
}

// Make chart available at localhost:8080
func render() {
	http.HandleFunc("/", drawChart)
	http.ListenAndServe(":8080", nil)
}

//Draw chart with retrieved data
func drawChart(res http.ResponseWriter, req *http.Request) {

	// create slice [1..60] for the x-axis of the chart

	arrayTime := []float64{}
	for i := 1; i <= 61; i++ {
		arrayTime = append(arrayTime, float64(i))
	}

	graph := chart.Chart{

		YAxis: chart.YAxis{
			Name:      "Bitcoin Prices in the past 30 minutes",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: arrayTime,
				YValues: listData,
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}
