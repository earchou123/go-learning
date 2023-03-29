package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"`
}

// 登录信息
type LoginMes struct {
	UserId   int    `json:"userId"`   // 用户id
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"username"` // 用户名
}

// 登录返回消息
type LoginResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
