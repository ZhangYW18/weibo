package main

import (
	"fmt"

	"github.com/ZhangYW18/weibo/controllers"
	"github.com/ZhangYW18/weibo/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := gin.Default()
	router.Static("/static", "/home/jcole/go/src/github.com/ZhangYW18/weibo/static")
	router.LoadHTMLGlob("/home/jcole/go/src/github.com/ZhangYW18/weibo/views/**/*")

	db, err := models.InitDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	db.SingularTable(true)

	setSessions(router)

	router.GET("/", controllers.IndexGet)
	router.GET("/index", controllers.IndexGet)
	router.GET("/my", controllers.MyGet)
	router.GET("/my/:username", controllers.MyAlternativeGet)
	router.GET("/login", controllers.LoginGet)
	router.GET("/register", controllers.RegisterGet)
	router.GET("/searchUser", controllers.SearchGet)
	router.GET("/navbar.html", controllers.NavbarGet)

	router.POST("/loginpost", controllers.LoginPost)
	router.POST("/registerpost", controllers.RegisterPost)
	router.POST("/weibopost", controllers.WeiboPost)
	router.POST("/commentpost", controllers.CommentPost)
	router.POST("/follow", controllers.Follow)
	router.POST("/unfollow", controllers.Unfollow)
	router.GET("/countfollow", controllers.CountFollow)
	router.GET("/logout", controllers.LogoutGet)

	router.Run(":8079")

	defer db.Close()
}

//setSessions initializes sessions & csrf middlewares
func setSessions(router *gin.Engine) {
	//https://github.com/gin-gonic/contrib/tree/master/sessions
	//	store := cookie.NewStore([]byte(os.Getenv("SESSION_KEY")))
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"})
	router.Use(sessions.Sessions("gin-session", store))
}
