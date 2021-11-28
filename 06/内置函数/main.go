package main

import "fmt"

/**
go设计者为了编程方便，提供了一些函数，这些函数可以直接使用，我们称为go的内置函数

https://studygolang.com/pkgdoc -> builtin部分
*/
func main() {
	// 1) len:用来求长度

	// 2) new:用来分配内存，主要用来分配值类型，比如int、float32、struct...返回的是指针
	/**
	func new(Type) *Type
	内建函数new分配内存，其第一个实参为类型，而非值。其返回值为指向该类型的新分配的零值的指针
	*/
	// new的使用
	num1 := 100
	fmt.Printf("num1的类型为 %T num1的值为 %v num1的地址为 %v \n", num1, num1, &num1)

	num2 := new(int) // 创建一个*int(int型指针变量 指针指向的空间初始值为0)
	fmt.Printf("num2的类型为 %T num2的值为 %v num2的地址为 %v num2这个指针指向的值为 %v \n", num2, num2, &num2, *num2)
	*num2 = 100 // 修改num2指向的值
	fmt.Printf("num2的类型为 %T num2的值为 %v num2的地址为 %v num2这个指针指向的值为 %v \n", num2, num2, &num2, *num2)

	// 3) make:用来分配内存，主要用来分配引用类型，比如chan,map,slice
}
