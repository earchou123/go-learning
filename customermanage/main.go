package main

import "customermanage/view"
import "customermanage/service"

func main() {
	customerView := view.CustomerView{
		Key:  "",
		Loop: true}
	customerView.CustomerService = service.NewCustomerService()

	customerView.MainMenu()
}
