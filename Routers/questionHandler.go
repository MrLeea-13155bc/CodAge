package Routers

import (
	"QuestionSearch/Models"
	"QuestionSearch/Utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetQuestions(c *gin.Context) {
	form := Utils.GetQuestionForm{}
	var err1, err2, err3 error
	sCategory := c.Query("chapter")
	sNum := c.Query("num")
	sType := c.Query("type")
	iUserId, _ := c.Get("userId")
	form.UserId = iUserId.(int64)
	form.SectionId, err1 = strconv.ParseInt(sCategory, 10, 64)
	form.Num, err2 = strconv.ParseInt(sNum, 10, 64)
	form.RequestType, err3 = strconv.ParseInt(sType, 10, 64)
	if err1 != nil || err2 != nil || err3 != nil {
		log.Println(err1, err2, err3)
		c.JSON(http.StatusBadRequest, gin.H{"message": "请求错误"})
		return
	}
	data, err := Models.GetQuestions(form)
	if err != nil {
		log.Println("[GetQuestions] err ", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, data)
}

func CheckAnswer(c *gin.Context) {
	iUserId, _ := c.Get("userId")
	userId := iUserId.(int64)
	var form []Utils.Answer
	c.Bind(&form)
	data, correct, total, err := Models.CheckAnswers(form, userId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data":    data,
		"correct": correct,
		"total":   total,
	})
}
