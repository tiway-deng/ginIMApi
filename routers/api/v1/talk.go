package v1

import (
	"ginIMApi/constants/e"
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
	var form validators.User
	c.Bind(&form)
	valid := validation.Validation{}
	check, _ := valid.Valid(form)
	if !check {
		utils.MarkErrors(valid.Errors)
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}

	//返回数据
	appG.Response(e.SUCCESS, nil)
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
