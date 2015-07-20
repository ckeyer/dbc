package main

import (
	_ "github.com/ckeyer/dbc/docs"
	_ "github.com/ckeyer/dbc/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
