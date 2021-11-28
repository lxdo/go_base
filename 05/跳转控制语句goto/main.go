package main

import "fmt"

/**
跳转控制语句 goto

介绍：
	1) go语言的goto语句可以无条件地转移到程序中指定的行
	2) goto语句通常与条件语句(if)配合使用。可用来实现条件转移，跳出循环体等功能
	3) 在go程序设计中一般不主张使用goto语句，以免造成程序流程的混乱，使理解和调试程序都产生困难

基本语法：
	goto label // 标签自定义
	...        // 这些代码不再执行
	label:statement  // 从goto标签对应的这一行代码执行

*/

func main() {
	// goto案例
	fmt.Println("ok1")
	goto label1   // 代码执行到这一行 直接跳到标签对应行执行代码
	fmt.Println("ok2") // 不执行
	fmt.Println("ok3") // 不执行
	fmt.Println("ok4") // 不执行
label1:
	fmt.Println("ok5") // 从这里继续开始执行代码
	fmt.Println("ok6")

}
