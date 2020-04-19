package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/ZhangYW18/weibo/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SearchGet(c *gin.Context) {
	showtype := c.Query("showtype")
	session := sessions.Default(c)
	if session.Get(SESSION_KEY) == nil {
		c.Redirect(http.StatusSeeOther, "/login")
	}
	userid := session.Get(SESSION_KEY).(int)
	user := models.FindUserByID(userid)
	user.Password = "***"
	if showtype == "follow" {
		queryUID, _ := strconv.Atoi(c.Query("userid"))
		queryUser := models.FindUserByID(queryUID)
		result := SearchFollowUser(queryUID)
		c.HTML(http.StatusOK, "searchPeople.html", gin.H{
			"title":               queryUser.Username + "关注的用户",
			"queryFollowUsername": queryUser.Username,
			"result":              result,
			"user":                user,
		})
	} else if showtype == "followed" {
		queryUID, _ := strconv.Atoi(c.Query("userid"))
		queryUser := models.FindUserByID(queryUID)
		result := SearchFollowedUser(queryUID)
		c.HTML(http.StatusOK, "searchPeople.html", gin.H{
			"title":                 "关注" + queryUser.Username + "的用户",
			"queryFollowedUsername": queryUser.Username,
			"result":                result,
			"user":                  user,
		})
	} else if showtype == "search" {
		name := c.Query("searchName")
		result := SearchUser(name, userid)
		count := len(result)
		c.HTML(http.StatusOK, "searchPeople.html", gin.H{
			"title":      "查询用户",
			"result":     result,
			"user":       user,
			"countusers": count,
		})
	}
}

func SearchUser(name string, userid int) []models.SearchUser {
	return models.SearchUserByName(name, userid)
}

func SearchFollowUser(userid int) []models.SearchUser {
	return models.SearchFollowUser(userid)
}

func SearchFollowedUser(userid int) []models.SearchUser {
	return models.SearchFollowedUser(userid)
}
