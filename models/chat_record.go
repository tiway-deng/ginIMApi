package models

import (
	"ginIMApi/packages/setting"
	"time"
)

type ChatRecord struct {
	Model
	Source    int       `json:"source"`
	MsgType   int       `json:"msg_type"`
	UserId    int       `json:"user_id"`
	ReceiveId int       `json:"receive_id"`
	Content   string    `json:"content"`
	IsRevoke  int       `json:"is_revoke"`
	DeletedAt time.Time `gorm:"default:0",json:"deleted_at"`
}

type userChatItem struct {
	ID        uint   `json:"id"`
	Avatar    string `json:"avatar"`
	CodeBlock struct {
		Code string `json:"code"`
	} `json:"code_block"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	File      struct {
		FileUrl string `json:"file_url"`
	} `json:"file"`
	Forward struct {
		Num  string `json:"num"`
		List string `json:"list"`
	} `json:"forward"`
	GroupAvatar string `json:"group_avatar"`
	GroupName   string `json:"group_name"`
	Invite      struct {
		Type string `json:"type"`
	}
	IsMark     int         `json:"is_mark"`
	IsRead     int         `json:"is_read"`
	IsRevoke   int         `json:"is_revoke"`
	Login      interface{} `json:"login"`
	MsgType    int         `json:"msg_type"`
	Nickname   string      `json:"nickname"`
	ReceiverId int         `json:"receiver_id"`
	TalkType   int         `json:"talk_type"`
	UserId     int         `json:"user_id"`
	Vote       interface{} `json:"vote"`
}

func (ChatRecord) TableName() string {
	return "chat_records"
}

//get user chat record list
func GetUserChatRecordList(uid string, toUid string, offset int) []userChatItem {
	var results []userChatItem

	tablePrefix := setting.DatabaseSetting.TablePrefix
	db.Table(tablePrefix + "chat_records as cr").
		//db.Model(&ChatRecord{}).
		Select("cr.*, u.nickname ,u.avatar").
		Joins("INNER Join " + tablePrefix + "users as u on u.id = cr.user_id").
		Where("(" +
			"(" + "cr.user_id = " + uid + " and cr.receive_id=" + toUid + ")" +
			" or " +
			"(" + "cr.user_id = " + toUid + " and cr.receive_id=" + uid + ")" +
			")").
		Order("cr.id desc").
		Offset(offset).
		Limit(30).
		Scan(&results)

	return results
}

func CreateChatRecord(u *ChatRecord) {
	db.Create(&u)
}
