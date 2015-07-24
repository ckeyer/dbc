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
	
	err := r.Set("hi", "hello world", 0).Err()
	if err != nil {
		log.Error(err.Error())
	}
	bs, err := r.Get("hi").Result()
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
