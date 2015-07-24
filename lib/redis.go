package lib

import (
	"fmt"
	"github.com/ckeyer/dbc/conf"
	"github.com/hoisie/redis"
)

var (
	redis_cli  redis.Client
	redis_conf = conf.Conf().Redis
)

func GetRedis() *redis.Client {
	return &redis_cli
}
func init() {
	var cli redis.Client
	cli.Addr = "172.17.42.1:6379"
	cli.Auth("abc")

	bs, err := cli.Get("hi")
	if err != nil {
		log.Error("HIHIHIHIHIH:::  ", err.Error())
	} else {
		log.Notice("HIHIHIHIHIH:::  ", string(bs))
	}
	return
	connstr := fmt.Sprintf("%s:%s", redis_conf.Host, redis_conf.Port)
	log.Debug(connstr)
	redis_cli.Addr = "172.17.42.1:6379"
	if redis_conf.Password != "" {
		// log.Debug(redis_conf.Passwor d)
		// err := redis_cli.Auth(redis_conf.Password)
		// if err != nil {
		// 	log.Error(err.Error())
		// }
		err := redis_cli.Auth("abc")
		if err != nil {
			log.Error(err.Error())
		} else {
			log.Notice("Connect Success")
		}
		bs, err := redis_cli.Get("hi")
		if err != nil {
			log.Error("HIHIHIHIHIH:::  ", err.Error())
		} else {
			log.Notice("HIHIHIHIHIH:::  ", string(bs))
		}
	}
}
