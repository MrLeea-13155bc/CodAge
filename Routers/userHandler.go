package Routers

import (
	"QuestionSearch/Models"
	"QuestionSearch/Utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var form = Utils.RegisterForm{}
	c.Bind(&form)
	if !Models.Register(form) {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func Login(c *gin.Context) {
	var form = Utils.RegisterForm{}
	c.Bind(&form)
	userId := Models.Login(form)
	if userId == -1 {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    Utils.CreateToken(userId),
		MaxAge:   Utils.MAXAGE,
		Path:     "/",
		Domain:   "",
		HttpOnly: false,
		SameSite: 4,
		Secure:   false,
	})
	c.JSON(http.StatusCreated, nil)
}

func UpdateUserInfo(c *gin.Context) {
	var form = Utils.UserInfo{}
	c.Bind(&form)
	iUserId, _ := c.Get("userId")
	form.UserId = iUserId.(int64)
	if !Models.UpdateUserInfo(form) {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func GetUserInfo(c *gin.Context) {
	iUserId, _ := c.Get("userId")
	userId := iUserId.(int64)
	data, err := Models.GetUserInfo(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetShowInfo(c *gin.Context) {
	iUserId, _ := c.Get("userId")
	userId := iUserId.(int64)
	data, err := Models.GetShowInfo(userId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	data.Subjects, err = Models.GetAllSubject(userId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, data)
}
