package main

import (
	"errors"
	"fmt"
)

/**
错误处理-自定义错误
	go中，也支持自定义错误，使用errors.New和panic内置函数
	1) errors.New("错误说明"),会返回一个error类型的值，表示一个错误
	2) panic 内置函数，接收一个interface{}类型的值(也就是任何值)作为参数，可以接收error类型的变量，输出错误信息，并退出程序

*/

// 案例说明
// 函数去读取以配置文件init.conf的信息 如果文件名传入不正确，就返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "init.conf" {
		// 读取配置

		return nil // 程序正常 错误信息为nil
	} else {
		return errors.New("读取文件错误...") // 程序异常 返回自定义错误
	}
}

func test() {
	err := readConf("init.conf2") // 读取错误
	if err != nil {               // 有自定义错误
		// 如果读取文件发生错误，就输出这个错误，并终止程序
		panic(err) // 把错误以panic输出、终止程序
	}

	fmt.Println("test继续执行")
}
func main() {
	test()
	fmt.Println("main继续执行")
}
