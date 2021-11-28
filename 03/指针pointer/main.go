package main

import "fmt"

/**
指针
	基本介绍
		1) 基本数据类型 ，变量存的就是值，也叫值类型
		2) 获取变量的地址 用& 如：&num
		3) 指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值 比如 var ptr *int = &num
		4) 获取指针类型所指向的值，使用 * ,比如 var ptr *int ，使用 *ptr 获取ptr指向的值

*/

func main() {
	// 基本数据类型在内存布局

	// 基本数据类型 ，变量存的就是值，也叫值类型
	var i int = 10
	// 获取变量的地址 用&   i 的地址是什么 &i
	fmt.Println("i 的地址 ", &i)

	// 指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值 比如 var ptr *int = &num
	var ptr *int = &i
	// 说明 1. ptr是一个指针变量 2.ptr的类型 *int(指向int) 3.ptr本身的值是&i
	fmt.Printf("ptr = %v \n", ptr)
	// ptr指针变量也有其变量地址
	fmt.Printf("ptr 的地址 =%v \n", &ptr)

	// 获取指针类型所指向的值，使用 * ,比如 var ptr *int ，使用 *ptr 获取ptr指向的值
	fmt.Printf("ptr 指向的值=%v \n", *ptr)

	// 练习
	// 1.写一个程序，获取一个int变量num的地址，并显示到终端
	var num int = 9
	fmt.Println("mum 的地址 ", &num)
	// 2.
	var pt *int = &num
	*pt = 10 // 这里修改时，会导致num的值变化
	fmt.Println("num 的值 ", num)

	/**
	指针细节说明
		1.值类型，都有对应的指针类型，形式为 *数据类型 ，比如int的对应的指针就是*int ，float32对应的指针类型就是 *float32,依次类推
		2.值类型包括：基本数据类型 int系列 、float系列、bool、string 、数组、结构体struct
	*/

	/**
	值类型和引用类型
		说明：
			常见的值类型：基本数据类型 int系列 ，float系列，bool，string,数组和结构体struct
			引用类型：指针、slice切片、map、管道chan、interface等都是引用类型

		使用特点：
			值类型：变量直接存储值，内存通常在栈中分配
			引用类型：变量存储的是一个地址，这个地址对应的空间才真正存储数据(值)，内存通常在堆上分配，
			当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾，由GC来回收
	*/

}
