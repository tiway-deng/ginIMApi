package models

import (
	"ginIMApi/packages/setting"
	"time"
)

type UsersFriendsApply struct {
	Model
	UserId    int       `json:"user_id"`
	FriendId  int       `json:"friend_id"`
	Remarks   string    `json:"remarks"`
	Status    int       `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RecordResult struct {
	ID        int       `json:"id"`
	UserId    int       `json:"user_id"`
	FriendId  int       `json:"friend_id"`
	Remarks   string    `json:"remarks"`
	Status    int       `json:"status"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	Mobile    string    `json:"mobile"`
	CreatedAt time.Time `json:"created_at"`
}

func (UsersFriendsApply) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "users_friends_apply"
}

func GetUserApplyRecord(userId interface{}, friendId interface{}) UsersFriendsApply {
	var userFriendApply UsersFriendsApply
	db.Model(UsersFriendsApply{}).Where("user_id = ? AND friend_id = ? AND status = 0", userId, friendId).Order("id desc").First(&userFriendApply)

	return userFriendApply
}

func UserApplyRecordExist(userId interface{}, friendId interface{}) bool {
	var userFriendApply UsersFriendsApply
	db.Model(UsersFriendsApply{}).Where("user_id = ? AND friend_id = ? AND status = 0", userId, friendId).Order("id desc").First(&userFriendApply)

	exist := false
	if userFriendApply.ID > 0 {
		exist = true
	}

	return exist
}

func UpdateApplyRecordRemark(id int, remark string) {
	db.Model(UsersFriendsApply{}).Where("id = ?", id).Updates(map[string]interface{}{"remarks": remark, "updated_at": time.Now()})
}

func UpdateApplyRecordStatus(id int, status int) {
	db.Model(UsersFriendsApply{}).Where("id = ?", id).Updates(map[string]interface{}{"status": status, "updated_at": time.Now()})
}

func CreateApplyRecord(apply *UsersFriendsApply) {
	db.Model(UsersFriendsApply{}).Create(&apply)
}

func GetUserApply(userId interface{}) []RecordResult {

	tablePrefix := setting.DatabaseSetting.TablePrefix
	var result []RecordResult
	db.Raw("SELECT users_friends_apply.id,users_friends_apply.status,users_friends_apply.remarks,users.nickname,users.avatar,users.mobile,users_friends_apply.user_id,users_friends_apply.friend_id,users_friends_apply.created_at from (SELECT id,status,remarks,user_id,friend_id,created_at from "+tablePrefix+"users_friends_apply WHERE user_id = ? ) users_friends_apply "+
		"LEFT JOIN "+tablePrefix+"users as users ON users.id = users_friends_apply.friend_id", userId).Scan(&result)

	return result
}

func GetUserApplyRecordById(id interface{}) UsersFriendsApply {
	var userFriendApply UsersFriendsApply
	db.Where("id = ? ", id).First(&userFriendApply)

	return userFriendApply
}
