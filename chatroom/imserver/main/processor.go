package main

import (
	"fmt"
	"go-learning/chatroom/common/message"
	"go-learning/chatroom/imserver/process2"
	"go-learning/chatroom/imserver/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 根据客户端发送的消息种类不同，决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// login
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//  register
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	default:
		fmt.Printf("消息类型不存在,mes=%v\n",mes)
	}
	return
}

func (this *Processor) process2() (err error) {

	for {
		fmt.Println("读取客户端发送的数据。。。")

		tf := &utils.Transfer{
			Conn: this.Conn,
		}

		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端关闭了链接。")
				return err
			} else {
				fmt.Printf("readPkg err=%v\n", err)
				return err
			}
		}
		fmt.Printf("读到的mes=%v\n", mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
