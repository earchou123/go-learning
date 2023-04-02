package process2

import "fmt"

var userMgr *UserMgr

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//初始化在线用户列表
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess),
	}
}

//添加在线用户
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

//删除在线用户
func (this *UserMgr) DeleteOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

//查询在线用户列表
func (this *UserMgr) GetOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

//通过id查询用户
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%不存在\n", userId)
		return
	}
	return
}
