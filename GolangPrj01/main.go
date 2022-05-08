package main

import (
	"GolangPrj01/gobal"
	"GolangPrj01/router"
	"github.com/spf13/viper"
	"os"
	"xorm.io/xorm"
)

func main() {

	InitConfig()
	db := gobal.SetupDBLink()
	defer func(db *xorm.Engine) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	engine := router.Router()
	err := engine.Run(":8081")
	if err != nil {
		return
	}

}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "\\config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}

}
