package models

type Contact struct {
	ID           int    `json:"id"`
	Nickname     string `json:"nickname"`
	Status       string `json:"online"`
	Motto        string `json:"motto"`
	Avatar       string `json:"avatar"`
	Gender       int    `json:"gender"`
	FriendRemark string `json:"friend_remark"`
}
