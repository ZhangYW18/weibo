package controllers

import (
	"fmt"
	"net/http"

	"github.com/VampireWeekend/weibo/models"
	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

func LoginGet(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get(SESSION_KEY) != nil {
		c.Redirect(http.StatusSeeOther, "/index")
	}
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "微博登录",
	})
}

func LoginPost(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username)
	var user models.User
	user = models.FindUserByName(username)
	//	fmt.Println(user)

	var flag bool
	if user.Username == "" {
		flag = false
	} else if user.Password == password {
		flag = true
	} else {
		flag = false
	}
	//	fmt.Println(flag)

	//保存用户信息
	if flag {
		session := sessions.Default(c)
		v := session.Get(SESSION_KEY)
		if v == nil {
			session.Set(SESSION_KEY, user.Userid)
			session.Save()
		}
		fmt.Println("session:")
		fmt.Println(session.Get(SESSION_KEY))
	}

	c.JSON(http.StatusOK, gin.H{
		"flag": flag,
	})
}
