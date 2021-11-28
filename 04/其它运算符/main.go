package main

import "fmt"

/**
其它运算符
说明：
	&  含义：返回变量存储地址 例：&a,将给出变量a的实际地址
	*  含义：指针变量        例：*a,是一个指针变量，获取a指向的值

	go不支持三元运算符 用 if  else 实现
*/
func main() {
	// &
	a := 100
	fmt.Println("a的地址 = ", &a)

	// *
	var ptr *int = &a

	fmt.Println("ptr的值 = ", ptr)
	fmt.Println("ptr指向的值 = ", *ptr)

	// 练习 1.求出两个数的最大值
	var n1 int = 10
	var n2 int = 40
	var max int
	// go不支持三元运算符 用 if  else 实现
	if n1 > n2 {
		max = n1
	} else {
		max = n2
	}
	fmt.Println("两个数最大值为 ", max)

	// 练习 2.求三个数最大值  思路：先求出两个数最大值。再和第三个数比较
	var n3 = 45

	if n3 > max {
		max = n3
	}

	fmt.Println("三个数最大值为 ", max)

}
