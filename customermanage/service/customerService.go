package service

import (
	"customermanage/model"
)

type CustomerService struct {
	customers   []model.Customer
	customerNum int //id
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "111", "zs@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// 返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

// 添客户到customers切片
func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum++ // id增加
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

// 从customers切片删除用户
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)

	if index == -1 {
		return false // 没找到
	}
	// 删除对应的切片
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

// 根据下标查找切片中的对应下标，如果没有该客户，则返回-1
func (this *CustomerService) FindById(id int) int {
	// 没找到则返回-1
	index := -1
	// 遍历切片
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			// 找到啦
			index = i

		}
	}
	return index
}
