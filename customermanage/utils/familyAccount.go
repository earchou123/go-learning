package utils

import "fmt"

type FamilyAccount struct {
	balance float64
	money   float64
	note    string
	loop    bool
	details string
	flag    bool
	key     string
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		balance: 10000.00,
		money:   0.0,
		note:    "",
		loop:    true,
		details: "收支\t账户余额\t收支金额\t 说 明",
		flag:    false,
		key:     "",
	}
}

func (this *FamilyAccount) showDetails() {
	fmt.Println("----------------当前收支明细记录----------------")
	if !this.flag {
		fmt.Println("暂无收支明细")
	} else {
		fmt.Println(this.details)
	}
}

func (this *FamilyAccount) income() {
	fmt.Println("----------------登记收入----------------")
	fmt.Println("请输入本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%8v\t%8v\t%5v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) outcome() {
	fmt.Println("----------------登记支出----------------")
	fmt.Println("请输入本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足！")
		// break
	} else {
		this.balance -= this.money
		fmt.Println("本次支出说明：")
		fmt.Scanln(&this.note)
		this.details += fmt.Sprintf("\n支出\t%8v\t%8v\t%5v", this.balance, this.money, this.note)
		this.flag = true
	}
}

func (this *FamilyAccount) exit() {
	fmt.Println("确定要退出吗？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("请输入y（退出） 或 n（继续）")
	}
	if choice == "y" {
		this.loop = false
	}
}

func (this *FamilyAccount) transfer() {
	fmt.Println("----------------转账----------------")
	fmt.Println("请输入转账金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足！")
	} else {
		this.balance -= this.money
		fmt.Println("输入转账对象：")
		fmt.Scanln(&this.note)
		this.details += fmt.Sprintf("\n转账\t%8v\t%8v\t%5v", this.balance, this.money, "转账给"+this.note)
		this.flag = true
	}
}

func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("----------------家庭收支记账软件----------------")
		fmt.Printf("%25s\n", "1 收支明细")
		fmt.Printf("%25s\n", "2 登记收入")
		fmt.Printf("%25s\n", "3 登记支出")
		fmt.Printf("%23s\n", "4 转账")
		fmt.Printf("%25s\n", "5 退出软件")
		fmt.Println("请选择（1-5）")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()

		case "3":
			this.outcome()
		case "4":
			this.transfer()
		case "5":
			this.exit()
		default:
			fmt.Println("请输入正确的选项....")

		}
		if !this.loop {
			break
		}
	}
	fmt.Println("退出家庭记账软件！")
}
