package main

import (
	_ "github.com/ckeyer/dbc_test/docs"
	_ "github.com/ckeyer/dbc_test/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
