package controllers

import (
	"net/http"
	"strconv"

	"github.com/VampireWeekend/weibo/models"

	"github.com/gin-gonic/gin"
)

func RegisterPost(c *gin.Context) {
	var user models.User
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	user.Sex, _ = strconv.Atoi(c.PostForm("sex"))
	user.Age, _ = strconv.Atoi(c.PostForm("age"))
	flag := models.AddUser(user)

	c.JSON(http.StatusOK, gin.H{
		"flag": flag,
	})
}
