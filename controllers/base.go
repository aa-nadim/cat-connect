package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

const CAT_API_BASE = "https://api.thecatapi.com/v1"

func (c *BaseController) Prepare() {
	c.Layout = "layout.html"
}
