package controllers

import (
	"fmt"
	"net/http"

	"github.com/VampireWeekend/weibo/models"

	"github.com/gin-gonic/gin"
)

func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username)
	var user models.User
	user = models.FindUserByName(username)
	fmt.Println(user)
	fmt.Println(password)
	var flag bool
	if user.Username == "" {
		flag = false
	} else if user.Password == password {
		flag = true
	} else {
		flag = false
	}
	fmt.Println(flag)

	c.JSON(http.StatusOK, gin.H{
		"flag": flag,
	})
}
