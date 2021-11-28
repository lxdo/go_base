package main

import "fmt"

/**
基本数据类型的转换

介绍：
	go和Java/c不同 ，go在不同类型的变量之间赋值时需要显式转换。也就是说go中数据类型不能自动转换
基本语法
	表达式T(v)将值v转换为类型T
	T：就是数据类型，比如int32,int64,float32等等
	v:就是需要转换的变量
*/

func main() {
	var i int32 = 100
	// i => float
	var f float32 = float32(i)
	var i2 int8 = int8(i)
	var i3 int64 = int64(i) // 低精度->高精度
	fmt.Printf("i=%v type=%T f=%v type=%T i2=%v type=%T i3=%v type=%T \n", i,i, f,f, i2,i2, i3,i3)
	/**
	细节说明
	1.go中，数据类型的转换可以是从表示范围小->表示范围大 ，也可以范围大->范围小
	*/
	// 2.被转换的是变量存储的数据（即值），变量本身的数据类型并没有变化
	fmt.Printf("i type is %T \n", i)

	// 3.在转换中，比如将int64转成int8（-128-127），编译时不会报错， 只是转换的结果是按溢出处理，和我们希望的结果不一样
	// 因此在转换时，需要考虑范围
	var n1 int64 = 999999
	var n2 int8 = int8(n1)
	fmt.Println("n2= ", n2) // 不会报错 会按溢出处理

	// 练习 基本数据类型的转换 判断是否能够通过编译
	// 第一题
	var a1 int32 = 12
	var a2 int64
	var a3 int8
	// a2 = a1 + 20   // 报错 a2 为int64 a1为int32 无法直接运算
	// a3 = a1 + 20   // 报错 a3 为int8  a1为int32 无法直接运算
	// 正确
	a2 = int64(a1) + 20
	a3 = int8(a1) + 20
	fmt.Println(a1, a2, a3)
	// 第二题
	var b1 int32 = 12
	var b3 int8
	var b4 int8
	b4 = int8(b1) + 127 // b4 为int8 b1为int32 需要类型转换  在求和时 12+127 结果在int8(-128-127)范围外 编译通过 结果按溢出处理
	// b3 = int8(b1) + 128   // 报错 b3为int8 b1为int32 需要类型转换 但 求和中128 已经超出 int8 (-128-127)表数范围 编译报错
	// 正确
	b5 := b1 + 128
	fmt.Println(b1, b3, b4, b5)

}
