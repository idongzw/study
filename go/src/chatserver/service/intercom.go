/*
 * @Author: dzw
 * @Date: 2020-03-11 09:32:56
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-12 17:12:19
 */

package service

import (
	"chatserver/chatlog"
	"chatserver/mysql"
	"chatserver/protos"
	"database/sql"
	"strconv"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
)

var (
	// 保存用户信息
	// map[int64]*protos.UserInfo
	userInfoMap sync.Map

	// 在线用户信息
	// map[int64]msg
	userOnlineMap sync.Map

	// 保存客户端channel信息
	// map[int64]chan *interMsg
	chanMsgMap sync.Map

	// 数据库操作
	db *sql.DB // sql.DB是一个数据库（操作）句柄，代表一个具有零到多个底层连接的连接池。它可以安全地被多个goroutine同时使用

	// msg process func map
	handleMsgFuncMap map[string]func(proto.Message, *Client)
)

func init() {
	handleMsgFuncMap = make(map[string]func(proto.Message, *Client), 5)

	handleMsgFuncMap["protos.SignupRequest"] = handleSignupReq
	handleMsgFuncMap["protos.LoginRequest"] = handleLoginReq
	handleMsgFuncMap["protos.ChatMsg"] = handleChatMsg
	handleMsgFuncMap["protos.LogoutRequest"] = handleLogoutReq
	handleMsgFuncMap["protos.HeartbeatMsg"] = handleHeartbeatMsg
}

// 内部通信
type interMsg struct {
	msgName string
	msg     proto.Message
}

func (im *interMsg) handleMsg(c *Client) {
	f, ok := handleMsgFuncMap[im.msgName]
	if ok {
		f(im.msg, c)
	} else {
		chatlog.Warning("msg type: [%v],not handle func", im.msgName)
	}
}

func handleSignupReq(pm proto.Message, c *Client) {
	signupReq, ok := pm.(*protos.SignupRequest)
	if ok {
		chatlog.Trace("%v", signupReq)
	}

	userID := time.Now().Unix() // timestamp
	password := strconv.FormatInt(userID, 10)
	signupRsp := &protos.SignupResponse{
		Result:   0,
		Id:       uint64(userID),
		Password: password,
	}

	user := &protos.UserInfo{
		Id:       uint64(userID),
		Password: password,
		Name:     signupReq.GetName(),
		Age:      signupReq.GetAge(),
		Gender:   signupReq.GetGender(),
	}

	sqlStr := "insert into userInfo(id, password, name, age, gender) values(?,?,?,?,?)"
	err := mysql.InsertInfo(db, sqlStr, user.Id, user.Password, user.Name, user.Age, user.Gender)
	if err != nil {
		chatlog.Error("insert user info to db failed, %v", err)
		return
	}

	userInfoMap.Store(user.GetId(), user)

	imRsp := &interMsg{
		msgName: "protos.SignupResponse",
		msg:     signupRsp,
	}
	c.chanSendMsg <- imRsp
}

func handleLoginReq(pm proto.Message, c *Client) {
	loginRsp := &protos.LoginResponse{
		Result: protos.LoginResult_SUCCESS,
	}
	imRsp := &interMsg{
		msgName: "protos.LoginResponse",
		msg:     loginRsp,
	}

	loginReq, ok := pm.(*protos.LoginRequest) // check msg type
	if !ok {
		chatlog.Warning("login failed")
		loginRsp.Result = protos.LoginResult_INTERERROR
		c.chanSendMsg <- imRsp
		return
	}

	userID := loginReq.GetId()

	user, ok := userInfoMap.Load(userID) // check if user exists
	if !ok {
		loginRsp.Result = protos.LoginResult_USERNOTEXIST
		c.chanSendMsg <- imRsp
		return
	}

	switch user.(type) {
	case *protos.UserInfo:
		userInfo := user.(*protos.UserInfo)
		if userInfo.GetPassword() != loginReq.GetPassword() { // check password
			loginRsp.Result = protos.LoginResult_PSWERROR
		}
	default:
		loginRsp.Result = protos.LoginResult_INTERERROR
	}
	chatlog.Trace("send client msg, msgName = [%v],msg = [%v]", imRsp.msgName, imRsp.msg)

	// check if user is online
	_, ok = userOnlineMap.Load(userID)
	if !ok { // user is offline
		userOnlineMap.Store(userID, "")         // store user to online map
		chanMsgMap.Store(userID, c.chanSendMsg) // store client send msg channel
		c.ID = userID                           // update user id to client
		go c.handleHeartbeat()                  // start hb
	} else {
		loginRsp.Result = protos.LoginResult_USERALDYONLINE
	}

	c.chanSendMsg <- imRsp // reply to client message
}

func handleChatMsg(pm proto.Message, c *Client) {
	chatMsg, ok := pm.(*protos.ChatMsg)

	if !ok {
		chatlog.Warning("chat msg type failed")
		return
	}

	msgType := chatMsg.GetMsgType()
	switch msgType {
	case protos.ChatMsgType_GROUP_CHAT:
		chatlog.Trace("broadcast user...")
		userOnlineMap.Range(func(k, v interface{}) bool {
			if k == c.ID { // 过滤掉自己
				return false
			}
			send, ok := chanMsgMap.Load(k)
			if ok {
				switch send.(type) {
				case chan *interMsg:
					imRsp := &interMsg{
						msgName: "protos.ChatMsg",
						msg:     chatMsg,
					}
					send.(chan *interMsg) <- imRsp
				default:
					chatlog.Warning("chat type failed")
				}
			}
			return true
		})
	case protos.ChatMsgType_PRIVATE_CHAT:
		// from := chatMsg.GetMsgFrom()
		to := chatMsg.GetMsgTo()
		_, ok := userOnlineMap.Load(to)
		if ok {
			send, ok := chanMsgMap.Load(to)
			if ok {
				imRsp := &interMsg{
					msgName: "protos.ChatMsg",
					msg:     chatMsg,
				}
				send.(chan *interMsg) <- imRsp
			}
		} else {
			chatlog.Info("user [%v] is offline \n", to)
		}
	default:
		chatlog.Warning("chat msg type [%v] error\n", msgType)
	}
}

func handleLogoutReq(pm proto.Message, c *Client) {
	logoutRsp := &protos.LogoutResponse{
		Result: protos.LogoutResult_LOGOUTSUCCESS,
	}
	imRsp := &interMsg{
		msgName: "protos.LogoutResponse",
		msg:     logoutRsp,
	}

	logoutReq, ok := pm.(*protos.LogoutRequest) // check msg type
	if !ok {
		chatlog.Warning("logout msg type failed")
		logoutRsp.Result = protos.LogoutResult_LOGOUTINTERERROR
		c.chanSendMsg <- imRsp
		return
	}

	userID := logoutReq.GetId()

	// check if user is online
	_, ok = userOnlineMap.Load(userID)
	if ok { // user is online
		userOnlineMap.Delete(userID) // delete user to online map
		chanMsgMap.Delete(userID)    // delete client send msg channel
		c.ID = 0                     // update user id to client
	} else {
		logoutRsp.Result = protos.LogoutResult_LOGOUTINTERERROR
	}

	c.chanSendMsg <- imRsp // reply to client message
}

func handleHeartbeatMsg(pm proto.Message, c *Client) {
	heartbeatMsg, ok := pm.(*protos.HeartbeatMsg) // check msg type
	if !ok {
		chatlog.Warning("heartbeat msg type failed")
		return
	}
	chatlog.Trace("recv heartbeat msg [%v] from user [%v]", heartbeatMsg.GetMsg(), c.ID)
	c.chanHbMsg <- true // notify heartbeat goroutine

	hbReplyMsg := &protos.HeartbeatReplyMsg{
		Msg: "Ack",
	}
	imRsp := &interMsg{
		msgName: "protos.HeartbeatReplyMsg",
		msg:     hbReplyMsg,
	}

	c.chanSendMsg <- imRsp
}
