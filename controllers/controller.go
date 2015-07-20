package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ckeyer/dbc/lib"
)

var (
	log = lib.GetLogger()
)

type Controller struct {
	beego.Controller
}

func (c *Controller) Prepare() {
	log.Notice("Start")
	// c.Ctx.WriteString("hello world")
}
