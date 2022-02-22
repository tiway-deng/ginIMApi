package v1

import (
	"ginIMApi/cache"
	"ginIMApi/constants/e"
	"ginIMApi/packages/utils"
	"ginIMApi/routers/ws"
	"ginIMApi/services/contactservice"
	"ginIMApi/validators"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"strconv"
)

//apply unread num
func ApplyUnreadNum(c *gin.Context) {
	appG := utils.Gin{C: c}

	userId := c.MustGet("user_id").(string)
	cacheApplyFriend := cache.NewApplyFriend()
	unreadNum, _ := cacheApplyFriend.GetApplyFriendUnRead(userId)
	//返回数据
	appG.Response(e.SUCCESS, map[string]interface{}{
		"unread_num": unreadNum,
	})
}

//get user contact list
func GetUserContactList(c *gin.Context) {
	appG := utils.Gin{C: c}
	userId := c.MustGet("user_id")

	contactList := contactservice.GetUserContactList(userId)
	//返回数据
	appG.Response(e.SUCCESS, contactList)
}


//add user contact
func AddUserContact(c *gin.Context) {
	appG := utils.Gin{C: c}
	userId := c.MustGet("user_id")

	var form validators.UserContactAdd
	c.Bind(&form)
	valid := validation.Validation{}
	check, _ := valid.Valid(form)
	if !check {
		utils.MarkErrors(valid.Errors)
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}
	//添加联系人
	contactList, _ := contactservice.AddUserContact(userId, form.FriendId, form.Remarks)
	//添加未读信息
	cacheApplyFriend := cache.NewApplyFriend()
	FriendIdStr := strconv.Itoa(form.FriendId)
	cacheApplyFriend.IncrApplyFriendUnRead(FriendIdStr)
	//消息通知
	ws.ApplyFriendMsg(FriendIdStr, map[string]interface{}{
		"sender":  userId,
		"receive": FriendIdStr,
		"type":    1,
		"status":  1,
		"remark":  form.Remarks,
	})
	//返回数据
	appG.Response(e.SUCCESS, contactList)
}


func SearchUserContact(c *gin.Context) {
	appG := utils.Gin{C: c}

	var form validators.UserContactSearch
	c.Bind(&form)
	valid := validation.Validation{}
	check, _ := valid.Valid(form)
	if !check {
		utils.MarkErrors(valid.Errors)
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}
	//添加联系人
	userInfo := contactservice.SearchUserContact(form.Mobile)

	//返回数据
	appG.Response(e.SUCCESS, userInfo)
}

func UserContactApplyRecords(c *gin.Context)  {
	appG := utils.Gin{C: c}
	userId := c.MustGet("user_id")

	contactList := contactservice.GetUserContactApplyRecords(userId)
	//返回数据
	appG.Response(e.SUCCESS, contactList)
}
