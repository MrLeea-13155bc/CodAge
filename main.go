package main

import (
	"QuestionSearch/MiddleWares"
	"QuestionSearch/Routers"
	"github.com/gin-contrib/cors"
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
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080", "http://localhost:8081"}
	config.AllowCredentials = true
	router.Use(cors.New(config))
	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/register", Routers.CreateUser)
			user.POST("/login", Routers.Login)
			user.PUT("/Info", MiddleWares.Auth(), Routers.UpdateUserInfo)
			user.GET("/Info", MiddleWares.Auth(), Routers.GetUserInfo)
			user.GET("/my", MiddleWares.Auth(), Routers.GetShowInfo)
		}
		question := api.Group("/question", MiddleWares.Auth())
		{
			question.GET("/get", Routers.GetQuestions)
		}
		subject := api.Group("/subject", MiddleWares.Auth())
		{
			subject.GET("/getAllSubject", Routers.GetAllSubject)
			subject.GET("/getAllChaptersById", Routers.GetAllChaptersById)
		}
	}
	router.RunTLS(":8081", "key/wonend.pem", "key/wonend.key")
}
