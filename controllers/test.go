package controllers

import (
	"github.com/ckeyer/dbc/conf"
	"github.com/ckeyer/dbc/lib"
)

type TestController struct {
	Controller
}

func (u *TestController) Get() {
	log.Debug("%v\n%v\n", *conf.Conf().Redis, *conf.Conf().Mysql)
	r := lib.GetRedis()
	// r.s
	// err := r.Set("hi", []byte("hello world"))
	// r.Set("hi", []byte("val"))
	// if err != nil {
	// 	log.Error(err.Error())
	// }
	bs, err := r.Get("hi")
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Notice(string(bs))
		u.Ctx.WriteString("Redis Test Success :" + string(bs) + "\n")
	}
	db := lib.GetMysql()
	if err = db.Ping(); err != nil {
		log.Error(err.Error())
	} else {
		u.Ctx.WriteString("Mysql Test Success\n")
	}
	u.Ctx.WriteString("Hello world!!!")
}
