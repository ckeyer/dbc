package controllers

import (
	"github.com/ckeyer/dbc/lib/mongo"
	"github.com/ckeyer/dbc/lib/mysql"
	"github.com/ckeyer/dbc/lib/redis"
)

type TestController struct {
	Controller
}

func (u *TestController) Get() {

	r, err := redis.GetRedis()
	if err != nil {
		log.Error(err.Error())
	}

	err = r.Set("hi", "hello world", 0).Err()
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

	db, err := mysql.GetMysql()
	if err != nil {
		log.Error(err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Error(err.Error())
	} else {
		u.Ctx.WriteString("Mysql Test Success\n")
	}
	u.Ctx.WriteString("Hello world!!!")

	mgo, err := mongo.GetMongo()
	if err != nil {
		log.Error(err.Error())
	}

	if mgo != nil {
		log.Notice("Mgo is Not Nil")
	} else {
		log.Error("Mgo is Nil")
	}

}
