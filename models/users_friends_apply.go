package models

import "ginIMApi/packages/setting"

type UsersFriendsApply struct {
	Model
	UserId   int `json:"user_id"`
	FriendId int `json:"friend_id"`
	Remarks  string `json:"remarks"`
	Status   int    `json:"status"`
}

func (UsersFriendsApply) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "users_friends_apply"
}

func GetUserApplyRecord(userId interface{}, friendId interface{}) UsersFriendsApply {
	var userFriendApply UsersFriendsApply
	db.Model(UsersFriendsApply{}).Where("user_id = ? AND friend_id = ?", userId, friendId).Order("id desc").First(&userFriendApply)

	return userFriendApply
}

func UpdateApplyRecord(id int, remark string) {
	db.Model(UsersFriendsApply{}).Where("id = ?", id).Updates(map[string]interface{}{"remarks": remark})
}

func CreateApplyRecord(apply *UsersFriendsApply) {
	db.Model(UsersFriendsApply{}).Create(&apply)
}
