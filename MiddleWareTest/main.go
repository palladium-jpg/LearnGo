package main

import "github.com/gin-gonic/gin"

func main() {

	//router := gin.New()

	//router.Use(gin.Logger(), gin.Recovery())
	router := gin.Default()
	//看源码发现Default函数等于上述两个函数
	//使用官方中间件
	controll(router)

}
