package main

import (
	_ "cat-connect/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Run()
}
