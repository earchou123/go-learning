package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-learning/chatroom/common/message"
	"go-learning/chatroom/imserver/utils"
	"net"
)

type UserProcess struct {
	//
}

//注册
func (this *UserProcess) Register(userId int, userName string, userPwd string) (err error) {
	// 连接服务器
	conn, err := net.Dial("tcp", "0.0.0.0:20000")
	if err != nil {
		fmt.Println("net.Dial err=%v\n", err)
		return
	}
	// 延时关闭
	defer conn.Close()

	var mes message.Message
	mes.Type = message.RegisterMesType

	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	// 将RregisterMes序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=%v", err)
		return
	}
	// 将data赋值给mes.Data
	mes.Data = string(data)

	// 将mes序列化
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
	// 读取服务器返回的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Printf("readPkg err=%v\n", err)
		return
	}
	// 将mes的Data反序列化
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Printf("json.Unmarshal err=%v\n", err)
		return
	}
	if registerResMes.Code == 200 {
		fmt.Printf("注册成功，请重新登录\n")
		//
	} else {
		fmt.Printf("注册失败：%v\n", registerResMes.Error)
	}
	return
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
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
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
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
		//CurUser初始化
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline
		// 展示在线用户列表，遍历loginResMes.UsersId
		fmt.Printf("------------当前在线用户列表------------\n")
		for _,v := range loginResMes.UsersId{
			fmt.Printf("ID：%v\n",v)

			//客户端在线用户初始化
			user := &message.User{
				UserId: v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Printf("------------end------------\n")

		//接收服务端发送的消息
		go serverProcessMes(conn)

		//显示菜单
		ShowMenu(&loginResMes)

	} else {
		fmt.Printf("登录失败：%v\n", loginResMes.Error)
	}

	return

}
