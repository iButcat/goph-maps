package controller

import (
	"log"

	"goph-maps/service"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetShortestPath(ctx *gin.Context)
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
