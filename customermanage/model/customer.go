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
