package Utils

import (
	"database/sql"
	"github.com/go-ini/ini"
	"github.com/go-redis/redis"
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

type rdbInfo struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Password string `ini:"password"`
	DB       int    `ini:"db"`
}

var db *sql.DB
var rdb *redis.Client

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
	var redisInfo rdbInfo
	conf.Section("redis").MapTo(&redisInfo)
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisInfo.Host + ":" + redisInfo.Port,
		Password: redisInfo.Password,
		DB:       redisInfo.DB,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		log.Panicln("[Redis]连接失败", err)
	}
	log.Println("[Redis]链接成功")
}

func MDB() *sql.DB {
	return db
}

func RDB() *redis.Client {
	return rdb
}
