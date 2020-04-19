package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ZhangYW18/weibo/models"
)

// 发评论
func CommentPost(c *gin.Context) {
	var comment models.Comment
	comment.CommentText = c.PostForm("text")
	comment.Userid, _ = strconv.Atoi(c.PostForm("userid"))
	comment.Username = c.PostForm("username")
	comment.Weiboid, _ = strconv.Atoi(c.PostForm("weiboid"))
	flag := models.AddComment(comment)

	c.JSON(http.StatusOK, gin.H{
		"flag": flag,
	})
}
