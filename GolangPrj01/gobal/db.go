package gobal

import (
	"fmt"
	"github.com/spf13/viper"
	"xorm.io/xorm"
)

func SetupDBLink() *xorm.Engine {
	driverName := viper.GetString("datasource.driverName")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=true",
		username,
		password,
		database,
		charset)
	db, err := xorm.NewEngine(driverName, args)
	if err != nil {
		db.ShowSQL()
		panic("")
	}
	return db
}
