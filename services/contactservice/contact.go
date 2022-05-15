package contactservice

import (
	"errors"
	"ginIMApi/models"
	"ginIMApi/packages/utils"
	"log"
	"strconv"
	"time"
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
	userIdInt, _ := strconv.Atoi(userId)
	friendIdInt, _ := strconv.Atoi(friendId)
	user1, user2 := utils.GetUserSort(userIdInt, friendIdInt)
	if models.IsUserFriend(user1, user2) {
		return true, nil
	}
	//添加好友
	friendApply := models.GetUserApplyRecord(userId, friendId)
	log.Print(friendApply)
	if friendApply.ID > 0 {
		models.UpdateApplyRecordRemark(friendApply.ID, remark)
	} else {
		apply := models.UsersFriendsApply{
			UserId:   userIdInt,
			FriendId: friendIdInt,
			Remarks:  remark,
		}
		models.CreateApplyRecord(&apply)
	}

	return true, nil
}

func SearchUserContact(mobile string) models.User {
	userInfo, _ := models.GetUserByMobile(mobile)

	return userInfo
}

func GetUserContactApplyRecords(userId interface{}) []models.RecordResult {

	//用户联系人信息
	userContactList := models.GetUserApply(userId)

	return userContactList
}

func AcceptInvitation(userId interface{}, applyId interface{}, remark string) int {

	//申请记录
	applyRecord := models.GetUserApplyRecordById(applyId)
	if applyRecord.ID == 0 {
		return applyRecord.ID
	}

	//更新记录
	models.UpdateApplyRecordStatus(applyRecord.ID, 1)
	//好友记录
	user1, user2 := utils.GetUserSort(applyRecord.UserId, applyRecord.FriendId)
	userFriend := models.GetUserFriend(user1, user2)

	if userFriend.ID > 0 {
		//谁是申请人
		active := 2
		if applyRecord.UserId == userFriend.User1 && applyRecord.FriendId == userFriend.User2 {
			active = 1
		}
		//更新好友信息
		userFriend.Active = active
		userFriend.Status = 1
		models.UpdateUserFriendStatus(userFriend.ID, userFriend)
	} else {
		//好友信息
		friendInfo, _ := models.GetUserByUserId(applyRecord.FriendId)
		userFriend = models.UsersFriends{User1: user1, User2: user2, User1Remark: friendInfo.Nickname, User2Remark: remark, Active: 1, Status: 1, AgreeTime: time.Now()}
		userIdInt, _ := strconv.Atoi(userId.(string))
		if user1 == userIdInt {
			userFriend.User1Remark = remark
			userFriend.Active = 2
			userFriend.User2Remark = friendInfo.Nickname
		}
		//添加好友信息
		models.CreateUserFriend(&userFriend)
	}

	return applyRecord.UserId
}
