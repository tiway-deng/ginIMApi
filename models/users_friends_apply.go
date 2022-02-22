package models

type UsersFriendsApply struct {
	Model
	Mobile      string `json:"mobile"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Gender      int    `json:"gender"`
	Password    string `json:"password"`
	Motto       string `json:"motto"`
	Status      int    `json:"status"`
	Description string `json:"description"`
	Email       string `json:"email"`
}


