package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"goph-maps/controller"
	"goph-maps/service"

	"github.com/gin-gonic/gin"
	"github.com/iButcat/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	/*
		graph := internal.Graph{}

		graph.Add("Tokyo")       // 0
		graph.Add("Beijing")     // 1
		graph.Add("Bangui")      // 2
		graph.Add("Berlin")      // 3
		graph.Add("Luxembourg")  // 4
		graph.Add("Mexico City") // 5
		graph.Add("Oslo")        // 6
		graph.Add("Bucharest")   // 7
		graph.Add("Singapore")   // 8
		graph.Add("Madrid")      // 9

		graph.AddEdge(0, 1, 100)
		graph.AddEdge(1, 2, 110)
		graph.AddEdge(2, 3, 120)
		graph.AddEdge(3, 4, 130)
		graph.AddEdge(4, 5, 140)
		graph.AddEdge(5, 6, 150)
		graph.AddEdge(6, 7, 160)
		graph.AddEdge(7, 8, 170)
		graph.AddEdge(8, 9, 180)

		path := graph.BFS(graph.GetVertexFromID(4), "Bucharest")
		fmt.Println("path: ", path)

	*/

	var db *gorm.DB
	{
		var err error
		db, err = gorm.Open(sqlite.Open("simple.db"), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	}

	repository := repository.NewRepo(db, log.Logger{})
	service := service.NewService(repository)
	controller := controller.NewController(service)

	var router = gin.New()
	router.GET("/test", controller.GetShortestPath)

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
		errs <- http.ListenAndServe(":8080", router)
	}()

	log.Println("exit", <-errs)
}
