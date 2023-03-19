package view

import (
	"customermanage/service"
	"fmt"
)

type CustomerView struct {
	Key             string
	Loop            bool
	CustomerService *service.CustomerService
}

func NewCustomerView(key string, loop bool) *CustomerView {
	return &CustomerView{
		Key:  key,
		Loop: loop,
	}
}
func (this *CustomerView) list() {
	customers := this.CustomerService.List()
	fmt.Println("------------------客户列表-----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("------------------客户列表end-----------------")
}

func (this *CustomerView) MainMenu() {
	for this.Loop {
		fmt.Println("------------------客户信息管理系统-----------------")
		fmt.Printf("%27v\n", "1 添加客户")
		fmt.Printf("%27v\n", "2 修改客户")
		fmt.Printf("%27v\n", "3 删除客户")
		fmt.Printf("%27v\n", "4 客户列表")
		fmt.Printf("%29v\n", "5 退    出")
		fmt.Println("请选择1-5")

		fmt.Scanln(&this.Key)
		switch this.Key {
		case "1":
			fmt.Println("添加客户")
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")

		case "4":
			this.list()

		case "5":
			this.Loop = false

		default:
			fmt.Println("请输入1-5")
		}
	}
	fmt.Println("退出系统.")
}
