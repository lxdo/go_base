package main

import "fmt"

/**
逻辑运算符

介绍
	用于连接多个条件(一般来讲就是关系表达式)，最终的结果也是一个bool值
*/
func main() {
	var age int = 40
	// && 用法
	// true && true = true
	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}

	// true && false = false
	if age > 30 && age < 40 {
		fmt.Println("ok2")
	}

	// || 用法
	// true || true = true
	if age > 30 || age < 50 {
		fmt.Println("ok3")
	}

	// true || false = true
	if age > 30 || age < 40 {
		fmt.Println("ok4")
	}

	// ! 用法
	if age > 30 {
		fmt.Println("ok5")
	}
	// !true = false
	if !(age > 30) {
		fmt.Println("ok6")
	}

	/**
	细节说明:
		&& 也叫短路与:如果第一个条件为false,则第二个条件不会判断，最终结果为 false
		|| 也叫短路或：如果第一个条件为true,则第二个条件不会判断，最终结果为true
	*/

	var i int = 10
	// && 也叫短路与:如果第一个条件为false,则第二个条件不会判断，最终结果为 false
	//  i<9 为false 第二个条件test()不会判断
	if i < 9 && test() {
		fmt.Println("ok7")
	}

	// || 也叫短路或：如果第一个条件为true,则第二个条件不会判断，最终结果为true
	// i>9为true 第二个条件test() 不会判断
	if i > 9 || test() {
		fmt.Println("ok8")
	}

}

// 声明一个函数(测试)
func test() bool {
	fmt.Println("test...")
	return true
}
