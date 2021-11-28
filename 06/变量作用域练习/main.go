package main

import "fmt"

/**
当全局变量和局部变量名相同,并且都在作用域内，会采用就近原则
*/
var name = "a" //全局变量

func test01() {
	fmt.Println(name)
}
func test02() {
	name := "b" // 局部变量 当全局变量和局部变量名相同,并且都在作用域内，会采用就近原则
	fmt.Println(name)
}

var Age int = 20
// Name := "tom" 错误  要用var 方式定义赋值全局变量
// Name := "tom" =>等价于 var Name string  Name="tom" 赋值语句不能在函数体外执行
func main() {
	fmt.Println(name) // a 当全局变量和局部变量名相同,并且都在作用域内，会采用就近原则
	test01()          // a
	test02()          // b 就近原则
	test01()          // a 局部变量不会影响全局变量

	fmt.Println("Age=",Age,"Name=",)
}



