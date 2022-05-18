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
	cid := c.Query("id")
	id, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	data, err := Models.GetAllChapter(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, data)
}
