package controller

import (
	"GolangPrj01/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type ArticleController struct {
}

func NewAC() ArticleController {
	return ArticleController{}
}

func (a ArticleController) GetOne(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println("id:" + id)

	articleId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.AbortWithStatus(400)
		fmt.Println(err.Error())
	}
	articleOne, er := service.GetOneArticle(articleId)
	if er != nil {
		ctx.AbortWithStatus(404)
		log.Fatal(er)
	} else {
		ctx.JSON(http.StatusOK, articleOne)
	}
}

func (a ArticleController) GetList(ctx *gin.Context) {

}
