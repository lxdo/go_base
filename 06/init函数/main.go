package main

import "fmt"


/**
init函数

基本介绍：
	每一个源文件都可以包含一个init函数，该函数会在main函数执行前，被go运行框架调用，也就是说init会在main函数前调用
	通常可以在init函数中完成初始化工作

细节：
	1) 如果一个文件同时包含全局变量定义，init函数和main函数，则执行的流程是 全局变量定义->init函数->main函数
	2) init函数最主要的作用，就是完成一些初始化的工作
		如果引入的包中也有init函数，则引入包中的init函数最先执行
		执行顺序： 引入的包init函数 -> 本包 全局变量定义-> 本包init函数 ->本包 main函数

面试题：
	如果main.go和utils.go都含有变量定义,init函数时，执行的流程又是怎么样呢

*/
var age = test() // 全局变量定义

func test() int {
	fmt.Println("age")
	return 90
}

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
}
