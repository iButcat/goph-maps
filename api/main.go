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
	"goph-maps/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iButcat/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	utils.GeoJsonToStruct("mets.geojson")

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

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/test", controller.GetShortestPath)
	router.GET("/hello", controller.DisplayPointandLineString)

	errs := make(chan error, 1)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		log.Println("Starting server...")
		errs <- http.ListenAndServe(":8080", router)
	}()

	log.Println("exit", <-errs)
}
