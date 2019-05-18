package controllers

import (
	"net/http"

	"github.com/VampireWeekend/weibo/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func NavbarGet(c *gin.Context) {
	session := sessions.Default(c)
	userid := session.Get(SESSION_KEY).(int)
	user := models.FindUserByID(userid)
	c.HTML(http.StatusOK, "navbar.html", gin.H{
		"user": user,
	})
}
