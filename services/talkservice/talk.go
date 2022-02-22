package talkservice

import "ginIMApi/models"

func GetUserChatRecordList(uid interface{}, toUserId string, page int) (recordId uint, data []interface{}) {

	userId := uid.(string)
	recordList := models.GetUserChatRecordList(userId, toUserId, page)

	for index, item := range recordList {
		if index == 0 {
			recordId = item.ID
		}
		tmp := map[string]interface{}{
			"avatar":       item.Avatar,
			"code_block":   item.CodeBlock,
			"content":      item.Content,
			"created_at":   item.CreatedAt.Format("2006-01-02 15:04:05"),
			"file":         item.File,
			"forward":      item.Forward,
			"group_avatar": item.GroupAvatar,
			"group_name":   item.GroupName,
			"invite":       item.Invite,
			"is_mark":      item.IsMark,
			"is_read":      item.IsRead,
			"is_revoke":    item.IsRevoke,
			"login":        item.Login,
			"msg_type":     item.MsgType,
			"nickname":     item.Nickname,
			"receiver_id":  item.ReceiverId,
			"talk_type":    item.TalkType,
			"user_id":      item.UserId,
			"vote":         item.Vote,
		}

		data = append(data, tmp)
	}

	return recordId, data
}
