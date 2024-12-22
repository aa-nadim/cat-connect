package main

import (
	_ "cat-connect/routers"
	"io/ioutil"
	"path/filepath"

	beego "github.com/beego/beego/v2/server/web"
	context "github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func main() {
	// Enable CORS for all origins
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	// Serve static files
	beego.SetStaticPath("/static", "static")

	// Handle all non-API routes by serving index.html
	beego.Get("/*", func(ctx *context.Context) {
		if ctx.Input.URL() != "/favicon.ico" {
			indexPath := filepath.Join("views", "index.tpl")
			content, err := ioutil.ReadFile(indexPath)
			if err != nil {
				ctx.Output.SetStatus(500)
				ctx.WriteString("Error loading page")
				return
			}

			ctx.Output.Header("Content-Type", "text/html")
			ctx.Output.Header("Cache-Control", "no-cache, no-store, must-revalidate")
			ctx.Output.Header("Pragma", "no-cache")
			ctx.Output.Header("Expires", "0")
			ctx.WriteString(string(content))
		}
	})

	beego.Run()
}
