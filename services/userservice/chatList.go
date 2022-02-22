package userservice

import (
	"ginIMApi/models"
)

func GetUserChatList(uid string, offset int) []interface{} {
	userChatList := models.GetUserChatList(uid, offset)
	var data []interface{}
	for _, item := range userChatList {
		tmp := map[string]interface{}{
			"id":          item.ID,
			"type":        item.Type,
			"friend_id":   item.FriendId,
			"group_id":    item.GroupId,
			"name":        item.GroupName,
			"avatar":      item.GroupAvatar,
			"remark_name": item.Nickname,
			"unread_num":  0,
			"msg_text":    "...",
			"updated_at":  item.UpdatedAt.Format("2006-01-02 15:04:05"),
			"online":      0,
			"is_top":      0,
			"not_disturb": item.NotDisturb,
		}
		if tmp["type"] == 1 {
			tmp["name"] = item.Nickname
			tmp["avatar"] = item.UserAvatar
			tmp["unread_num"] = 2
			tmp["online"] = item.UserStatus
		}

		data = append(data, tmp)
	}

	return data
}
