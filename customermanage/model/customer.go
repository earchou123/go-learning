package model

import "fmt"

type Customer struct {
	id     int
	name   string
	age    int
	gender string
	phone  string
	email  string
}

func NewCustomer(id int, name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		id:     id,
		name:   name,
		gender: gender,
		age:    age,
		phone:  phone,
		email:  email,
	}
}

func (this *Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.id, this.name, this.age, this.gender, this.phone, this.email)
	return info
}

// func (this *NewCustomer) Add() {
// 	fmt.Println("------------------添加客户-----------------")
// 	this.id = len(this) + 1
// 	fmt.Println("请输入姓名：")
// 	fmt.Scanln(&this.name)
// 	fmt.Println("请输入性别：")
// 	fmt.Scanln(&this.gender)
// 	fmt.Println("请输入年龄：")
// 	fmt.Scanln(&this.age)
// 	fmt.Println("请输入电话：")
// 	fmt.Scanln(&this.phone)
// 	fmt.Println("请输入邮箱：")
// 	fmt.Scanln(&this.email)
// 	fmt.Println("------------------添加成功✅-----------------")
// }
// func (this *NewCustomer) Edit() {
// 	fmt.Println("------------------修改客户-----------------")
// 	fmt.Println("请输入待修改的客户ID：")
// 	fmt.Scanln(&this.id)
// 	fmt.Printf("姓名（%v）：（	直接回车表示不修改）\n", this.name)
// 	fmt.Scanln(&this.name)
// 	fmt.Printf("性别(%v)：\n", this.gender)
// 	fmt.Scanln(&this.gender)
// 	fmt.Printf("年龄(%v)：\n", this.age)
// 	fmt.Scanln(&this.age)
// 	fmt.Printf("电话(%v)：\n", this.phone)
// 	fmt.Scanln(&this.phone)
// 	fmt.Printf("邮箱(%v)：\n", this.email)
// 	fmt.Scanln(&this.email)
// 	fmt.Println("------------------修改成功✅-----------------")
// }

// func (this *NewCustomer) Delete() {
// 	fmt.Println("-----------------删除客户-----------------")
// 	fmt.Println("请要删除的的客户ID：")
// 	fmt.Scanln(&this.id)
// 	fmt.Println("------------------删除成功✅-----------------")
// }
