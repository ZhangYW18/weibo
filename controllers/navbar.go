package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/ZhangYW18/weibo/models"
)

func NavbarGet(c *gin.Context) {
	session := sessions.Default(c)
	userid := session.Get(SESSION_KEY).(int)
	user := models.FindUserByID(userid)
	c.HTML(http.StatusOK, "navbar.html", gin.H{
		"user": user,
	})
}
