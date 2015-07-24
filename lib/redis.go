package lib

import (
	"fmt"
	"github.com/ckeyer/dbc/conf"
	"github.com/go-redis/redis"
)

var (
	redis_cli  *redis.Client
	redis_conf = conf.Conf().Redis
)

func GetRedis() *redis.Client {
	if redis_cli == nil {
		connstr := fmt.Sprintf("%s:%s", redis_conf.Host, redis_conf.Port)
		redis_cli.Addr = connstr
		redis_cli = &redis.NewClient(&redis.Options{
			Addr:     connstr,
			Password: redis_conf.Password,
			DB:      0,
		})
		_, err := client.Ping().Result()
		if err!=nil{
			log.Error(err.Error())
		}
	}
	return redis_cli
}
