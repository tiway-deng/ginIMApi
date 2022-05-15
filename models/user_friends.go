package models

import (
	"time"
)

type UsersFriends struct {
	Model
	User1       int       `json:"user1"`
	User2       int       `json:"user2"`
	User1Remark string    `json:"user1_remark"`
	User2Remark string    `json:"user2_remark"`
	Active      int       `json:"active"`
	Status      int       `json:"status"`
	AgreeTime   time.Time `json:"agree_time"`
}


//return true if it is user's friend
func IsUserFriend(user1 int, user2 int) bool {
	var userFriend UsersFriends
	isFriend := false
	userFriend = GetUserFriend(user1, user2)
	if userFriend.ID > 0 {
		isFriend = true
	}
	return isFriend
}

func GetUserFriend(user1 int, user2 int) UsersFriends {
	var userFriend UsersFriends
	db.Select("id").Where("user1 = ? AND user2 = ?", user1, user2).First(&userFriend)

	return userFriend
}

//get user friends
func GetUserFriends(userId interface{}, field []string) []UsersFriends {
	var userFriends []UsersFriends
	db.Model(UsersFriends{}).Select(field).Where("user1 = ? OR user2 = ?", userId, userId).Scan(&userFriends)

	return userFriends
}

func CreateUserFriend(userFriend *UsersFriends) {
	db.Model(UsersFriends{}).Create(&userFriend)
}

func UpdateUserFriendStatus(id int, friends UsersFriends) {
	db.Model(UsersFriends{}).Where("id = ?", id).Updates(UsersFriends{Active: friends.Active, Status: friends.Status})

}
