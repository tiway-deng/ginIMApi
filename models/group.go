package models

import "time"

type Group struct {
	Model
	CreatorId   string    `json:"creator_id"`
	GroupName   string    `json:"group_name"`
	Profile     string    `json:"profile"`
	Avatar      int       `json:"avatar"`
	MaxNum      string    `json:"max_num"`
	IsOvert     string    `json:"is_overt"`
	IsDismiss   string    `json:"is_dismiss"`
	DismissedAt string    `json:"is_dismiss"`
	DeletedAt   time.Time `gorm:"default:0",json:"deleted_at"`
}

//get group info by id
func GetGroupById(id int) Group {
	var group Group
	db.Where("id = ?", id).First(&group)
	return group
}
