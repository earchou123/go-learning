package process

import (
	"encoding/json"
	"fmt"
	"go-learning/chatroom/common/message"
	"go-learning/chatroom/imserver/utils"
	"net"
)

func ShowMenu(loginMes *message.LoginResMes) {
	loop := true
	for loop {
		fmt.Printf("------------恭喜%v[ID:%d]登录成功------------\n",loginMes.UserName,loginMes.UserId)
		fmt.Printf("%20v\n", "1 显示在线用户列表")
		fmt.Printf("%20v\n", "2 发送消息")
		fmt.Printf("%20v\n", "3 信息列表")
		fmt.Printf("%20v\n", "4 退出系统")
		fmt.Println("请选择（1-4）")

		var key int
		fmt.Scanf("%d\n", &key)

		var content string
		switch key {
		case 1:
			outputOnlineUser()
		case 2:
			fmt.Println("发送消息")
			fmt.Scanf("%v\n",&content)
			var smsProcess SmsProcess
			_ = smsProcess.SendGroupMes(content)
		case 3:
			fmt.Println("信息列表")
		case 4:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("您输入的选项不正确，请重新输入")
		}
	}
}

func serverProcessMes(conn net.Conn) {
	tf := utils.Transfer{
		Conn: conn,
	}

	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Printf("ReadPkg err=%v\n", err)
			return
		}
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data),&notifyUserStatusMes)
			//更新用户状态
			updataUserStatus(&notifyUserStatusMes)
		default:
			fmt.Println("服务器返回了客户端无法识别的消息类型。")
		}
		//fmt.Printf("mess=%v\n", mes)
	}
}
