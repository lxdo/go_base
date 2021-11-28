package main

import (
	"fmt"
	"go_base/11面向对象下/02封装/model"
)

/**
面向对象编程三大特性
	基本介绍
		go仍然有面向对象编程的继承，封装和多态的特性，只是实现的方式和其他oop语言不一样

	封装
		封装的介绍：
		封装就是把抽象出的字段和对字段的操作封装在一起，数据被保护在内部，程序的其他包只有通过被授权的操作(方法)
		才能对字段进行操作
		封装的理解和好处：
		1) 隐藏实现细节
		2) 可以对数据进行验证,保证安全合理
		如何体现封装：
		1) 对结构体中属性进行封装
		2) 通过方法，包 实现封装

	封装的实现步骤:
		1) 将结构体、字段(属性)的首字母小写(不能导出了，其他包不能使用，类似private)
		2) 给结构体所在包提供一个工厂模式的函数，首字母大写。类似一个构造函数
		3) 提供一个首字母大写的Set方法(类似其它语言的public),用于对属性判断并赋值
		func (var 结构体类型名) Setxxx(参数列表) (返回值列表) {  // var 为结构体变量名
			// .... 此处加入数据验证的业务逻辑

			// 对结构体属性进行赋值
			var.字段 = 参数
		}
		4) 提供一个首字母大写的Get方法(类似其他语言的public),用于获取属性的值
		func (var 结构体类型名) Getxxx(){
			return var.字段
		}

		特别说明:在go开发中没有特别强调封装，go本身对面向对象的特性做了简化的

*/

/**
案例
	请大家看一个程序(person.go) 不能随便查看人的年龄，工资等隐私，并对输入的年龄进行合理的验证
*/
func main() {
	// 用工厂模式 函数获得一个指向结构体变量数据空间的指针变量
	p := model.NewPerson("smith")
	fmt.Println(*p)
	// 调用方法给结构体私有字段赋值
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(*p)

	// 调用方法读取结构体私有字段
	fmt.Println(p.Name, "age=", p.GetAge(), "sal=", p.GetSal())

	/**
	案例
		创建程序 在model包中定义account结构体，在main函数中使用
		1) account结构体要求具有字段:账户(长度在6-10之间) 、余额必要>20、密码(必须6位)
		2) 通过Setxxx的方法给account的字段赋值
		3) 在main函数中测试
	*/

	account := model.NewAccount("zhang111", "888888", 40)

	if account != nil {
		fmt.Println("create account success",*account)
	}else{
		fmt.Println("create account fail")
	}

}
