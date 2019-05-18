package controllers

import (
	"fmt"
	"net/http"

	"github.com/VampireWeekend/weibo/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func MyGet(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get(SESSION_KEY) == nil {
		c.Redirect(http.StatusSeeOther, "/login")
	}
	uid := session.Get(SESSION_KEY).(int)
	user := models.FindUserByID(uid)
	weibo := GetAllWeibo(user.Userid)
	fmt.Println(uid)
	c.HTML(http.StatusOK, "my.html", gin.H{
		"userid":     uid,
		"username":   user.Username,
		"user":       user,
		"weibo":      weibo,
		"weibocount": models.FindWeiboCountByUserID(user.Userid),
	})
}

func MyAlternativeGet(c *gin.Context) {
	username := c.Param("username")
	user := models.FindUserByName(username)
	user.Password = "***"
	weibo := GetAllWeibo(user.Userid)
	c.HTML(http.StatusOK, "my.html", gin.H{
		"userid":     user.Userid,
		"username":   username,
		"user":       user,
		"weibo":      weibo,
		"weibocount": models.FindWeiboCountByUserID(user.Userid),
	})
}
