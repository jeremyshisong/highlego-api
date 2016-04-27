package main

import (
	_ "github.com/shawncode/highlego-api/docs"
	_ "github.com/shawncode/highlego-api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shawncode/highlego-api/task"
	"net/url"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "dev:199105@tcp(123.56.227.116:3306)/highlego?charset=utf8&loc="+url.QueryEscape("Asia/Shanghai")+"")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	go task.Task()
	beego.Run()
}

