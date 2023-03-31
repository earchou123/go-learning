package model

import (
	"context"
	"encoding/json"
	"fmt"
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
