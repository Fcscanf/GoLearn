package main

import "fmt"

func main() {
	/*菜单栏选择key*/
	var key string
	/*循环控制变量*/
	loop := true

	/*定义帐号余额*/
	balance := 10000.0
	/*每次收支金额*/
	money := 0.0
	/*收支说明*/
	note := ""
	/*记录收支的行为*/
	flag := false
	/*收支的详情*/
	details := "收支\t帐号余额\t收支金额\t说明"
	/*主菜单的显示*/
	for {
		fmt.Println("\n----------------------家庭收支记账软件---------------------")
		fmt.Println("                        1 收支明细")
		fmt.Println("                        2 登记收入")
		fmt.Println("                        3 登记支出")
		fmt.Println("                        4 退出软件")
		fmt.Println("                   请选择你要执行操作的序号：")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("\n--------------------当前收支明细记录-------------------")
			if flag {
				fmt.Println(details)
			} else {
				fmt.Println("当前没有收支明细。。。来一笔吧！")
			}
		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
			flag = true
		case "3":
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足！！！")
				break
			}
			balance -= money
			fmt.Println("本次支出说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
			flag = true
		case "4":
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
				loop = false
			}
		default:
			fmt.Println("请输入正确的选项！")
		}
		if !loop {
			break
		}
	}
	fmt.Println("你已经退出了家庭记账软件...")
}
