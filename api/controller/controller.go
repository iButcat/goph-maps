package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

func GetShortestPath(ctx *gin.Context) {
	var getShortestPathRequest struct {
		start  float64 `json:"start"`
		finish float64 `json:"finish"`
	}

	if err := ctx.BindJSON(&getShortestPathRequest); err != nil {
		log.Println(err)
		return
	}
}
