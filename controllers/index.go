package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/ZhangYW18/weibo/models"
)

func IndexGet(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get(SESSION_KEY) == nil {
		c.Redirect(http.StatusSeeOther, "/login")
	}
	uid := session.Get(SESSION_KEY).(int)
	user := models.FindUserByID(uid)
	fmt.Println(uid)
	weibo := GetAllIndexWeibo(uid)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"userid":     uid,
		"username":   user.Username,
		"user":       user,
		"weibo":      weibo,
		"weibocount": models.FindWeiboCountByUserID(user.Userid),
	})
}
