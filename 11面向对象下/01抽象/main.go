package main

import "fmt"

/**
面向对象编程思想-抽象

	我们在前面去定义一个结构体时候,实际上就是把一类事物的共有属性(字段)和行为(方法)提取出来
	形成一个物理模型(模板)。这种研究问题的方法称为抽象


*/

// 定义一个结构体
// Account银行账户
type Account struct {
	// 共有属性
	Name    string  // 账号名
	Pwd     int     // 密码
	Balance float64 // 余额
}

// 共有方法
// 存款
func (account *Account) Deposite(money float64, pwd int) {
	// 验证密码是否正确
	if pwd != account.Pwd {
		fmt.Println("pwd error")
		return
	}
	// 看看存款金额是否正确
	if money <= 0 {
		fmt.Println("money error")
		return
	}
	// 存款成功
	account.Balance += money
	fmt.Println("deposite success")
}

// 取款
func (account *Account) Withdraw(money float64, pwd int) {
	// 验证密码是否正确
	if pwd != account.Pwd {
		fmt.Println("pwd error")
		return
	}
	// 取款不能小于等于0 或者 不能取超过存款
	if money <= 0 || money > account.Balance {
		fmt.Println("money error")
		return
	}
	// 取款成功
	account.Balance -= money
	fmt.Println("withDraw success")
}

// 查询余额
func (account *Account) Query(pwd int) {
	// 验证密码是否正确
	if pwd != account.Pwd {
		fmt.Println("pwd error")
		return
	}
	fmt.Println("balance=", account.Balance)
}
func main() {
	// 创建结构体实例

	account := Account{
		Name:    "zhang", // 结构体字段赋值
		Pwd:     666666,
		Balance: 100.0,
	}

	// 查询余额
	account.Query(666666)

	// 存款
	account.Deposite(2000,666666)
	account.Query(666666)

	// 取款
	account.Withdraw(1000,666666)
	account.Query(666666)


}
