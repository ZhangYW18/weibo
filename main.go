package main

//import "github.com/gin-gonic/gin"

import (
	"fmt"
	"net/http"

	"github.com/VampireWeekend/weibo/controllers"
	"github.com/VampireWeekend/weibo/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := gin.Default()
	//	router.Static("/weibo", "./views")
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
		//根据完整文件名渲染模板，并传递参数
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "微博登录",
		})
	})

	router.GET("/register", func(c *gin.Context) {
		//根据完整文件名渲染模板，并传递参数
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "微博注册",
		})
	})

	router.GET("/navbar.html", func(c *gin.Context) {
		//根据完整文件名渲染模板，并传递参数
		c.HTML(http.StatusOK, "navbar.html", gin.H{})
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

	router.Run(":8091")

	defer db.Close()
}
