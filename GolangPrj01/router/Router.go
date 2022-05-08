package router

import (
	"GolangPrj01/controller"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	engine := gin.Default()
	engine.GET("/article/getone/:id", controller.ArticleController{}.GetOne)
	engine.GET("/artivle/list", controller.ArticleController{}.GetList)
	return engine
}
