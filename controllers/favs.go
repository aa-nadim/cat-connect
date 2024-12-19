package controllers

type FavsController struct {
	BaseController
}

func (c *FavsController) Get() {
	c.TplName = "favs.html"
}
