package main

import (
	"fmt"
)

/**
方法快速入门

*/
// 定义Person结构体
type Person struct {
	Name string
}

// 给结构体Person绑定speak方法
func (p Person) speak() {
	fmt.Println(p.Name, " good")
}

// 1)在方法中可以进行各种运算
// 给结构体Person添加jisuan方法
func (p Person) jisuan() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println(p.Name, " sum=", sum)
}

// 2)方法可以接收外面传参
// 给结构体Person添加jisuan2方法
//                     参数类型
func (p Person) jisuan2(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println(p.Name, " sum=", sum)
}

// 3)方法有返回值
// 给结构体Person添加getSum方法
//                     参数类型      返回值类型
func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2 // 方法返回结果
}

// 声明一个结构体Circle ,字段为radius
type Circle struct {
	radius float64
}

// 声明一个方法area和Circle绑定,可以返回面积
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}
func main() {
	var p Person   // 定义结构体变量
	p.Name = "tom" // 结构体变量字段赋值
	p.speak()      // 结构体变量p调用speak方法 结构体变量p传入speak方法(值传递 字段方法内改变不影响方法外字段)
	p.jisuan()     // 结构体变量p调用jisuan方法
	p.jisuan2(10)  // 结构体变量p调用jisuan2方法 并传入参数

	n1 := 10
	n2 := 20
	sum := p.getSum(n1, n2) // 结构体变量p调用getSum方法,并传入参数,接收返回值
	fmt.Println("sum=", sum)

	/**
	方法的调用和传参机制原理
		说明:方法的调用和传参机制 和函数基本一样,不一样的地方是方法调用时,会将调用方法的变量(实例),当做实参也传递给方法
	         如果参数是值类型,则进行值拷贝,如果参数是引用类型,则进行地址拷贝

	*/

	// 案例
	// 1) 声明一个结构体Circle ,字段为radius
	// 2) 声明一个方法area和Circle绑定,可以返回面积
	// 创建一个Circle结构体变量c
	var c Circle
	c.radius = 4.0  // 给c的字段赋值
	res := c.area() // 调用Circle结构体绑定的方法  结构体变量c作为参数传入area(值传递)
	fmt.Println("res=", res)
	
}
