package main

import "fmt"

/**
结构体的注意事项和使用细节:
	1) 结构体的所有字段在内存中是连续的
	2) 结构体是用户单独定义的类型,和其他结构体类型进行转换时需要有完全相同的字段(名字、个数和类型)
	3) 结构体进行type重新定义(相当于取别名),go认为是新的数据类型,但是相互间可以强转
*/

type A struct {
	Num int
}

type B struct {
	Num int
}

// 定义Student结构体
type Student struct {
	Name string
	Age  int
}

// 对结构体Student进行type重新定义(相当于取别名)
type Stu Student

// 基本数据类型进行type重新定义(相当于取别名),go也是认为是新的数据类型,但是相互间可以强转
// 对数据类型int进行type重新定义(取别名)
type integer int

func main() {
	// 2) 结构体是用户单独定义的类型,和其他结构体类型进行转换时需要有完全相同的字段(名字、个数和类型)
	var a A
	var b B
	fmt.Println("a=", a, "b=", b)
	// 结构体之间进行转换,需要有完全相同的字段(名字、个数和类型)
	a = A(b)
	fmt.Println("a=", a, "b=", b)

	// 3) 结构体进行type重新定义(相当于取别名),go认为是新的数据类型,但是相互间可以强转
	var stu1 Student
	var stu2 Stu
	// Stu是Student的重新定义(别名),go认为是新的数据类型,互相间可以强转,但不能直接 stu2 = stu1
	stu2 = Stu(stu1)
	fmt.Println("stu1=", stu1, "stu2=", stu2)

	// 基本数据类型进行type重新定义(相当于取别名),go也是认为是新的数据类型,但是相互间可以强转
	var i integer = 10
	var j int = 20
	// integer是int的重新定义(别名),go认为是新的数据类型,互相间可以强转,但不能直接 j = i
	j = int(i)
	fmt.Println("i=", i, "j=", j)
}
