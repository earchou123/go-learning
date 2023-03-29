package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-learning/chatroom/common/message"
	"net"
)

func login(userId int, userPwd string) (err error) {
	// 连接服务器
	conn, err := net.Dial("tcp", "0.0.0.0:20000")
	if err != nil {
		fmt.Println("net.Dial err=%v\n", err)
		return
	}
	// 延时关闭
	defer conn.Close()

	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// 将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=%v", err)
		return
	}
	// 将data赋值给mes.Data
	mes.Data = string(data)

	// 将loginMes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=%v", err)
		return
	}

	// -----发消息-------
	// 定义消息长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	// 将消息长度在转成字节
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	// 发送消息长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Printf("conn.Write err=%v\n", err)
		return
	}
	// fmt.Printf("客户端发送的消息长度=%d 内容=%v\n", len(data), string(data))
	// 发送消息本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Printf("conn.Write err=%v\n", err)
		return
	}
	mes, err = readPkg(conn)
	if err != nil {
		fmt.Printf("readPkg err=%v\n", err)
		return
	}
	// 将mes的Data反序列化
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Printf("json.Unmarshal err=%v\n", err)
		return
	}
	if loginResMes.Code == 200 {
		fmt.Printf("登录成功\n")
	} else if loginResMes.Code == 500 {
		fmt.Printf("登录失败：%v\n", loginResMes.Error)
	}

	return

}
