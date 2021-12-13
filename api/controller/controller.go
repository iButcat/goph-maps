package controller

import (
	"log"

	"goph-maps/service"
	"goph-maps/utils"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetShortestPath(ctx *gin.Context)
	DisplayPointandLineString(ctx *gin.Context)
}

type controller struct {
	service service.Service
}

func NewController(service service.Service) Controller {
	return &controller{
		service: service,
	}
}

func (c *controller) GetShortestPath(ctx *gin.Context) {
	var getShortestPathRequest struct {
		start  float64 `json:"start"`
		finish float64 `json:"finish"`
	}

	if err := ctx.BindJSON(&getShortestPathRequest); err != nil {
		log.Println(err)
		return
	}
}

func (c *controller) DisplayPointandLineString(ctx *gin.Context) {
	graph := *utils.Graph
	ctx.JSON(200, gin.H{"graph": graph.Vertices})
}
