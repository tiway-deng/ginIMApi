package models

import (
	"ginIMApi/packages/setting"
	"ginIMApi/validators"
	"log"
)

type User struct {
	Model
	Mobile      string    `json:"mobile"`
	Nickname    string    `json:"nickname"`
	Avatar      string    `json:"avatar"`
	Gender      int       `json:"gender"`
	Password    string    `json:"password"`
	Motto       string    `json:"motto"`
	Status      int       `json:"status"`
	//Description string    `json:"description"`
	Email       string    `json:"email"`
}

const StatusOnline = 1
const StatusOffline = 0

//get user info by mobile
func GetUserByMobile(m string) (user User, err error) {
	db.Where("mobile = ?", m).First(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

//get user info by user name
func GetUserByUsername(m string) (user User, err error) {
	db.Where("mobile = ?", m).Or("nickname = ?", m).First(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

//create user info
func CreatUser(u *User) {
	db.Create(&u)
}

//get user info by id
func GetUserByUserId(userId interface{}) (user User, err error) {
	db.Where("id = ?", userId).First(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

//update user info by user id
func UpdateUser(userId interface{}, u validators.User) (userInfo User, err error) {

	userInfo, err = GetUserByUserId(userId)
	db.Model(&userInfo).Update(u)

	return userInfo, nil
}

//update user status by user id
func UpdateUserStatus(userId interface{}, status int) User {
	userInfo, _ := GetUserByUserId(userId)
	db.Model(&userInfo).Update(map[string]interface{}{
		"status": status,
	})

	return userInfo
}

func GetUserContactList(userId interface{}) []Contact {

	tablePrefix := setting.DatabaseSetting.TablePrefix
	var contactList []Contact
	res := db.Raw("SELECT users.id,users.nickname,users.avatar,users.status,users.motto,users.gender,tmp_table.friend_remark from "+tablePrefix+"users users "+
		"INNER join("+
		" SELECT id as rid,user2 as uid,user1_remark as friend_remark from "+tablePrefix+"users_friends where user1 = ? and `status` = 1 "+
		"UNION all"+
		" SELECT id as rid,user1 as uid,user2_remark as friend_remark from "+tablePrefix+"users_friends where user2 = ? and `status` = 1 "+
		") tmp_table on tmp_table.uid = users.id", userId, userId).Scan(&contactList)

	log.Println(res)

	return contactList
}
