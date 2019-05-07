package controllers

import (
	"github.com/VampireWeekend/weibo/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SearchUser(name string, userid int) []models.SearchUser {
	return models.SearchUserByName(name, userid)
}

func SearchFollowUser(userid int) []models.SearchUser {
	return models.SearchFollowUser(userid)
}

func SearchFollowedUser(userid int) []models.SearchUser {
	return models.SearchFollowedUser(userid)
}
