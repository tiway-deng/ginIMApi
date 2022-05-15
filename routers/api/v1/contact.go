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
		"unread_num": len(unreadNum),
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
	json := make(map[string]interface{})
	_ = c.BindJSON(&json)
	userId := c.MustGet("user_id")
	FriendIdStr := strconv.FormatFloat(json["friend_id"].(float64),'f', 0, 64)
	//添加联系人
	isSuccess, err := contactservice.AddUserContact(userId.(string), FriendIdStr, json["remarks"].(string))
	if err != nil {
		appG.Response(e.ERROR, err)
	}
	//添加未读信息
	cacheApplyFriend := cache.NewApplyFriend()
	_,err = cacheApplyFriend.IncrApplyFriendUnRead(FriendIdStr)
	if err != nil {
		appG.Response(e.ERROR, "失败")
	}
	//消息通知
	ws.ApplyFriendMsg(FriendIdStr, map[string]interface{}{
		"sender":  userId,
		"receive": FriendIdStr,
		"type":    1,
		"status":  1,
		"remark":  json["remarks"],
	})
	//返回数据
	appG.Response(e.SUCCESS, isSuccess)
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
	if (userInfo.ID == 0) {
		appG.Response(e.INVALID_PARAMS,nil)
	}

	//返回数据
	appG.Response(e.SUCCESS, userInfo)
}

func UserContactApplyRecords(c *gin.Context)  {
	appG := utils.Gin{C: c}
	userId := c.MustGet("user_id")

	contactList := contactservice.GetUserContactApplyRecords(userId)
	data := map[string]interface{}{
		"rows":     contactList,
		"page": 1,
		"page_size":      100,
		"page_total": len(contactList),
		"total": len(contactList),
	}

	//返回数据
	appG.Response(e.SUCCESS, data)
}

func AcceptUser(c *gin.Context)  {
	appG := utils.Gin{C: c}
	userId := c.MustGet("user_id")

	json := make(map[string]interface{})
	c.BindJSON(&json)
	//var form validators.AcceptInvitation
	//c.Bind(&form)
	//
	//valid := validation.Validation{}
	//check, _ := valid.Valid(form)
	//if !check {
	//	utils.MarkErrors(valid.Errors)
	//	appG.Response(e.INVALID_PARAMS, nil)
	//	return
	//}

	//添加好友
	friendId := contactservice.AcceptInvitation(userId,json["apply_id"],json["remarks"].(string))
	if friendId == 0 {
		appG.Response(e.INVALID_PARAMS, "处理失败")
		return
	}
	//判断对方是否在线
	friendIdStr := strconv.Itoa(friendId)
	ws.ApplyFriendMsg(friendIdStr, map[string]interface{}{
		"sender":  userId,
		"receive": friendIdStr,
		"type":    1,
		"status":  1,
		"remark":  "",
	})

	//返回数据
	appG.Response(e.SUCCESS, "")
}
