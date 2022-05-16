package v1

import (
	"ginIMApi/constants/e"
	"ginIMApi/models"
	"ginIMApi/packages/utils"
	"ginIMApi/services/talkservice"
	"ginIMApi/services/userservice"
	"ginIMApi/validators"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"strconv"
)

func List(c *gin.Context) {
	appG := utils.Gin{C: c}
	userId := c.MustGet("user_id")

	talkList := userservice.GetUserChatList(userId.(string), utils.GetPageOffset(c))

	//返回数据
	appG.Response(e.SUCCESS, talkList)
}

func Create(c *gin.Context) {
	appG := utils.Gin{C: c}
	//参数验证
	json := make(utils.JsonStruct)
	c.BindJSON(&json)

	userId := c.MustGet("user_id")

	if json["type"] == 1 {
		userIdInt,_ :=strconv.Atoi(userId.(string))
		user1, user2 := utils.GetUserSort(userIdInt, int(json["receive_id"].(float64)))
		isFriend := models.IsUserFriend(user1,user2)
		if !isFriend {
			appG.Response(e.ERROR, "你们还不是好友")
		}
	}
	//添加聊天列表
	userIdInt,_ := strconv.Atoi(userId.(string))
	receiveId,_ := strconv.Atoi(json["receive_id"].(string))
	result :=models.UpsertChatItem(userIdInt,receiveId,int(json["type"].(float64)))

	//返回数据
	appG.Response(e.SUCCESS, map[string]interface{}{
		"talkItem":result,
	})
}

func UpdateUnreadNum(c *gin.Context) {
	appG := utils.Gin{C: c}
	//返回数据
	appG.Response(e.SUCCESS, nil)
}

func UserChatRecords(c *gin.Context) {
	appG := utils.Gin{C: c}
	//参数验证
	var form validators.TalkRecord
	c.Bind(&form)
	valid := validation.Validation{}
	check, _ := valid.Valid(form)
	if !check {
		utils.MarkErrors(valid.Errors)
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}

	userId := c.MustGet("user_id")
	toUid := strconv.Itoa(form.ReceiveId)

	//todo 是否属于群成员
	recordId, recordList := talkservice.GetUserChatRecordList(userId, toUid, utils.GetPageOffset(c))

	//返回数据
	appG.Response(e.SUCCESS, map[string]interface{}{
		"limit":     30,
		"record_id": recordId,
		"rows":      recordList,
	})
}
