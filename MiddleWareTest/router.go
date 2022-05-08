package main

import (
	"MiddleWareTest/Han"
	"MiddleWareTest/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func controll(engine *gin.Engine) {

	authorized := engine.Group("/admin")
	authorized.Use(middleware.TimeMiddleWare)
	{
		authorized.POST("/login", handlFunc)
		inneed := authorized.Group("indeed")
		inneed.GET("/analytics", Han.HandlFunc)
	}
	engine.GET("/mid", middleware.TimeMiddleWare, middleware.CookieMiddleware(), handlFunc)

	err := engine.Run(":8000")
	if err != nil {
		return
	}

}

func handlFunc(ctx *gin.Context) {
	fmt.Println("正常的开始捏")
	ctx.JSON(http.StatusOK, gin.H{
		"Kabuto": "Hyper",
	})
}
