package main

import (
	"fmt"
	"go_code/chapter04-struct/dome8-abstract/model"
)

type account struct {
	accountNo string
	pwd       string
	balance   float64
}

func NewAccount(accountNo string, pwd string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("帐号的长度不正确！")
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("密码的长度不对！")
		return nil
	}
	if balance < 20 {
		fmt.Println("余额数目不正确！")
		return nil
	}
	return &account{
		accountNo: accountNo,
		pwd:       pwd,
		balance:   balance,
	}
}

/*存款*/
func (account *account) Deposite(money float64, pwd string) {
	if pwd != account.pwd {
		fmt.Println("密码错误")
		return
	}
	if money <= 0 {
		fmt.Println("输入的金额错误！")
	}
	account.balance += money
	fmt.Println("存款成功！")
}

/*取款*/
func (account *account) WithDraw(money float64, pwd string) {
	if pwd != account.pwd {
		fmt.Println("密码错误")
		return
	}
	if money <= 0 || money > account.balance {
		fmt.Println("输入的金额错误！")
	}
	account.balance -= money
	fmt.Println("取款成功！")
}

/*查询余额*/
func (account *account) Query(pwd string) {
	if pwd != account.pwd {
		fmt.Println("密码错误")
		return
	}
	fmt.Println(account.accountNo, "的帐号余额为：", account.balance)
}

func main() {
	account := &account{
		accountNo: "ICBC",
		pwd:       "123",
		balance:   100,
	}
	account.Query("123")
	account.Deposite(100, "123")
	account.Query("123")
	account.WithDraw(100, "123")
	account.Query("123")

	p := model.NewPerson("TGH")
	fmt.Println(*p)
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(*p)
	fmt.Println(p.Name, "的年龄是：", p.GetAge(), "的薪水是：", p.GetSal())
}
