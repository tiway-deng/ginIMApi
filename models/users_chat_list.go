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
	UpdatedAt  time.Time `gorm:"default:0",json:"updated_at"`
}

type UsersChatListDetail struct {
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
func GetUserChatList(uid string, offset int) []UsersChatListDetail {
	var results []UsersChatListDetail

	tablePrefix := setting.DatabaseSetting.TablePrefix
	db.Table(tablePrefix+"users_chat_lists as list").
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

func GetChatItem(userId int, receiveId int, chatType int) (chatItem UsersChatList) {

	db.Table("im_users_chat_lists").Where("uid = ? AND type = ? AND friend_id = ?",userId,chatType,receiveId).First(&chatItem)

	return chatItem
}

func UpsertChatItem(userId int,receiveId int,chatType int) map[string]interface{}{
	chatItem := GetChatItem(userId,receiveId,chatType)
	if chatItem.ID == 0 {
		chatItem.Type = chatType
		chatItem.Uid = userId
		chatItem.Status = 1
		chatItem.FriendId = receiveId
		chatItem.GroupId = 0
		chatItem.CreatedAt = time.Now()
	}
	chatItem.Status = 1
	chatItem.UpdatedAt = time.Now()
	log.Println(chatItem)
	db.Save(&chatItem)

	return map[string]interface{}{
		"id":chatItem.ID,
		"type":chatItem.Type,
		"friend_id":chatItem.FriendId,
		"group_id":chatItem.GroupId,
	}
}

