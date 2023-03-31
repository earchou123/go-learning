package model

import (
	"context"
	"encoding/json"
	"fmt"
	"go-learning/chatroom/common/message"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var MyUserDao *UserDao

type UserDao struct {
	client *redis.Client
}

func NewUserDao(client *redis.Client) (userDao *UserDao) {
	userDao = &UserDao{
		client: client,
	}
	return
}

// 根据一个用户id返回一个user实例
func (this *UserDao) getUserById(id int) (user *User, err error) {
	// 通过给定的id ，去redis查询该用户
	jsonStr, err := this.client.HGet(context.Background(), "users", strconv.Itoa(id)).Result()
	if err != nil {
		if err == redis.Nil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}
	err = json.Unmarshal([]byte(jsonStr), user)
	if err != nil {
		fmt.Printf("json.Unmarshal err=%v\n", err)
	}
	return
}

// 增加一个用户
func (this *UserDao) setUser(key int, value string) (err error) {
	// redis添加用户
	err = this.client.HSet(context.Background(), "users", key, value).Err()
	if err != nil {
		fmt.Printf("添加用户失败: err=%$v\n", err)
		return
	}
	return
}

//登录
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	user, err = this.getUserById(userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}

	return
}

func (this *UserDao) Register(user *message.User) (err error) {
	//先查询，如果用户存在，返回
	_, err = this.getUserById(user.UserId)
	if err == nil {
		if err != ERROR_USER_NOTEXISTS {
			err = ERROR_USER_TEXISTED
			return
		}
	}
	//接收mes对象，把序列化，再转成字符串传给setUser
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("json Marshal err=%v\n", err)
	}
	//调用setUser，添加用户
	err = this.setUser(user.UserId, string(data))
	return
}
