package process

import (
	"encoding/json"
	"fmt"
	"go-learning/chatroom/common/message"
	"go-learning/chatroom/imserver/utils"
)

type SmsProcess struct {

}

func (this *SmsProcess) SendGroupMes(content string)  (err error){
	var mes  message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes

	smsMes.Content = content
	smsMes.User = CurUser.User
	data,err := json.Marshal(smsMes)
	if err != nil{
		fmt.Printf("SendGroupMes json.Marshal err=%v\n",err)
		return
	}
	mes.Data = string(data)
	data ,err = json.Marshal(mes)
	if err !=nil {
		fmt.Printf("SendGroupMes json.Marshal err=%v\n",err)
		return
	}

	//发送SendGroupMes给服务器
	tf :=&utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil{
		fmt.Printf("SendGroupMes WritePkg err=%v\n",err)
		return
	}

	return

}