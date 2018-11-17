package utils

import "fmt"

type FamilyAccount struct {
	/*菜单栏选择key*/
	key string
	/*循环控制变量*/
	loop bool

	/*定义帐号余额*/
	balance float64
	/*每次收支金额*/
	money float64
	/*收支说明*/
	note string
	/*记录收支的行为*/
	flag bool
	/*收支的详情*/
	details string
}

/*工厂模式初始化*/
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t帐号余额\t收支金额\t说明",
	}
}

func (this *FamilyAccount) showDetails() {
	fmt.Println("\n--------------------当前收支明细记录-------------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细。。。来一笔吧！")
	}
}

func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足！！！")
	}
	this.balance -= this.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) exit() {
	fmt.Println("你确定要退出吗？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}
	if choice == "y" {
		this.loop = false
	}
}

func (this *FamilyAccount) MainMenu() {
	/*主菜单的显示*/
	for {
		fmt.Println("\n----------------------家庭收支记账软件---------------------")
		fmt.Println("                        1 收支明细")
		fmt.Println("                        2 登记收入")
		fmt.Println("                        3 登记支出")
		fmt.Println("                        4 退出软件")
		fmt.Println("                   请选择你要执行操作的序号：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项！")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你已经退出了家庭记账软件...")
}
