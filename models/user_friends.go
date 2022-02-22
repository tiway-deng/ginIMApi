package models

import (
	"ginIMApi/packages/setting"
)

type UserFriends struct {
	Model
	User1       int       `json:"user1"`
	User2       int       `json:"user2"`
	User1Remark int       `json:"user1_remark"`
	User2Remark int       `json:"user2_remark"`
	Active      int       `json:"active"`
	Status      int       `json:"status"`
	AgreeTime   int       `json:"agree_time"`
}

//set table name
func (UserFriends) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "users_friends"
}

//return true if it is user's friend
func IsUserFriend(user1 int, user2 int) bool {
	var userFriend UserFriends
	isFriend := false
	db.Model(UserFriends{}).Select("id").Where("user1 = ? AND user2 = ?", user1, user2).Find(&userFriend)
	if userFriend.ID > 0 {
		isFriend = true
	}
	return isFriend
}

//get user friends
func GetUserFriends(userId interface{}, field []string) []UserFriends {
	var userFriends []UserFriends
	db.Model(UserFriends{}).Select(field).Where("user1 = ? OR user2 = ?", userId, userId).Scan(&userFriends)

	return userFriends
}

func AddUserFriend(user1 int, user2 int, remark string) UserFriends {
	var userFriend UserFriends
	db.Model(userFriend).Where("user1 = ? OR user2 = ?", user1, user2).Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})

	return userFriend
}
