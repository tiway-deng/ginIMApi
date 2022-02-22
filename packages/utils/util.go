package utils

import "ginIMApi/packages/setting"

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}

//返回用户id从小到大
func GetUserSort(userId1 int, userId2 int) (int, int) {
	if userId1 > userId2 {
		userId1, userId2 = userId2, userId1
	}
	return userId1, userId2
}
