package v1

import (
	//"ginIMApi/cache"
	"ginIMApi/constants/e"
	"ginIMApi/models"
	"ginIMApi/packages/utils"
	"ginIMApi/services/userservice"
	"ginIMApi/validators"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	//"log"
	//"strconv"
)

func GetUserDetail(c *gin.Context) {
	appG := utils.Gin{C: c}
	var form validators.UserSearch
	c.Bind(&form)

	//userId := c.MustGet("user_id")

	userInfo, err := models.GetUserByUserId(form.UserId)
	if err != nil {
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}

	//好友请求
	//cacheApplyFriend := cache.NewApplyFriend()
	//friendIdStr := strconv.Itoa(form.UserId)
	//friendApply,error := cacheApplyFriend.GetApplyFriendUnRead(friendIdStr)
	//log.Print(friendApply,error)
	////好友状态
	//user1, user2 := utils.GetUserSort(userId.(int), form.UserId)
	//friendStatus := models.IsUserFriend(user1, user2)
	//log.Print(friendStatus)


	//返回数据
	appG.Response(e.SUCCESS, map[string]interface{}{
		"id":   userInfo.ID,
		"mobile":   userInfo.Mobile,
		"nickname": userInfo.Nickname,
		"avatar":   "https://im.gzydong.club/public/media/image/avatar/20220224/fa7cd682ba942a6d9f04218eb82b7e75_200x200.png",
		"motto":    userInfo.Motto,
		"email":    userInfo.Email,
		"gender":   userInfo.Gender,
		"friend_apply":   1,
		"friend_status":   1,
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

