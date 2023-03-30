package process2

import (
	"encoding/json"
	"fmt"
	"go-learning/chatroom/common/message"
	"go-learning/chatroom/imserver/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

// 处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}