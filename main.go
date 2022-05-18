package main

import (
	"QuestionSearch/MiddleWares"
	"QuestionSearch/Routers"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"log"
)

var (
	port = ":8081"
)

func init() {
	conf, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalln(err)
	}
	portData, err := conf.Section("gin").GetKey("port")
	if err != nil {
		log.Fatalln(err)
	}
	port = ":" + portData.Value()
}

func main() {
	router := gin.Default()
	api := router.Group("/api")
	{
		question := api.Group("/question")
		{
			question.GET("/", Routers.GetQuestions)
		}
		subject := api.Group("/subject", MiddleWares.Auth())
		{
			subject.GET("/getAllSubject", Routers.GetAllSubject)
		}
	}
	router.Run(port)
}
