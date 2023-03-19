package service

import (
	"customermanage/model"
)

type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "111", "zs@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}
