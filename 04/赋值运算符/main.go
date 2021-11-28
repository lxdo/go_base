package main

import "fmt"

/**
介绍：
	赋值运算符就是将某个运算后的值，赋给指定的变量
*/
func main() {
	// 赋值运算符的使用
	var i int
	i = 10 // 基本赋值
	fmt.Println("i=", i)
	// 有两个变量,a和b ，要求将其值进行交换 实现 a=2 b=9
	a := 9
	b := 2
	fmt.Printf("交换前 a= %v ,b= %v \n", a, b) // a=9 b=2
	// 1) 定义一个临时变量
	t := a
	a = b
	b = t
	fmt.Printf("交换后 a= %v ,b= %v \n", a, b) // a=2 b=9
	// 2) 不使用中间变量，直接交换
	a, b = b, a
	fmt.Printf("交换后 a= %v ,b= %v \n", a, b) // a=9 b=2
	// 复合赋值的操作
	a += 17 // 等价 a=a+17
	fmt.Println("a= ", a)

	/**
	细节说明
		1)运算顺序从右向左
		2)赋值运算符的左边只能是变量 右边可以是变量、表达式(任何有值的都可以被看做表达式)、常量值
	*/

	// 赋值运算符的左边只能是变量 右边可以是变量、表达式(任何有值的都可以被看做表达式)、常量值
	var d int
	d = test() + 90 // = 左边是变量  =右边是表达式

	fmt.Println("d = ", d)

	// 面试题:有两个变量，n1和n2,要求将其进行交换，但是不允许使用中间变量
	var n1 int = 10
	var n2 int = 20

	n1 = n1 + n2 //
	n2 = n1 - n2 // n2 = n1+n2-n2 => n2 =n1
	n1 = n1 - n2 // n1 = n1+n2-n1 => n1=n2
	fmt.Printf("n1 = %v n2 = %v ", n1, n2)
}

func test() int {
	return 90
}
