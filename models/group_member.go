package models

import (
	"ginIMApi/packages/setting"
	"time"
)

type GroupMember struct {
	Model
	GroupId   int       `json:"group_id"`
	UserId    int       `json:"user_id"`
	Leader    int       `json:"leader"`
	IsMute    int       `json:"is_mute"`
	IsQuit    int       `json:"is_quit"`
	UserCard  string    `json:"user_card"`
	DeletedAt time.Time `gorm:"default:0",json:"deleted_at"`
}

//set table name
func (GroupMember) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "group_member"
}

//get group members by group id
func GetGroupMemberList(groupId int, fields []string) []GroupMember {
	var results []GroupMember
	if len(fields) == 0 {
		fields = []string{"id", "group_id"}
	}
	db.Select(fields).Where("group_id = ? AND is_quit = ?", groupId, 0).Find(&results)
	return results
}

//get group member info by group id and user id
func GroupMemberInfo(groupId int, userId int, fields []string) GroupMember {
	var member GroupMember
	db.Select(fields).Where("group_id = ? AND user_id = ?", groupId, userId).First(&member)
	return member
}

//return true if it is exist
func IsGroupMember(groupId int, userId int) bool {
	isGroupMember := false
	groupMemberInfo := GroupMemberInfo(groupId, userId, []string{"id"})
	if groupMemberInfo.ID > 0 {
		isGroupMember = true
	}
	return isGroupMember
}
