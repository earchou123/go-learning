package process

import (
	"fmt"
	"go-learning/chatroom/imserver/utils"
	"net"
)

func ShowMenu() {
	loop := true
	for loop {
		fmt.Println("------------恭喜XXX登录成功------------")
		fmt.Printf("%20v\n", "1 显示在线用户列表")
		fmt.Printf("%20v\n", "2 发送消息")
		fmt.Printf("%20v\n", "3 信息列表")
		fmt.Printf("%20v\n", "4 退出系统")
		fmt.Println("请选择（1-4）")

		var key int
		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("显示在线用户列表")
		case 2:
			fmt.Println("发送消息")
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
		fmt.Printf("mess=%v\n", mes)
	}
}
