package main

import "fmt"

/**
5) 如果一个类型实现了String()这个方法,那么fmt.Println默认会调用这个类型的String()进行输出
*/

// 定义结构体Student
type Student struct {
	Name string
	Age  int
}

// 给结构体Student指针变量(指向结构体Student变量的空间)实现(绑定)方法String()
func (stu *Student) String() string {
	// fmt.Sprintf 格式化字符串
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)

	return str
}

// 定义结构体Person
type Person struct {
	Name string
	Age  int
}

// 给结构体实现(绑定)String方法
func (p Person) String() string {
	// fmt.Sprintf 格式化字符串
	str := fmt.Sprintf("Name=[%v] Age=[%v]", p.Name, p.Age)

	return str
}
func main() {
	// 定义Student结构体变量stu
	stu := Student{"tom", 20}
	// 调用结构体指针变量绑定的String方法
	fmt.Println((&stu).String()) // 标准写法   结构体指针变量调用结构体中方法
	fmt.Println(stu.String())    // 简便写法   结构体指针变量(&省略)调用结构体中方法
	fmt.Println(&stu)            // 超简写法 fmt.Println输出时默认会调用 结构体指针变量绑定的String方法返回值 进行输出
	// 超简写法 如果是指针变量的话 还是要加上& 不然无法区别是stu还是&stu


	p := Person{"jack", 22}
	// 调用结构体绑定的String方法
	fmt.Println(p.String()) // 标准写法  结构体变量调用结构体方法
	fmt.Println(p)          // 超简写法 fmt.Println输出时默认会调用 结构体 绑定的String方法返回值 进行输出

}
