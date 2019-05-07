package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/VampireWeekend/weibo/controllers"
	"github.com/VampireWeekend/weibo/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := gin.Default()
	router.Static("/static", "/home/jcole/go/src/github.com/VampireWeekend/weibo/static")
	router.LoadHTMLGlob("/home/jcole/go/src/github.com/VampireWeekend/weibo/views/**/*")

	router.GET("/index/:username", func(c *gin.Context) {
		username := c.Param("username")
		user := models.FindUserByName(username)
		user.Password = "***"
		weibo := controllers.GetAllIndexWeibo(user.Userid)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"userid":   user.Userid,
			"username": username,
			"user":     user,
			"weibo":    weibo,
		})
	})

	router.GET("/my/:username", func(c *gin.Context) {
		username := c.Param("username")
		user := models.FindUserByName(username)
		user.Password = "***"
		weibo := controllers.GetAllWeibo(user.Userid)
		c.HTML(http.StatusOK, "my.html", gin.H{
			"userid":   user.Userid,
			"username": username,
			"user":     user,
			"weibo":    weibo,
		})
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "微博登录",
		})
	})

	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "微博注册",
		})
	})

	router.GET("/searchUser", func(c *gin.Context) {
		showtype := c.Query("showtype")
		userid, _ := strconv.Atoi(c.Query("userid"))
		user := models.FindUserByID(userid)
		user.Password = "***"
		if showtype == "follow" {
			result := controllers.SearchFollowUser(userid)
			c.HTML(http.StatusOK, "searchPeople.html", gin.H{
				"title":  user.Username + "关注的用户",
				"result": result,
				"user":   user,
			})
		} else if showtype == "followed" {
			result := controllers.SearchFollowedUser(userid)
			c.HTML(http.StatusOK, "searchPeople.html", gin.H{
				"title":  "关注" + user.Username + "的用户",
				"result": result,
				"user":   user,
			})
		} else if showtype == "search" {
			name := c.Query("searchName")
			result := controllers.SearchUser(name, userid)
			c.HTML(http.StatusOK, "searchPeople.html", gin.H{
				"title":  "查询用户",
				"result": result,
				"user":   user,
			})
		}
	})

	router.GET("/navbar.html", func(c *gin.Context) {
		userid, _ := strconv.Atoi(c.Query("userid"))
		user := models.FindUserByID(userid)
		c.HTML(http.StatusOK, "navbar.html", gin.H{
			"user": user,
		})
	})

	db, err := models.InitDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	db.SingularTable(true)

	router.POST("/loginpost", controllers.LoginPost)
	router.POST("/registerpost", controllers.RegisterPost)
	router.POST("/weibopost", controllers.WeiboPost)
	router.POST("/commentpost", controllers.CommentPost)
	router.POST("/follow", controllers.Follow)
	router.POST("/unfollow", controllers.Unfollow)
	router.GET("/countfollow", controllers.CountFollow)

	router.Run(":8087")

	defer db.Close()
}
