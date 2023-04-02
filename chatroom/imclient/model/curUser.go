package model

import (
	"go-learning/chatroom/common/message"
	"net"
)

//当前登录用户
type CurUser struct {
	Conn net.Conn
	message.User
}
