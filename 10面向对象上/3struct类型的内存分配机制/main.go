package main

import "fmt"

/**
struct 类型的内存分配机制
*/

// 定义一个Person结构体
type Person struct {
	Name string
	Age  int
}

func main() {
	var p1 Person
	p1.Name = "小明"
	p1.Age = 20
	// 结构体变量默认是值拷贝
	var p2 Person = p1
	fmt.Println("p2=", p2) // p2= {小明 20}
	p2.Name = "tom"
	fmt.Println("p2=", p2) // p2= {tom 20}
	fmt.Println("p1=", p1) // p1= {小明 20}

	// 变量总是存在内存中的

	var p3 *Person = &p1 // 把结构体变量p1的数据空间地址 给 指针变量 p3
	// 标准写法 (*p3).Name
	fmt.Println((*p3).Name) // 小明
	// 简易写法 p3.Name 两者等价
	fmt.Println(p3.Name) // 小明
	p3.Name = "jack"     // 等价 (*p3).Name="jack"
	// p3是指向结构体变量p1的数据空间的指针变量 p3的值就是p1的地址 通过指针修改该地址的字段值 也会影响p1字段值
	fmt.Println(p1.Name) // jack
	fmt.Println(p3.Name) // jack p3.Name 等价 (*p3).Name

	// fmt.Println(*p3.Name) 报错  如果用标准写法 *p3 一定要加() => (*p3)  原因: .的运算符优先级比*高

}
