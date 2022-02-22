package v1

import (
	"ginIMApi/constants/e"
	"ginIMApi/models"
	"ginIMApi/packages/utils"
	"ginIMApi/services/userservice"
	"ginIMApi/validators"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func GetUserDetail(c *gin.Context) {
	appG := utils.Gin{C: c}
	userId := c.MustGet("user_id")
	userInfo, err := models.GetUserByUserId(userId)
	if err != nil {
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}

	//返回数据
	appG.Response(e.SUCCESS, map[string]interface{}{
		"mobile":   userInfo.Mobile,
		"nickname": userInfo.Nickname,
		"avatar":   userInfo.Avatar,
		"motto":    userInfo.Motto,
		"email":    userInfo.Email,
		"gender":   userInfo.Gender,
	})

}

func GetUserSetting(c *gin.Context) {
	appG := utils.Gin{C: c}
	userId := c.MustGet("user_id")
	userInfo, err := models.GetUserByUserId(userId)
	if err != nil {
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}

	//返回数据
	appG.Response(e.SUCCESS, map[string]interface{}{
		"user_info": utils.JsonStruct{
			"uid":      userInfo.ID,
			"nickname": userInfo.Nickname,
			"mobile":   userInfo.Mobile,
			"avatar":   userInfo.Avatar,
			"gender":   userInfo.Gender,
			"motto":    userInfo.Motto,
		},
		"setting": utils.JsonStruct{
			"theme_mode":            "",
			"theme_bag_img":         "",
			"theme_color":           "",
			"notify_cue_tone":       "",
			"keyboard_event_notify": "",
		},
	})
}

func EditUserDetail(c *gin.Context) {
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

	userId := c.MustGet("user_id")

	_, err := userservice.UpdateUser(userId, form)
	if err != nil {
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}

	appG.Response(e.SUCCESS, nil)

}

