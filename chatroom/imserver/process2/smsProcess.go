package process2

import (
	"encoding/json"
	"fmt"
	"go-learning/chatroom/common/message"
	"go-learning/chatroom/imserver/utils"
	"net"
)

type SmsProcess struct {
}

//群发消息
func (this SmsProcess) SendGroupMes(mes *message.Message) (err error) {
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Printf("SendGroupMes json.Unmarshal err=%v\n", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Printf("SendGroupMes json.Marshal err=%v\n", err)
		return
	}

	//遍历当前在线用户列表，分别给他们发消息
	for id, up := range userMgr.onlineUsers {
		//过滤当前登录用户
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
	return
}

//发消息给用户
func (this SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) (err error) {
	//发送消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	return err
}
