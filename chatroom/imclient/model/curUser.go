package model

import (
	"go-learning/chatroom/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}