package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/VampireWeekend/weibo/models"

	"github.com/gin-gonic/gin"
)

func UserGet(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Param("userid"))
	var user models.User
	user = models.FindUserByID(userid)
	fmt.Println(user)
	//	fmt.Println(query.Username)
	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"title":    "微博",
	})
}
