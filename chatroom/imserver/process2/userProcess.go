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
	UserId int
}
//通知其他用户，我上线了
func (this *UserProcess)NotifyOtherOnlineUser(userId int){
	//遍历在线用户，拿到每个在线用户的指针，通知他们传进来的这个userID的状态
	for id,up :=range userMgr.onlineUsers{
		//如果上线id等于当前登录用户的id，则跳过
		if id == userId{
			continue
		}
		//拿到当前在线的指针up，调用NotifyMeOnline，发送当前登录用户的装修消息
		up.NotifyMeOnline(userId)
	}
}

//发送用户状态消息
func (this *UserProcess)NotifyMeOnline(userId int){
	//userId指要发送的这个人的上线状态。
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	data,err := json.Marshal(notifyUserStatusMes)
	if err != nil{
		fmt.Printf("json.Marshal err=%v\n",err)
		return
	}
	mes.Data = string(data)
	data,err = json.Marshal(mes)
	if err != nil{
		fmt.Printf("json.Marshal err=%v\n",err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err =tf.WritePkg(data)
	if err != nil{
		fmt.Printf("NotifyMeOline err=%v\n",err)
		return
	}
	return
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
	resMes.Type = message.RegisterResMesType

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
		//this.UserId = registerMes.User.UserId
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
		//登录成功，将id添加到userMgr中
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		//通知其他用户，我上线了
		this.NotifyOtherOnlineUser(this.UserId)
		//当前在线用户放入logingResMes.UserId
		//遍历在线用户
		for  id,_ := range userMgr.onlineUsers{
			loginResMes.UsersId = append(loginResMes.UsersId,id)
		}

		fmt.Printf("user=%v\t 登录成功\n", user)
	}

	loginResMes.UserId = this.UserId
	loginResMes.UserName = user.UserName
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
