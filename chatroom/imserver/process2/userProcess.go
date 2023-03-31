package process2

import (
	"encoding/json"
	"fmt"
	"go-learning/chatroom/common/message"
	"go-learning/chatroom/imserver/model"
	"go-learning/chatroom/imserver/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

//处理注册请求
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	// 先从mes中读取mes.Data，并反序列化程RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Printf("josn.Unmarshal err=%v\n", err)
		return
	}

	// 声明一个Mes作为返回消息
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var registerResMes message.RegisterResMes

	//注册用户
	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		if err == model.ERROR_USER_TEXISTED {
			registerResMes.Code = 403
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 505
			registerResMes.Error = "服务器内部错误。。"
		}
	} else {
		registerResMes.Code = 200
		fmt.Printf("注册成功\n")
	}

	// 序列化registerResMes
	data, err := json.Marshal(registerResMes)
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

	//发送response给客户端
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
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

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误。。"
		}
	} else {
		loginResMes.Code = 200
		fmt.Printf("user=%v\t 登录成功\n", user)
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
	//发送response给客户端
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
