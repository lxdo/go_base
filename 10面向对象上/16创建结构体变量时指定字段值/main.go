package main

import "fmt"

/**
创建结构体变量时指定字段值

说明:
	go在创建结构体实例(变量)时，可以直接指定字段的值
*/
type Stu struct {
	Name string
	Age  int
}

func main() {
	// 在创建结构体变量时，就直接指定字段的值
	var stu1 = Stu{"tom", 18} // 必须对应结构体中字段的位置顺序
	fmt.Println("stu1=", stu1)

	// 在创建结构体变量时，把字段名和字段值写在一起
	stu2 := Stu{
		Name: "jack",
		Age:  20, // 最后也要有个逗号
	}
	fmt.Println("stu2=", stu2)
	// 这种写法不依赖结构体中字段的位置顺序
	// 在给结构体变量赋值时,可以不按照结构体中字段的位置顺序
	stu3 := Stu{
		Age:  20,
		Name: "jack",
	}
	fmt.Println("stu3=", stu3)

	// 返回结构体变量的指针类型  直接指定字段的值
	var stu4 = &Stu{"zhang", 29} // stu4 是指向结构体变量Stu{"zhang", 29}数据空间的指针变量
	// stu4--->地址---> 结构体变量Stu{"zhang", 29}
	fmt.Println("stu4=", *stu4)

	// 返回结构体变量的指针类型  把字段名和字段值写在一起，这种写法 不依赖结构体中字段的位置顺序
	stu5 := &Stu{
		Age:  29,
		Name: "zhang", // 最后要有逗号
	}
	fmt.Println("stu5=", *stu5) // 获取指针变量指向的数据空间的值 *stu5
}
