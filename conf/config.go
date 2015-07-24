package conf

import (
	"os"
)

var (
	config *Config
)

type Config struct {
	Redis *Redis
	Mysql *Mysql
	Mongo *Mongo
}

// redis
type Redis struct {
	Host     string
	Port     string
	Password string
}

// mysql
type Mysql struct {
	TcpConn  string
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// mongo
type Mongo struct {
	Addr     string
	Host     string
	Port     string
	User     string
	Password string
	Instance string
}

func init() {
	if config == nil {
		config = &Config{
			Redis: &Redis{
				Host:     os.Getenv("REDIS_PORT_6379_TCP_ADDR"),
				Port:     os.Getenv("REDIS_PORT_6379_TCP_PORT"),
				Password: os.Getenv("REDIS_PASSWORD"),
			},
			Mysql: &Mysql{
				TcpConn:  os.Getenv("MYSQL_PORT_3306_TCP"),
				Host:     os.Getenv("MYSQL_PORT_3306_TCP_ADDR"),
				Port:     os.Getenv("MYSQL_PORT_3306_TCP_PORT"),
				User:     os.Getenv("MYSQL_USERNAME"),
				Password: os.Getenv("MYSQL_PASSWORD"),
			},
			Mongo: &Mongo{
				Addr:     os.Getenv("MONGODB_PORT_27017_TCP"),
				Host:     os.Getenv("MONGODB_PORT_27017_TCP_ADDR"),
				Port:     os.Getenv("MONGODB_PORT_27017_TCP_PORT"),
				User:     os.Getenv("MONGODB_USERNAME"),
				Password: os.Getenv("MONGODB_PASSWORD"),
				Instance: os.Getenv("MONGODB_INSTANCE_NAME"),
			},
		}
	}
}

func Conf() *Config {
	return config
}
