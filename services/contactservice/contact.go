package contactservice

import (
	"errors"
	"ginIMApi/models"
	"ginIMApi/packages/utils"
	"log"
	"strconv"
)

func GetUserContactList(userId interface{}) []models.Contact {

	//用户联系人信息
	userContactList := models.GetUserContactList(userId)

	return userContactList
}

func AddUserContact(userId string, friendId string, remark string) (bool, error) {

	//用户是否存在
	friendInfo, _ := models.GetUserByUserId(friendId)
	if friendInfo.ID == 0 {
		return false, errors.New("用户不存在")
	}
	//是否是好友
	userIdInt,_ := strconv.Atoi(userId)
	friendIdInt,_ := strconv.Atoi(friendId)
	user1, user2 := utils.GetUserSort(userIdInt, friendIdInt)
	if models.IsUserFriend(user1, user2) {
		return true,nil
	}
	//添加好友
	friendApply := models.GetUserApplyRecord(userId,friendId)
	log.Print(friendApply)
	if friendApply.ID > 0 {
		models.UpdateApplyRecord(friendApply.ID,remark)
	}else{
		apply := models.UsersFriendsApply{
			UserId: userIdInt,
			FriendId:friendIdInt,
			Remarks:remark,
		}
		models.CreateApplyRecord(&apply)
	}

	return true, nil
}

func SearchUserContact(mobile string) models.User {
	userInfo,_ := models.GetUserByMobile(mobile)

	return userInfo
}

func GetUserContactApplyRecords(userId interface{}) []models.RecordResult {

	//用户联系人信息
	userContactList := models.GetUserApply(userId)

	return userContactList
}
