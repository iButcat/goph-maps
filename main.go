package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var guangdongMapData = map[string]float64{
	"深圳市": float64(rand.Intn(150)),
	"广州市": float64(rand.Intn(150)),
	"湛江市": float64(rand.Intn(150)),
	"汕头市": float64(rand.Intn(150)),
	"东莞市": float64(rand.Intn(150)),
	"佛山市": float64(rand.Intn(150)),
	"云浮市": float64(rand.Intn(150)),
	"肇庆市": float64(rand.Intn(150)),
	"梅州市": float64(rand.Intn(150)),
}

func generateMapData(data map[string]float64) (items []opts.MapData) {
	items = make([]opts.MapData, 0)
	for k, v := range data {
		items = append(items, opts.MapData{Name: k, Value: v})
	}
	return
}

func mapRegion() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("广东")
	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Guangdong province",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			InRange:    &opts.VisualMapInRange{Color: []string{"#50a3ba", "#eac736", "#d94e5d"}},
		}),
	)

	mc.AddSeries("map", generateMapData(guangdongMapData))
	return mc
}

func main() {
	page := components.NewPage()
	page.AddCharts(
		mapRegion(),
	)

	f, err := os.Create("static/bar.html")
	if err != nil {
		panic(err)
	}
	page.Render(f)

	errs := make(chan error, 1)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		log.Println("Starting server...")
		fileServer := http.FileServer(http.Dir("./static"))
		http.Handle("/", fileServer)
		errs <- http.ListenAndServe(":8080", nil)
	}()

	log.Println("exit", <-errs)
}
