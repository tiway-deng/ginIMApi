package models

import (
	"ginIMApi/packages/setting"
	"log"
	"time"
)

type UsersChatList struct {
	Model
	Type       int       `json:"type"`
	Uid        int       `json:"uid"`
	FriendId   int       `json:"friend_id"`
	GroupId    int       `json:"group_id"`
	Status     int       `json:"status"`
	IsTop      int       `json:"is_top"`
	NotDisturb int       `json:"not_disturb"`
	DeletedAt  time.Time `gorm:"default:0",json:"deleted_at"`
}

type ChatListInfo struct {
	ID          uint      `json:"id"`
	Type        int       `json:"type"`
	FriendId    int       `json:"friend_id"`
	GroupId     int       `json:"group_id"`
	UpdatedAt   time.Time `json:"updated_at"`
	NotDisturb  int       `json:"not_disturb"`
	IsTop       int       `json:"is_top"`
	UserId      int       `json:"user_id"`
	UserAvatar  string    `json:"user_avatar"`
	Nickname    string    `json:"nickname"`
	UserStatus  int       `json:"user_status"`
	GroupName   string    `json:"group_name"`
	GroupAvatar string    `json:"group_avatar"`
}

//get user chat list
func GetUserChatList(uid string, offset int) []ChatListInfo {
	var results []ChatListInfo

	tablePrefix := setting.DatabaseSetting.TablePrefix
	db.Table(tablePrefix+"users_chat_list as list").
		Select(
			"list.id,list.type,list.friend_id,list.group_id,list.updated_at,list.not_disturb,list.is_top,"+
				"u.id as user_id,u.avatar as user_avatar,u.nickname,u.status as user_status,"+
				"g.group_name, g.avatar as group_avatar").
		Joins("left Join "+tablePrefix+"users as u on u.id = list.friend_id").
		Joins("left Join "+tablePrefix+"group as g on g.id = list.group_id").
		Where("list.uid = ? AND list.status = ?", uid, 1).
		Order("updated_at desc").
		Offset(offset).
		Limit(30).
		Find(&results)

	log.Println(results)

	return results
}
