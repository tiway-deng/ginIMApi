package validators

type User struct {
	Mobile      string `form:"mobile",valid:"Required; Required"`
	Nickname    string `form:"nickname",valid:"Required; MaxSize(20)"`
	Avatar      string `form:"avatar",valid:"Required; MaxSize(255)"`
	Gender      int    `form:"gender",valid:"Required; range(0,2)"`
	Status      int    `form:"status",valid:"Required; range(0,1)"`
	Password    string `form:"password",valid:"Required; MaxSize(255)"`
	Password2   string `form:"password",valid:"Required; MaxSize(255)"`
	Email       string `form:"email",valid:"Required; Email"`
}

type UserContactAdd struct {
	FriendId int `form:"friend_id",valid:"Required; Numeric"`
	Remarks  string `form:"remarks",valid:"Required; MaxSize(50)"`
}

type UserContactSearch struct {
	Mobile string `form:"mobile",valid:"Required; Phone"`
}

type UserSearch struct {
	UserId int `form:"user_id",valid:"Required; Numeric"`
}
