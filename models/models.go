package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User xxx
type User struct {
	Userid   int `gorm:"primary_key;AUTO_INCREMENT"`
	Username string
	Password string
	Sex      int
	Age      int
}

type Comment struct {
	Commentid   int `gorm:"primary_key;AUTO_INCREMENT"`
	Weiboid     int
	Userid      int
	Username    string
	CreatedAt   time.Time
	CommentText string
}

type Weibo struct {
	Weiboid      int `gorm:"primary_key;AUTO_INCREMENT"`
	Userid       int
	Username     string
	CreatedAt    time.Time
	Text         string
	Like         int
	CommentCount int
}

type Weibos struct {
	Weibo
	Comment []Comment
}

type Follow struct {
	Userid   int `gorm:"primary_key"`
	Followid int `gorm:"primary_key"`
}

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {

	db, err := gorm.Open("mysql", "debian-sys-maint:#VictorOladipo#@tcp(127.0.0.1:3306)/weibo?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		db.SingularTable(true)
		db.AutoMigrate(&User{}, &Weibo{}, &Comment{}, &Follow{})
		DB = db
		return db, err
	}
	return nil, err
}

func FindUserByID(userid int) User {
	var query User
	DB.Where("userid=?", userid).First(&query)
	return query
}

func FindUserByName(username string) User {
	var query User
	DB.Where("username=?", username).First(&query)
	return query
}

func AddUser(user User) bool {
	var query, nowUser User
	DB.Last(&query)
	var maxUserID int
	maxUserID = query.Userid
	nowUser = user
	nowUser.Userid = maxUserID + 1
	fmt.Println(nowUser)
	if err := DB.Create(&nowUser).Error; err != nil {
		fmt.Println("Insert Error")
		fmt.Println(err)
		return false
	}
	return true
}

//根据用户自己，查找微博
func FindAllWeiboByUserID(userid int) []Weibos {
	var query []Weibo
	var ans []Weibos
	DB.Where("userid=?", userid).Order("created_at desc").Find(&query)
	for _, weibo := range query {
		var this Weibos
		this.Weibo = weibo
		DB.Where("weiboid=?", weibo.Weiboid).Order("created_at desc").Find(&this.Comment)
		ans = append(ans, this)
	}
	return ans
}

//根据用户自己和关注，查找微博
func FindAllIndexWeibo(userid int) []Weibos {
	var query []Weibo
	var ans []Weibos
	DB.Raw("SELECT * FROM weibo WHERE userid = ?  or userid in "+
		"(SELECT followid FROM follow WHERE userid = ?) order by created_at desc", userid, userid).Scan(&query)
	for _, weibo := range query {
		var this Weibos
		this.Weibo = weibo
		DB.Where("weiboid=?", weibo.Weiboid).Order("created_at desc").Find(&this.Comment)
		ans = append(ans, this)
	}
	return ans
}

//添加微博
func AddWeibo(weibo Weibo) bool {
	fmt.Println(weibo)
	if err := DB.Create(&weibo).Error; err != nil {
		fmt.Println("Insert Error")
		fmt.Println(err)
		return false
	}
	return true
}

//添加微博评论
func AddComment(comment Comment) bool {
	fmt.Println(comment)
	if err := DB.Create(&comment).Error; err != nil {
		return false
	}
	//增加对应微博的评论数量
	err := DB.Table("weibo").Where("weiboid = ?", comment.Weiboid).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error
	if err != nil {
		return false
	}
	return true
}
