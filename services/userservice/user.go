package userservice

import (
	"errors"
	"ginIMApi/constants/e"
	"ginIMApi/models"
	"ginIMApi/packages/utils"
	"ginIMApi/validators"
)

func CheckUser(m string, p string) (user models.User, ok bool) {
	user, _ = models.GetUserByUsername(m)
	if user.ID > 0 {
		//密码验证
		md5Pwd := utils.EncodeMD5(p)
		if user.Password == md5Pwd {
			return user,true
		}
	}
	return user,false
}

func RegisterUser(u *models.User) (int, error) {

	//手机号不同重复注册
	mobile := u.Mobile
	userInfo, _ := models.GetUserByUsername(mobile)
	if userInfo.ID > 0 {
		return userInfo.ID, errors.New(e.GetMsg(e.ERROR_USER_REGISTER_USER_SAME))
	}
	//密码加密
	u.Password = utils.EncodeMD5(u.Password)

	//生成用户信息
	models.CreatUser(u)

	return u.ID, nil
}

func UpdateUser(userId interface{}, u validators.User) (userInfo models.User, err error) {
	userInfo, err = models.UpdateUser(userId,u)

	return userInfo,err
}



