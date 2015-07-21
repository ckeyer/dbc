package lib

import (
	"database/sql"
	"fmt"
	"github.com/ckeyer/dbc/conf"
	_ "github.com/go-sql-driver/mysql"
)

var (
	mysql_conf = conf.Conf().Mysql
)

// var mysql *mysql.
func GetMysql() *sql.DB {
	connstr := fmt.Sprintf("%s:%s@tcp(%s:%s)?charset=utf8",
		mysql_conf.User, mysql_conf.Password, mysql_conf.Host, mysql_conf.Port)
	db, err := sql.Open("mysql", connstr)
	if err != nil {
		log.Error(err.Error())
	}
	return db
}
