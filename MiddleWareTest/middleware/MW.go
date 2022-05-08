package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func CookieMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("CookieMiddleware Start")
		user := context.Query("user")
		fmt.Printf("context.Request.Host: %v\n", strings.Split(context.Request.Host, ":")[0])
		context.SetCookie("user", user, 0, "", strings.Split(context.Request.Host, ":")[0], false, true)
		fmt.Println("CookieMiddleWare End")
	}
}

func TimeMiddleWare(ctx *gin.Context) {
	fmt.Println("TimeMiddleWare Start")
	start := time.Now()
	ctx.Next()
	Since := time.Since(start)
	fmt.Println(Since)
	fmt.Println("TimeMiddleWare End")
}

func LoginMiddleWare(ctx *gin.Context) {
	fmt.Println("LoginMiddleWare Start")
	loign := ctx.Query("login")
	if loign != "" {
		ctx.Next()
		fmt.Println("Hi," + loign + ",Welcome to Our daily day")
	} else {
		ctx.Abort()
		fmt.Println("未登录 gunna")
	}
	fmt.Println("End of LoginMiddleWare")
}
