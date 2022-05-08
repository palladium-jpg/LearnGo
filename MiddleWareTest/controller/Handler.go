package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handlFunc(ctx *gin.Context) {
	fmt.Println("正常的开始捏")
	ctx.JSON(http.StatusOK, gin.H{
		"Kabuto": "Hyper",
	})
}
