package lib

import (
	"fmt"
	"github.com/ckeyer/dbc/conf"
	redis "gopkg.in/redis.v3"
)

var (
	redis_cli  *redis.Client
	redis_conf = conf.Conf().Redis
)

func GetRedis() *redis.Client {
	return redis_cli
}
func init() {
	if redis_cli == nil {
		connstr := fmt.Sprintf("%s:%s", redis_conf.Host, redis_conf.Port)
		redis_cli = redis.NewClient(&redis.Options{
			Addr:     connstr,
			Password: redis_conf.Password,
			DB:      0,
		})
		_, err := redis_cli.Ping().Result()
		if err!=nil{
			log.Error(err.Error())
		}
	}
}
