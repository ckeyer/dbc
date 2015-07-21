package lib

import (
	"github.com/ckeyer/dbc/conf"
	"github.com/hoisie/redis"
)

var (
	redis_cli  *redis.Client
	redis_conf = conf.Conf().Redis
)

func GetRedis() *redis.Client {
	if redis_cli == nil {
		redis_cli.Addr = redis_conf.Host + ":" + redis_conf.Port
		err := redis_cli.Auth(redis_conf.Password)
		if err != nil {
			log.Error(err.Error())
		}
	}
	return redis_cli
}
