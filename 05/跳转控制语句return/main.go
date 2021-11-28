package main

import "fmt"

/**
跳转控制语句 return
介绍：
	return 使用在方法或者函数中，表示跳出所在的方法或函数

说明：
	1) 如果return 是在普通的函数，则表示跳出该函数，即不再执行函数中return后面代码，也可以理解成终止函数
	2) 如果return 是在main函数 ，表示终止main函数 ，也就是说终止程序
*/
func main() {
	for i := 1; i <= 10; i++ {
		if i == 3 {
			return // 直接结束所在方法 后面所有代码不再执行
		}
		fmt.Println("i =", i)
	}
	fmt.Println("hello world")
}
