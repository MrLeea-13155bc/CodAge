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
	sCategory := c.Query("category")
	sNum := c.Query("num")
	sType := c.Query("type")
	form.Category, err1 = strconv.ParseInt(sCategory, 10, 64)
	form.Num, err2 = strconv.ParseInt(sNum, 10, 64)
	form.RequestType, err3 = strconv.ParseInt(sType, 10, 64)
	if err1 != nil || err2 != nil || err3 != nil {
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
