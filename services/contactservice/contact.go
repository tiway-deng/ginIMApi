package contactservice

import (
	"errors"
	"ginIMApi/models"
	"ginIMApi/packages/utils"
)

func GetUserContactList(userId interface{}) []models.Contact {

	//用户联系人信息
	userContactList := models.GetUserContactList(userId)

	return userContactList
}

func AddUserContact(userId interface{}, friendId int, remark string) (bool, error) {

	//用户是否存在
	friendInfo, _ := models.GetUserByUserId(userId)
	if friendInfo.ID > 0 {
		return false, errors.New("用户不存在")
	}
	//是否是好友
	user1, user2 := utils.GetUserSort(userId.(int), friendId)
	if models.IsUserFriend(user1, user2) {
		return true,nil
	}
	//添加好友

	return false, nil
}

func SearchUserContact(mobile string) models.User {
	userInfo,_ := models.GetUserByMobile(mobile)

	return userInfo
}

func GetUserContactApplyRecords(userId interface{}) []models.Contact {

	//用户联系人信息
	userContactList := models.GetUserContactList(userId)

	return userContactList
}
