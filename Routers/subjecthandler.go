package Routers

import (
	"QuestionSearch/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func GetAllChaptersById(c *gin.Context) {
	sSectionId := c.Query("id")
	iUserId, _ := c.Get("userId")
	userId := iUserId.(int64)
	sectionId, err := strconv.ParseInt(sSectionId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	data, err := Models.GetAllChapter(sectionId, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, data)
}
