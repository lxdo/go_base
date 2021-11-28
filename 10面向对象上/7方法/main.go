package main

import "fmt"

/**
方法
	基本介绍:
		1) 在某些情况下，我们需要声明(定义)方法。比如Person结构体:除了有一些字段外(年龄、姓名...),
		Person结构体还有一些行为:可以说话、跑步...,这时就要用方法才能完成
		2) go中的方法是作用在指定的数据类型上的(即:和指定的数据类型绑定),因此自定义类型,都可以有方法,而不仅仅是struct

	方法的声明和调用:
		type A struct {
			Num int
		}

		func (a A)test(){
			fmt.Println(a.Num)
		}

		1) func (a A)test(){} 表示A结构体有一个方法,方法名为test
		2) (a A) 体现 test方法是和A类型绑定的


*/

// 声明类型Person和定义结构体Person
type Person struct {
	Name string
}

// 给Person类型绑定一个方法test
// p表示调用test方法的结构体变量的副本(值传递)
func (p Person) test() {
	p.Name = "jack"                     // 在方法内改变结构体变量字段值 不会 影响方法外调用该方法的结构体变量的字段值
	fmt.Println("test p.Name=", p.Name) // test p.Name= jack
}
func main() {
	var p Person   // 结构体实例化
	p.Name = "tom" // 结构体变量(实例) 字段赋值
	p.test()       // 调用结构体Person(类型Person)的方法test
	// 结构体实例p会在调用时传递给test方法(值传递)

	/**
	1)test方法和Person类型绑定
	2)test方法只能通过Person类型的变量(Person结构体的实例)来调用,而不能直接调用,也不能使用其他类型变量来调用

	相当于php类中的方法,只能用类的实例来调用类中方法
		class Person
		{
			public function test()
			{

			}
		}

	3) func (p Person)  中p表示调用test方法的结构体变量的副本(值拷贝)
	   在方法内改变结构体变量字段值不会影响方法外调用该方法的结构体变量字段值
	4) func (p Person) 中 p这个形参名,名字是自定义的
	*/

	fmt.Println("main p.Name=", p.Name) // main p.Name= tom

}
