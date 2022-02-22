package v1

import (
	"ginIMApi/constants/e"
	"ginIMApi/models"
	"ginIMApi/packages/utils"
	userServe "ginIMApi/services/userservice"
	"ginIMApi/validators"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"regexp"
)


func Login(c *gin.Context){
	appG := utils.Gin{C: c}
	valid := validation.Validation{}

	json := make(map[string]string)
	c.BindJSON(&json)
	username := json["mobile"]
	password := json["password"]
	//参数验证
	a := validators.User{Password: password}
	isMobile, _ := regexp.MatchString(`^[\+-]?\d+$`, username)
	if isMobile {
		a.Mobile = username
	} else {
		a.Nickname = username
	}
	ok, _ := valid.Valid(&a)
	if !ok {
		utils.MarkErrors(valid.Errors)
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}
	//登录验证
	userInfo, ok := userServe.CheckUser(username, password)
	if !ok {
		appG.Response(e.ERROR_USER_CHECK_FAIL, nil)
		return
	}
	//生成jwt token
	token, expiredIn, err := utils.GenerateToken(userInfo)
	if err != nil {
		appG.Response(e.ERROR_AUTH_TOKEN, nil)
		return
	}

	//返回数据
	appG.Response(e.SUCCESS, map[string]interface{}{
		"user_info": utils.JsonStruct{
			"uid":       userInfo.ID,
			"nickname":  userInfo.Nickname,
			"mobile":    userInfo.Mobile,
			"avatar":    userInfo.Avatar,
			"sex":       userInfo.Gender,
			"signature": "",
		},
		"authorize": utils.JsonStruct{
			"author_type":  "Bearer",
			"access_token": token,
			"expired_in":   expiredIn,
		},
	})

}


func Register(c *gin.Context) {
	appG := utils.Gin{C: c}
	//参数验证
	var form validators.User
	c.Bind(&form)
	//密码是否相同
	//if form.Password != form.Password2 {
	//	appG.Response(e.ERROR_USER_REGISTER_PASS_SAME, nil)
	//	return
	//}
	valid := validation.Validation{}
	check, _ := valid.Valid(form)
	if !check {
		utils.MarkErrors(valid.Errors)
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}
	//登录验证
	userForm := models.User{Mobile: form.Mobile, Nickname: form.Nickname, Password: form.Password, Avatar: form.Avatar, Gender: form.Gender, Email: form.Email}
	userId, err := userServe.RegisterUser(&userForm)
	if err != nil {
		appG.Response(e.ERROR_USER_REGISTER_USER_SAME, nil)
		return
	}
	//返回数据
	appG.Response(e.SUCCESS, map[string]int{
		"user_id": userId,
	})
}