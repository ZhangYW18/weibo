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

type SearchUser struct {
	User
	Relation int
	// Relation 0:无关系 1:被关注 2:关注 3:互相关注
}

var DB *gorm.DB

//连接数据库
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

//根据ID查询用户
func FindUserByName(userid string) User {
	var query User
	DB.Where("username=?", userid).First(&query)
	return query
}

//根据ID查询用户
func FindUserByID(userid int) User {
	var query User
	DB.Where("userid=?", userid).First(&query)
	return query
}

//添加用户
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

// 查找一个用户和其他用户的关系
// Relation属性取值 0:无关系 1:被关注 2:关注 3:互相关注
func FindRelation(userid int, users []User) []SearchUser {
	var result []SearchUser
	for _, user := range users {
		var this SearchUser
		this.User = user
		this.Relation = 0
		var count int
		DB.Table("follow").Where("userid = ? and followid = ?", userid, user.Userid).Count(&count)
		if count > 0 {
			this.Relation = this.Relation + 2
		}
		DB.Table("follow").Where("userid = ? and followid = ?", user.Userid, userid).Count(&count)
		if count > 0 {
			this.Relation = this.Relation + 1
		}
		result = append(result, this)
	}
	fmt.Println(result)
	return result
}

//根据用户名查找用户和关系
func SearchUserByName(name string, userid int) []SearchUser {
	var users []User
	DB.Where("username LIKE ?", "%"+name+"%").Find(&users)
	return FindRelation(userid, users)
}

//查找我关注的用户
func SearchFollowUser(userid int) []SearchUser {
	var users []User
	DB.Raw("SELECT * FROM user WHERE userid in "+
		"(SELECT followid FROM follow WHERE userid = ?)", userid).Scan(&users)
	return FindRelation(userid, users)
}

//查找关注我的用户
func SearchFollowedUser(userid int) []SearchUser {
	var users []User
	DB.Raw("SELECT * FROM user WHERE userid in "+
		"(SELECT userid FROM follow WHERE followid = ?)", userid).Scan(&users)
	return FindRelation(userid, users)
}

//添加关注关系
func AddFollow(userid int, followid int) bool {
	if err := DB.Exec("INSERT INTO follow (`userid`, `followid`) VALUES (?, ?)", userid, followid).Error; err != nil {
		fmt.Println("Insert Error")
		fmt.Println(err)
		return false
	}
	return true
}

//删除关注关系
func DeleteFollow(userid int, followid int) bool {
	var follow Follow
	follow.Userid = userid
	follow.Followid = followid
	if err := DB.Delete(&follow).Error; err != nil {
		fmt.Println("Delete Error")
		fmt.Println(err)
		return false
	}
	return true
}

//查找关注和被关注总数
func CountFollow(userid int) (int, int) {
	var cf, cfd int
	DB.Table("follow").Where("userid = ?", userid).Count(&cf)
	DB.Table("follow").Where("followid = ?", userid).Count(&cfd)
	return cf, cfd
}
