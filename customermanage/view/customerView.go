package view

import (
	"customermanage/model"
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

// 显示客户列表
func (this *CustomerView) list() {
	customers := this.CustomerService.List()
	fmt.Println("------------------客户列表-----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("------------------客户列表end-----------------")
}

// 添加客户
func (this *CustomerView) add() {
	fmt.Println("------------------添加客户-----------------")
	fmt.Printf("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Printf("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Printf("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Printf("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Printf("邮件：")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name, gender, age, phone, email)

	if this.CustomerService.Add(customer) {
		fmt.Println("------------------添加完成-----------------")
	} else {
		fmt.Println("------------------添加失败-----------------")
	}
}

// 删除客户
func (this *CustomerView) delete() {
	fmt.Println("------------------删除客户-----------------")
	fmt.Println("请输入要删除的客户ID(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return // 放弃删除
	}
	customers := this.CustomerService.List()
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	fmt.Println(customers[id-1].GetInfo())

	fmt.Println("确定要删除吗？y/n")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if this.CustomerService.Delete(id) {
			fmt.Println("------------------删除完成-----------------")
		} else {
			fmt.Println("------------------删除失败-----------------")
		}
	}

}

// 编辑修改客户
func (this *CustomerView) edit() {
	fmt.Println("------------------修改客户-----------------")
	fmt.Printf("请输入要修改的客户ID(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return // 放弃修改
	}
	index := this.CustomerService.FindById(id)
	if index == -1 {
		fmt.Println("用户id不存在")
		return
	}
	customer := &(this.CustomerService.List())[id-1]
	fmt.Printf("姓名(%v)：<直接回车则不修改>", customer.Name)
	name := ""
	fmt.Scanln(&name)
	fmt.Printf("性别(%v)：", customer.Gender)
	gender := ""
	fmt.Scanln(&gender)
	fmt.Printf("年龄(%v)：", customer.Age)
	age := 0
	fmt.Scanln(&age)
	fmt.Printf("电话(%v)：", customer.Phone)
	phone := ""
	fmt.Scanln(&phone)
	fmt.Printf("邮件(%v)：", customer.Email)
	email := ""
	fmt.Scanln(&email)

	if this.CustomerService.Edit(id, name, gender, age, phone, email) {
		fmt.Println("\n编号\t姓名\t性别\t年龄\t电话\t邮箱")
		fmt.Println(customer.GetInfo())
		fmt.Println("------------------修改完成-----------------")
	} else {
		fmt.Println("------------------修改失败-----------------")
	}
}

// 菜单
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
			this.add()
		case "2":
			this.edit()
		case "3":
			this.delete()

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
