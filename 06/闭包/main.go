package main

import (
	"fmt"
	"strings"
)

/**
闭包 【难点】
	介绍：
		闭包就是一个函数和与其相关的引用环境组合的一个整体(实体)

		一般来说闭包是一个匿名函数和函数引用到的外部变量组成
		外部变量只会初始化赋值一次，后面会在匿名函数调用中不断改变
*/

// 案例：累加器
func addUpper() func(int) int {
	// 下面变量n、str和匿名函数就组成了一个闭包

	//一般来说闭包是一个匿名函数和函数引用到的外部变量组成
	var n int = 10    // 闭包中变量初始化赋值只执行一次
	var str = "hello" // 闭包中变量初始化赋值只执行一次
	fmt.Println("初始化 str=", str)
	return func(x int) int {

		n += x
		str += "1"
		fmt.Println("匿名函数 str=", str) // 变量值在调用中不断改变
		return n
	}

}

/**

说明：
		1)addUpper 是一个函数，返回的数据类型 是func (int) int
		2)addUpper返回的是一个匿名函数，但是这个匿名函数引用到函数外的n,因此这个匿名函数就和n形成一个整体，构成闭包
		3)可以理解为：闭包是类，匿名函数是操作(方法)，n是属性字段  匿名函数和它使用到的n构成闭包
		4)当我们反复调用下面的f函数时，因为n 是初始化一次，因此每调用一次就进行累计
		5)我们要搞清楚闭包的关键，就是要分析出返回的函数它使用(引用)到哪些变量。因为函数和它引用到的变量共同构成闭包

*/

/**
闭包的最佳实践:
	请编写一个程序，具体要求如下
		1) 编写一个函数 makeSuffix (suffix string) 可以接收一个文件后缀名(比如.jpg)，并返回一个闭包
		2) 调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(如果.jpg),则返回文件名.jpg，如果已经有.jpg后缀,则返回原文件名
		3) 要求用闭包的方式完成
		4) strings.HasSuffix // 该函数可以判断某个字符串是否有指定的后缀
*/

// 闭包实现
func makeSuffix(suffix string) func(string) string {
	// return 返回的就是一个闭包 ：匿名函数使用到外部变量suffix
	return func(name string) string {
		// 如果name 没有指定后缀，则加上，否则返回原来的名字
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}

		return name
	}
}

// 传统方法使用
func makeSuffix2(suffix string, name string) string {
	// 如果name 没有指定后缀，则加上，否则返回原来的名字
	if !strings.HasSuffix(name, suffix) {
		return name + suffix
	}

	return name
}

/**
说明：
	返回的匿名函数和makeSuffix(suffix string)函数的suffix变量组合成一个闭包，因为返回的函数引用到suffix这个变量
	如果使用传统的方法，也可以轻松实现这个功能，但是传统方法需要每次都传入后缀名，比如.jpg。而闭包因为可以保留上次引用的某个值，
	所以我们传入一次就可以反复使用
*/
func main() {
	// 调用累加器案例
	// 返回值是一个函数 返回 func(x int) int { n+=x return n } 函数
	f := addUpper()
	// 调用累加器函数返回的函数 即调用func(x int) int { n+=x return n } 函数
	fmt.Println("a=", f(1)) // 11
	fmt.Println("a=", f(2)) // 13
	fmt.Println("a=", f(3)) // 16

	// 闭包实现 使用makeSuffix
	// 返回一个闭包
	ms := makeSuffix(".jpg")
	fmt.Println("文件名处理后=", ms("winter"))
	fmt.Println("文件名处理后=", ms("bind.jpg"))
	// 传统方式实现 使用makeSuffix2
	// 传统方法需要每次都传入后缀名
	fmt.Println("文件名处理后=", makeSuffix2(".jpg", "winter"))
	fmt.Println("文件名处理后=", makeSuffix2(".jpg", "bind.jpg"))
}
