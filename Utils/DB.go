package Utils

import (
	"database/sql"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type dbInfo struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

var db *sql.DB

func init() {
	conf, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalln(err)
	}
	var mysqlInfo dbInfo
	conf.Section("mysql").StrictMapTo(&mysqlInfo)
	dsn := mysqlInfo.User + ":" + mysqlInfo.Password + "@tcp(" + mysqlInfo.Host + ":" + mysqlInfo.Port + ")/" + mysqlInfo.Database
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln("[Mysql] Connect Default! err:", err)
	}
	log.Println("[Mysql] Connect Success!")
}

func MDB() *sql.DB {
	return db
}
