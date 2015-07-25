package mongo

import (
	"fmt"
	"github.com/ckeyer/dbc/conf"
	mgo "gopkg.in/mgo.v2"
)

var (
	db       *mgo.Database
	mgo_conf = conf.Conf().Mongo
)

func GetMgoCollection(table string) *mgo.Collection {
	if db == nil {
		_, err := GetMongo()
		if err != nil {
			return nil
		}
	}
	return db.C(table)
}

func GetMongo() (*mgo.Database, error) {
	if db == nil {
		connstr := fmt.Sprintf("%s:%s/%s", mgo_conf.Host, mgo_conf.Port, mgo_conf.Instance)
		session, err := mgo.Dial(connstr)
		if err != nil {
			return nil, err
		}
		db = session.DB(mgo_conf.Instance)
	}
	return db, nil
}
