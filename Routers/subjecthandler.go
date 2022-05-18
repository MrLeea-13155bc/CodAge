package Routers

import (
	"QuestionSearch/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllSubject(c *gin.Context) {
	iUserId, _ := c.Get("userId")
	userId := iUserId.(int64)
	data, err := Models.GetAllSubject(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	} else {
		c.JSON(http.StatusOK, data)
	}
}
