package main

import "familyaccount/utils"

func main() {
	utils.NewFamilyAccount().Login()
}

// --------------------------
// import (
// 	"fmt"
// )

// func main() {
// 	//	保存用户输入的变量
// 	key := ""
// 	// 控制是否退出循环
// 	loop := true
// 	// 账户余额
// 	balance := 10000.0
// 	// 每次收支金额
// 	money := 100.0
// 	// 收支说明
// 	note := ""
// 	// 收支详情字符串
// 	details := "收支\t账户金额\t收支金额\t说明"
// 	// 收支明细是否为空
// 	flag := false

// 	for {
// 		fmt.Println("----------------家庭收支记账软件----------------")
// 		fmt.Printf("%25s\n", "1 收支明细")
// 		fmt.Printf("%25s\n", "2 登记收入")
// 		fmt.Printf("%25s\n", "3 登记支出")
// 		fmt.Printf("%25s\n", "4 退出软件")
// 		fmt.Println("请选择（1-4）")
// 		fmt.Scanln(&key)

// 		switch key {
// 		case "1":
// 			fmt.Println("----------------当前收支明细记录----------------")
// 			if !flag {
// 				fmt.Println("暂无收支明细")
// 			} else {
// 				fmt.Println(details)
// 			}
// 		case "2":
// 			fmt.Println("----------------登记收入----------------")
// 			fmt.Println("请输入本次收入金额：")
// 			fmt.Scanln(&money)
// 			balance += money
// 			fmt.Println("本次收入说明：")
// 			fmt.Scanln(&note)
// 			details += fmt.Sprintf("\n收入\t%8v\t%8v\t%v", balance, money, note)
// 			flag = true

// 		case "3":
// 			fmt.Println("----------------登记支出----------------")
// 			fmt.Println("请输入本次支出金额：")
// 			fmt.Scanln(&money)
// 			if money > balance {
// 				fmt.Println("余额不足！")
// 				break
// 			}

// 			balance -= money
// 			fmt.Println("本次支出说明：")
// 			fmt.Scanln(&note)
// 			details += fmt.Sprintf("\n支出\t%v\t-%v\t%v", balance, money, note)
// 			flag = true

// 		case "4":
// 			fmt.Println("确定要退出吗？y/n")
// 			choice := ""
// 			for {
// 				fmt.Scanln(&choice)
// 				if choice == "y" || choice == "n" {
// 					break
// 				}
// 				fmt.Println("请输入y（退出） 或 n（继续）")
// 			}
// 			if choice == "y" {
// 				loop = false
// 			}
// 		default:
// 			fmt.Println("请输入正确的选项....")

// 		}
// 		if !loop {
// 			break
// 		}
// 	}
// 	fmt.Println("退出家庭记账软件！")
// }
