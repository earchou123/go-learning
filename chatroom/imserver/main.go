package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-learning/chatroom/common/message"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Printf("conn.Read fail err=%v\n", err)

	}

	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Printf("json.Unmarshal err=%v\n", err)
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	// 发送包长度

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

	// 发送data本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Printf("conn.Write err=%v\n", err)
		return
	}
	return
}

// 处理登录请求
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	// 先从mes中读取mes.Data，并反序列化程LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Printf("josn.Unmarshal err=%v\n", err)
		return
	}

	// 声明一个Mes作为返回消息
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes

	// 判断用户id和密码
	if loginMes.UserId == 1 && loginMes.UserPwd == "123" {
		loginResMes.Code = 200

	} else {
		loginResMes.Code = 500
		loginResMes.Error = "用户不存在，请注册再使用"
	}

	// 序列化loginResMes
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Printf("json.Marshal err=%v\n", err)
		return
	}

	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Printf("json.Marshal err=%v\n", err)
		return
	}

	err = writePkg(conn, data)
	return
}

// 根据客户端发送的消息种类不同，决定调用哪个函数来处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// login
		err = serverProcessLogin(conn, mes)
	case message.LoginResMesType:
		// response
	default:
		fmt.Println("消息类型不存在")
	}
	return
}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		fmt.Println("读取客户端发送的数据。。。")
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端关闭了链接。")
				return
			} else {
				fmt.Printf("readPkg err=%v\n", err)
				return
			}
		}
		fmt.Printf("读到的mes=%v\n", mes)
		err = serverProcessMes(conn, &mes)
		if err != nil {
			return
		}
	}
}

func main() {
	fmt.Println("服务器在监听20000端口")
	listen, err := net.Listen("tcp", "0.0.0.0:20000")
	defer listen.Close()
	if err != nil {
		fmt.Printf("net.Listen err=%v\n", err)
		return
	}
	for {
		fmt.Println("等待客户端连接服务器..")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("listen.Accept err=%v\n", err)
		}
		go process(conn)
	}

}
