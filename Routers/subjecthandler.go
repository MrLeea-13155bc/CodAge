package Routers

import (
	"QuestionSearch/Models"
	"QuestionSearch/Utils"
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

func SearchSubject(c *gin.Context) {
	subject := c.Query("name")
	if len([]rune(subject)) < 2 {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	userId, _ := c.Get("userId")
	result, err := Models.SearchSubject(subject, userId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, result)
}

func AddSubject(c *gin.Context) {
	userId, _ := c.Get("userId")
	var info Utils.Subject
	c.Bind(&info)
	if !Models.AddSubject(info.SubjectId, userId.(int64)) {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func InsertSubject(c *gin.Context) {
	var info Utils.Subject
	c.Bind(&info)
	if info.SubjectName == "" || !Models.InsertSubject(info) {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusCreated, nil)
}
func InsertSection(c *gin.Context) {
	var info Utils.SectionInfo
	c.Bind(&info)
	if info.SectionName == "" || !Models.InsertSection(info) {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusCreated, nil)
}
