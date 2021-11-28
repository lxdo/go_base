package main

import "fmt"

/**
错误处理
*/
func test() {
	// go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理
	// 使用defer+recover来捕获和处理异常
	defer func() {
		err := recover() // recover为内置函数，可以捕获到异常
		if err != nil {  // 说明捕获到错误
			//	if err := recover(); err != nil { 	// 上面两行可以简写成这个
			fmt.Println("err=", err)
			// 这里就可以将错误信息发送给管理员
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2 // 除数分母不能为0
	fmt.Println("res=", res)
}

/**
1) 默认情况下，当发生错误后(panic)，程序就会退出(崩溃)
2) 如果我们希望：当发生错误后，可以捕获到错误，并进行处理，保证程序可以继续执行。还可以在捕获错误后，给管理员发生通知(邮件、短信等)
*/

/**
基本说明
	1) go语言追求简洁优雅，所以go不支持传统的try...catch...finally 这种处理
	2) go中引入的处理方式为: defer , panic, recover
	3) 这几个异常的使用场景：go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理

错误处理的好处
	进行错误处理后，程序不会轻易挂掉，如果加入预警代码，就可以让程序更加的健壮
*/
func main() {
	test() // 默认情况下，当发生错误后(panic)，程序就会退出(崩溃) 不会向下执行
	fmt.Println("test执行后的代码")
}
