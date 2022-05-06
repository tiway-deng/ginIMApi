package routers

import (
	"ginIMApi/middlewares/cors"
	"ginIMApi/middlewares/jwt"
	_ "ginIMApi/middlewares/jwt"
	"ginIMApi/routers/api"
	v1 "ginIMApi/routers/api/v1"
	"ginIMApi/routers/ws"
	"github.com/gin-gonic/gin"
	_ "net/http"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Cors())
	//
	apiG := r.Group("/api")
	apiG.POST("/upload", api.UploadImage)

	//api 接口
	apiv1 := r.Group("/api/v1")
	//用户登录注册
	auth := apiv1.Group("/auth")
	{
		auth.POST("login", v1.Login)
		auth.POST("register", v1.Register)
	}

	//jwt 鉴权
	apiv1.Use(jwt.JWT()).Use(cors.Cors())
	{
		//用户信息
		users := apiv1.Group("/users")
		{
			users.GET("setting", v1.GetUserSetting)
			users.GET("detail", v1.GetUserDetail)
			users.GET("search-user", v1.GetUserDetail)
		}

		talk := apiv1.Group("/talk")
		{
			talk.GET("list", v1.List)
			talk.POST("create", v1.Create)
			talk.POST("update-unread-num", v1.UpdateUnreadNum)

			talk.GET("records", v1.UserChatRecords)
		}

		contacts := apiv1.Group("/contacts")
		{
			contacts.GET("list", v1.GetUserContactList)
			contacts.GET("apply-unread-num", v1.ApplyUnreadNum)
			contacts.POST("add", v1.AddUserContact)
			contacts.GET("search", v1.SearchUserContact)
			contacts.GET("apply-records", v1.UserContactApplyRecords)
		}

		emoticon := apiv1.Group("/emoticon")
		{
			emoticon.GET("user-emoticon", v1.UserEmoticon)
		}

	}

	//websocket
	r.GET("socket.io", ws.Run)

	return r
}
