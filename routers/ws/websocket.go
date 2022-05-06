package ws

import (
	"encoding/json"
	"ginIMApi/models"
	"ginIMApi/packages/utils"
	_ "ginIMApi/packages/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"html"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// 客户端连接详情
type wsClients struct {
	Conn *websocket.Conn `json:"conn"`

	RemoteAddr string `json:"remote_addr"`

	Uid float64 `json:"uid"`

	Username string `json:"username"`

	Avatar string `json:"avatar"`
}

// client & serve 的消息体
type msg struct {
	Event string          `json:"event"`
	Data  interface{}     `json:"data"`
	Conn  *websocket.Conn `json:"conn"`
}

// 变量定义初始化
var (
	wsUpgrader = websocket.Upgrader{}

	clientMsg = msg{}

	userClients = map[string]wsClients{}

	mutex = sync.Mutex{}

	sMsg = make(chan msg)

	chNotify = make(chan int, 1)
)

const EventTalk = "event_talk"
const EventOnLineStatus = "event_online_status"
const EventKeyboard = "event_keyboard"
const EventFriendApply = "event_friend_apply"

//用户在线状态
const OnlineStatusOn = 1
const OnlineStatusOff = 0

func Run(gin *gin.Context) {

	//根据jwt token 获取用户信息
	userId, _ := getUserIdByToken(gin)

	// @see https://github.com/gorilla/websocket/issues/523
	wsUpgrader.CheckOrigin = func(r *http.Request) bool { return true }

	c, _ := wsUpgrader.Upgrade(gin.Writer, gin.Request, nil)

	//建立连接
	handleConnClients(c, userId)
	defer c.Close()
	go read(c)
	go write()

	select {}
}

func read(c *websocket.Conn) {

	defer func() {
		//捕获read抛出的panic
		if err := recover(); err != nil {
			log.Println("read发生错误", err)
		}
	}()

	for {
		_, message, err := c.ReadMessage()
		if err != nil { // 离线通知
			log.Println("===================  离线 ============================")
			disconnect(c)
			c.Close()
			return
		}

		serveMsgStr := message
		// 处理心跳响应
		if string(serveMsgStr) == "PING" {
			c.WriteMessage(websocket.TextMessage, []byte(`{"data":"PONG"}`))
			continue
		}

		json.Unmarshal(message, &clientMsg)
		log.Println("来自客户端的消息", clientMsg)
		if clientMsg.Data == nil {
			return
			//mainProcess(c)
		}
		// 聊天
		if clientMsg.Event == EventTalk {
			source, _ := strconv.Atoi(clientMsg.Data.(map[string]interface{})["source_type"].(string))
			receiveId, _ := strconv.Atoi(clientMsg.Data.(map[string]interface{})["receive_user"].(string))
			userId := clientMsg.Data.(map[string]interface{})["send_user"].(float64)

			var group models.Group
			var userInfo models.User
			var talkingUser = map[string]wsClients{}
			//群聊天
			if source == 2 {
				//群信息
				group = models.GetGroupById(receiveId)
				if group.ID == 0 {
					return
				}
				//判断是否属于群成员
				if !models.IsGroupMember(receiveId, int(userId)) {
					return
				}
				//群成员用户的在线链接信息
				groupMembers := models.GetGroupMemberList(receiveId, []string{"user_id"})
				for _, item := range groupMembers {
					groupMemberId := strconv.Itoa(item.UserId)
					if memberCon, ok := userClients[groupMemberId]; ok {
						talkingUser[groupMemberId] = memberCon
					}
				}
			}
			//私聊
			if source == 1 {
				//判断是否是好友
				user1, user2 := utils.GetUserSort(receiveId, int(userId))
				if !models.IsUserFriend(user1, user2) {
					return
				}
				//好友链接信息
				friendId := strconv.Itoa(receiveId)
				if memberCon, ok := userClients[friendId]; ok {
					talkingUser[friendId] = memberCon
				}
				//用户信息
				userInfo, _ = models.GetUserByUserId(userId)
			}

			//创建聊天记录
			chatRecord := models.ChatRecord{
				Source:    source,
				MsgType:   1,
				UserId:    int(userId),
				ReceiveId: receiveId,
				Content:   html.EscapeString(clientMsg.Data.(map[string]interface{})["text_message"].(string)),
			}
			models.CreateChatRecord(&chatRecord)
			if chatRecord.ID == 0 {
				return
			}

			//返回信息
			talkType, _ := strconv.Atoi(clientMsg.Data.(map[string]interface{})["source_type"].(string))
			dataRes := []interface{}{
				EventTalk,
				map[string]interface{}{
					"send_user":    chatRecord.UserId,
					"receive_user": chatRecord.ReceiveId,
					"source_type":  talkType,
					"data": map[string]interface{}{
						"id":           chatRecord.ID,
						"talk_type":    talkType,
						"msg_type":     chatRecord.MsgType,
						"user_id":      chatRecord.UserId,
						"receiver_id":  chatRecord.ReceiveId,
						"nickname":     userInfo.Nickname,
						"avatar":       userInfo.Avatar,
						"group_name":   group.GroupName,
						"group_avatar": group.Avatar,
						"file":         nil,
						"code_block":   nil,
						"forward":      nil,
						"invite":       nil,
						"vote":         nil,
						"login":        nil,
						"content":      chatRecord.Content,
						"created_at":   chatRecord.CreatedAt,
						"is_revoke":    chatRecord.IsRevoke,
						"is_mark":      0,
						"is_read":      0,
					},
				},
			}
			////用户本人推送信息
			jsonStrServeMsg := msg{
				Event: EventTalk,
				Data:  dataRes,
				Conn:  c,
			}
			sMsg <- jsonStrServeMsg

			//在线用户推送
			if len(talkingUser) > 0 {
				for _, item := range talkingUser {
					//聊天好友推送信息
					jsonToUserStrServeMsg := msg{
						Event: EventTalk,
						Data:  dataRes,
						Conn:  item.Conn,
					}
					sMsg <- jsonToUserStrServeMsg
				}
			}
		}
		//对方正在输入中
		if clientMsg.Event == EventKeyboard {
			receiveId, _ := clientMsg.Data.(map[string]interface{})["receive_user"].(string)
			userId := clientMsg.Data.(map[string]interface{})["send_user"].(float64)
			if userC, ok := userClients[receiveId]; ok {
				//对方正在输入中
				dataRes := []interface{}{
					EventKeyboard,
					map[string]interface{}{
						"send_user":    userId,
						"receive_user": receiveId,
					},
				}
				jsonStrServeMsg := msg{
					Event: EventTalk,
					Data:  dataRes,
					Conn:  userC.Conn,
				}
				sMsg <- jsonStrServeMsg
			}
		}

	}
}

func write() {
	defer func() {
		//捕获write抛出的panic
		if err := recover(); err != nil {
			log.Println("write发生错误", err)
		}
	}()

	for {
		select {
		case cl := <-sMsg:
			chNotify <- 1
			serveMsgStr, _ := json.Marshal(cl.Data)
			notify(cl.Conn, string(serveMsgStr))
			<-chNotify
			//case o := <-offline:

		}
	}
}

func getUserIdByToken(c *gin.Context) (string, error) {
	token := c.Query("token")
	claims, err := utils.ParseToken(token)
	if err != nil {
		log.Println("用户token失效")
		return "", err
	}
	return claims.UserId, nil
}

// 处理建立连接的用户
func handleConnClients(c *websocket.Conn, userId string) {
	//用户信息
	userInfo := models.UpdateUserStatus(userId, models.StatusOnline)
	mutex.Lock()
	userClients[userId] = wsClients{
		Conn:       c,
		RemoteAddr: c.RemoteAddr().String(),
		Uid:        float64(userInfo.ID),
		Username:   userInfo.Nickname,
		Avatar:     userInfo.Avatar,
	}
	mutex.Unlock()

	onlineStatusNotify(userId, OnlineStatusOn)

}

func onlineStatusNotify(userId string, status int) {
	//在线好友
	userIdInt, _ := strconv.Atoi(userId)
	userFriends := models.GetUserFriends(userId, []string{"user1", "user2"})

	notifyMsg := "用户上线通知"
	if status == OnlineStatusOff {
		notifyMsg = "用户下线通知"
	}
	for _, item := range userFriends {
		var friendUserId string
		if item.User1 == userIdInt {
			friendUserId = strconv.Itoa(item.User2)
		} else {
			friendUserId = strconv.Itoa(item.User1)
		}
		if friendC, ok := userClients[friendUserId]; ok {
			//好友上线状态通知
			dataRes := []interface{}{
				EventOnLineStatus,
				map[string]interface{}{
					"user_id": userId,
					"status":  status,
					"notify":  notifyMsg,
				},
			}
			jsonStrServeMsg := msg{
				Event: EventTalk,
				Data:  dataRes,
				Conn:  friendC.Conn,
			}
			sMsg <- jsonStrServeMsg
		}
	}
}

// 统一消息发放
func notify(conn *websocket.Conn, msg string) {
	conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

// 离线通知
func disconnect(conn *websocket.Conn) {
	for index, con := range userClients {
		if con.RemoteAddr == conn.RemoteAddr().String() {
			log.Println("----断开链接----", con.Uid)
			data := map[string]interface{}{
				"username": con.Username,
				"uid":      con.Uid,
				"time":     time.Now().UnixNano() / 1e6, // 13位  10位 => now.Unix()
			}
			jsonStrServeMsg := msg{
				Event: EventOnLineStatus,
				Data:  data,
			}
			serveMsgStr, _ := json.Marshal(jsonStrServeMsg)

			//删除下线用户
			mutex.Lock()
			models.UpdateUserStatus(con.Uid, models.StatusOffline)
			delete(userClients, index)
			mutex.Unlock()
			con.Conn.Close()
			notify(conn, string(serveMsgStr))

			onlineStatusNotify(strconv.Itoa(int(con.Uid)), OnlineStatusOff)
		}
	}
}

//好友申请消息
func ApplyFriendMsg(friendId string,data map[string]interface{}) {
	if userC, ok := userClients[friendId]; ok {
		//对方正在输入中
		dataRes := []interface{}{
			EventFriendApply,
			data,
		}
		jsonStrServeMsg := msg{
			Event: EventFriendApply,
			Data:  dataRes,
			Conn:  userC.Conn,
		}
		sMsg <- jsonStrServeMsg
	}
}


