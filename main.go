// main.go
package main

import (
	_ "cat-connect/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.BConfig.WebConfig.ViewsPath = "views"
	web.BConfig.WebConfig.StaticDir["/static"] = "static"
	web.BConfig.WebConfig.Session.SessionOn = true
	web.Run()
}
