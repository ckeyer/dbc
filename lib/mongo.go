package lib

import (
	"fmt"
	"github.com/ckeyer/dbc/conf"
	mgo "gopkg.in/mgo.v2"
)

var (
	db       *mgo.Database
	mgo_conf = conf.Conf().Mongo
)

func GetMongo(table string) *mgo.Collection {
	if db == nil {
		GetDB()
	}
	return db.C(table)
}

func GetDB() *mgo.Database {
	if db == nil {
		connstr := fmt.Sprintf("%s:%s/%s", mgo_conf.Host, mgo_conf.Port, mgo_conf.Instance)
		session, err := mgo.Dial(connstr)
		if err != nil {
			log.Error(err.Error())
			return nil
		}
		db = session.DB(mgo_conf.Instance)
	}
	return db
}
