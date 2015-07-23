package lib

import (
	"fmt"
	"github.com/ckeyer/dbc/conf"
	"github.com/hoisie/redis"
)

var (
	redis_cli  *redis.Client
	redis_conf = conf.Conf().Redis
)

func GetRedis() *redis.Client {
	if redis_cli == nil {
		redis_cli = &redis.Client{}
		connstr := fmt.Sprintf("%s:%s", redis_conf.Host, redis_conf.Port)
		log.Debug(connstr)
		redis_cli.Addr = connstr
		if redis_conf.Password != ""{
			log.Debug(redis_conf.Password)
			err := redis_cli.Auth(redis_conf.Password)
			if err != nil {
				log.Error(err.Error())
			}
			err = redis_cli.Auth("MNZiMQKXAbeOyq6")
			if err != nil {
				log.Error(err.Error())
			}
		}
	}
	return redis_cli
}
