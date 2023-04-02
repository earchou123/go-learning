package process

import (
	"encoding/json"
	"fmt"
	"go-learning/chatroom/common/message"
)

//打印群发消息
func outputGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)

	if err != nil {
		fmt.Printf("outputGroupMes json.UnMarshal err=%v\n", err)
		return
	}

	//打印信息
	info := fmt.Sprintf("用户id：%d 对大家说：%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
}
