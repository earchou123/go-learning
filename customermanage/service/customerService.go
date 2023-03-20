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
