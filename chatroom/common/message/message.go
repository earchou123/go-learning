package message

const (
	LoginMesType            = "LoginMes"            //登录消息
	LoginResMesType         = "LoginResMes"         //登录返回消息
	RegisterMesType         = "RegisterMes"         //注册消息
	RegisterResMesType      = "RegisterResMes"      //注册返回消息
	NotifyUserStatusMesType = "NotifyUserStatusMes" // 通知用户状态消息
	SmsMesType              = "SmsMes"              // 群发消息
)

//用户状态
const (
	UserOnline  = iota // 在线
	UserOffline        // 不在线
	UserBusy           // 忙碌
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"`
}

// 登录信息
type LoginMes struct {
	UserId   int    `json:"userId"`   // 用户id
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` // 用户名
}

// 登录返回消息
type LoginResMes struct {
	Code     int    `json:"code"`
	UsersId  []int  `json:"userIds"` //保存用户切片
	UserName string // 返回当前用户名
	UserId   int    // 返回当前用户id
	Error    string `json:"error"`
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `josn:"code"`
	Error string `json:"error"`
}

//推送用户状态消息
type NotifyUserStatusMes struct {
	UserId int `json:"userID"`
	Status int `json:"status"` // 用户在线状态
}

//发送消息
type SmsMes struct {
	Content string `json:"content"`
	User
}
