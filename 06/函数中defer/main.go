package main

import "fmt"

/**
函数中 defer

为什么需要defer：
	在函数中，程序员经常需要创建资源(比如：数据库连接、文件句柄、锁等)，
	为了在函数执行完毕后，及时的释放资源，go的设计者提供defer(延时机制)


细节说明：
	1) 当go执行到一个defer时，不会立即执行defer后的语句，而是将defer后的语句压入到一个栈中(暂称defer栈)，然后继续执行函数下一个语句
	2) 当函数执行完毕后，再从defer栈中，依次从栈顶取出语句执行(先入后出)
	3) 在defer将语句放入到栈时，也会将相关的值拷贝同时入栈

最佳实践
	defer最主要的价值在，当函数执行完毕后，可以及时的释放函数创建的资源
	在go中的通常做法是，创建资源后，比如(打开了文件，获取了数据库连接，或者是锁资源)，可以立即使用defer关闭

*/

func sum(n1, n2 int) int {
	// 当执行到defer时 ，暂时不执行，会将defer后面的语句压入到独立的栈(defer栈)
	// 当函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行

	defer fmt.Println("ok1 n1=", n1) // defer  // 先入后出 第三执行
	defer fmt.Println("ok2 n2=", n2) // defer  // 先入后出 第二执行

	// 在defer将语句放入到栈时，也会将相关的值拷贝同时入栈
	// 即使defer在函数执行完执行 但这里不会对defer中n1、n2的值产生影响
	n1++
	n2++

	s := n1 + n2
	fmt.Println("ok3 s=", s) // 第一执行
	return s
}
func main() {
	a := sum(10, 20)
	fmt.Println("ok4 a=", a) // 第四执行
}
