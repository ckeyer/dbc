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
}

type Redis struct {
	Host     string
	Port     string
	Password string
}

type Mysql struct {
	TcpConn  string
	Host     string
	Port     string
	User     string
	Password string
	Database string
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
		}
	}
}

func Conf() *Config {
	return config
}
