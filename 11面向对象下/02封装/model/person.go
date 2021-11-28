package model

import (
	"fmt"
)

/**
案例
	请大家看一个程序(person.go) 不能随便查看人的年龄，工资等隐私，并对输入的年龄进行合理的验证
*/

// 定义person结构体
type person struct { // 其他包无法直接使用person结构体
	Name string  // 其他包可直接访问
	age  int     // 其他不包可直接访问 私有字段
	sal  float64 // 其他不包可直接访问 私有字段
}

// 写一个工厂模式的函数，相当于构造函数 返回指向结构体变量数据空间的指针变量
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// 给person绑定一个方法 为了给结构体中私有字段赋值 age
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("age error")
	}
}

// 给person绑定一个方法 为了访问结构体中私有字段 age
func (p *person) GetAge() int {
	return p.age
}

// 给person绑定一个方法 为了给结构体中私有字段赋值 sal
func (p *person) SetSal(sal float64) {
	if sal >= 3000 && sal <= 30000 {
		p.sal = sal
	} else {
		fmt.Println("sal error")
	}
}

// 给person绑定一个方法 为了访问结构体中私有字段 sal
func (p *person) GetSal() float64 {
	return p.sal
}

/**
案例
	创建程序 在model包中定义account结构体，在main函数中使用
	1) account结构体要求具有字段:账户(长度在6-10之间) 、余额必要>20、密码(必须6位)
	2) 通过Setxxx的方法给account的字段赋值
	3) 在main函数中测试
*/
// 定义一个私有结构体
// account银行账户 
type account struct {
	// 共有属性
	accountNo string  // 账户
	pwd       string  // 密码
	balance   float64 // 余额
}

// 工厂模式的函数-构造函数  返回指向结构体变量数据空间的指针变量
func NewAccount(accountNo string, pwd string, balance float64) *account {
	// 账户不能小于6位 不能大于10位
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("accountNo error")

		return nil
	}
	// 密码长度6位
	if len(pwd) != 6 {
		fmt.Println("pwd error")
		return nil
	}

	// 余额不能小于20
	if balance < 20 {
		fmt.Println("balance error")
		return nil
	}

	// 返回指向结构体变量数据空间的指针变量
	return &account{
		accountNo: accountNo,
		pwd:       pwd,
		balance:   balance,
	}

}

// 结构体绑定设置密码方法
func (account *account) SetPwd(pwd string) {
	if len(pwd) != 6 {
		fmt.Println("pwd error")
		return
	}
	account.pwd = pwd
}

// 结构体绑定获取密码方法
func (account *account) GetPwd() string {
	return account.pwd
}

// 共有方法
// 存款
func (account *account) Deposite(money float64, pwd string) {
	// 验证密码是否正确
	if pwd != account.GetPwd() {
		fmt.Println("pwd error")
		return
	}
	// 看看存款金额是否正确
	if money <= 0 {
		fmt.Println("money error")
		return
	}
	// 存款成功
	account.balance += money
	fmt.Println("deposite success")
}

// 取款
func (account *account) Withdraw(money float64, pwd string) {
	// 验证密码是否正确
	if pwd != account.GetPwd() {
		fmt.Println("pwd error")
		return
	}
	// 取款不能小于等于0 或者 不能取超过存款
	if money <= 0 || money > account.balance {
		fmt.Println("money error")
		return
	}
	// 取款成功
	account.balance -= money
	fmt.Println("withDraw success")
}

// 查询余额
func (account *account) Query(pwd string) {
	// 验证密码是否正确
	if pwd != account.GetPwd() {
		fmt.Println("pwd error")
		return
	}
	fmt.Println("balance=", account.balance)
}
