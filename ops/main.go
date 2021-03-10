package main

import (
	"fmt"
	"log"

	_ "ops/controllers"
	_ "ops/models"
	_ "ops/routers"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=PRC&parseTime=true",
		beego.AppConfig.DefaultString("mysql::User", ""),
		beego.AppConfig.DefaultString("mysql::Password", ""),
		beego.AppConfig.DefaultString("mysql::Host", ""),
		beego.AppConfig.DefaultString("mysql::Port", ""),
		beego.AppConfig.DefaultString("mysql::DBName", ""),
	)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RunCommand()

	if db, err := orm.GetDB("default"); err != nil {
		log.Fatal(err)
	} else if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	beego.Run()
}
