package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-learning/chatroom/common/message"
	"net"
)

func login(userId int, userPwd string) (err error) {
	// fmt.Printf("userId=%v,userPwd=%s\n", userId, userPwd)
	// return nil
	conn, err := net.Dial("tcp", "0.0.0.0:20000")
	if err != nil {
		fmt.Println("net.Dial err=%v\n", err)
		return
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=%v", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=%v", err)
		return
	}

	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte

	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Printf("conn.Write err=%v\n", err)
		return
	}
	fmt.Printf("客户端发送的消息长度=%d 内容=%v\n", len(data), string(data))
	return

}
