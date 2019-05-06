package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/VampireWeekend/weibo/models"
	"github.com/gin-gonic/gin"
)

// 获取一个用户自己的微博
func GetAllWeibo(userid int) []models.Weibos {
	var weibo []models.Weibos
	weibo = models.FindAllWeiboByUserID(userid)
	return weibo
}

// 获取一个用户关注的所有微博和自己的微博
func GetAllIndexWeibo(userid int) []models.Weibos {
	var weibo []models.Weibos
	weibo = models.FindAllIndexWeibo(userid)
	//	fmt.Println(weibo)
	//	fmt.Println(query.Username)
	return weibo
}

// 发微博
func WeiboPost(c *gin.Context) {
	var weibo models.Weibo
	weibo.Text = c.PostForm("text")
	weibo.Userid, _ = strconv.Atoi(c.PostForm("userid"))
	weibo.Username = c.PostForm("username")
	fmt.Println(c.PostForm("userid"))
	fmt.Println(c.PostForm("text"))
	//	weibo.CreatedAt = time.Now()
	weibo.Like = 0
	weibo.CommentCount = 0
	flag := models.AddWeibo(weibo)

	c.JSON(http.StatusOK, gin.H{
		"flag": flag,
	})
}
