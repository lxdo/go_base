package main

import "fmt"

/**
结构体
	如何声明结构体:
		type 结构体名称 struct {  // 如果结构体名称首字母大写 则该结构体可以在其他包使用 小写只能在本包使用
			field1 type // 字段名称 字段数据类型  如果字段名称首字母大写 则该字段可以在其他包使用 小写只能在本包使用
			field2 type
		}

	字段/属性:
		基本介绍:
			1) 从概念或叫法上看:结构体字段=属性=field
			2) 字段是结构体的一个组成部分,一般是基本数据类型、数组,也可以是引用类型
		注意事项和细节说明:
			1) 字段声明语法同变量, 格式: 字段名 字段数据类型
			2) 字段的类型可以为:基本类型、数组或引用类型
			3) 在创建一个结构体变量(实例)后,如果没有给字段赋值,都对应一个零值(默认值)
               默认值的规则:
					(1)布尔类型是false 整型float型是0 字符串是""
		       		(2)数组类型的默认值和它的元素数据类型相关,比如score [3]int 则默认值为[0,0,0]
					(3)slice和map的零值都是nil,即还没有分配空间,需要先make(分配数据空间)后使用
			4) 不同结构体变量的字段是独立,互不影响,一个结构体变量字段的更改，不影响另外一个





*/

// 定义一个cat结构体,将cat的各个字段/属性信息,放入到cat结构体进行管理
type Cat struct {
	Name  string // 属性 可以在结构体外用 结构体实例.属性名 访问
	Age   int
	Color string
}

// 3) 在创建一个结构体变量(实例)后,如果没有给字段赋值,都对应一个零值(默认值)
// (3)slice和map的零值都是nil,即还没有分配空间,需要先make(分配数据空间)后使用
type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int              // 指针
	slice  []int             // 切片
	map1   map[string]string // map
}

func main() {
	// 创建一个结构体Cat的实例
	var cat1 Cat
	fmt.Println("cat1=", cat1) // 结构体实例化 就已经分配了数据空间 值为默认值
	// 给结构体实例中属性赋值
	cat1.Name = "tom" // 属性 可用 结构体实例.属性名 访问
	cat1.Age = 3
	cat1.Color = "white"
	fmt.Println("cat1=", cat1)

	/**
	结构体和结构体变量(实例)的区别和联系:
		1) 结构体是自定义的数据类型,代表一类事物   如上面的struct Cat
		2) 结构体变量(实例)是具体的,实际的,代表一个具体变量  如上面的struct 实例cat1
	*/

	// 3) 在创建一个结构体变量(实例)后,如果没有给字段赋值,都对应一个零值(默认值)
	// (3)slice和map的零值都是nil,即还没有分配空间,需要先make后使用
	// 定义结构体变量(实例)
	var p1 Person
	fmt.Println("p1=", p1)
	// slice、map 默认值是nil 不可使用
	if p1.slice == nil && p1.map1 == nil {
		fmt.Println("slice、map 默认值是nil 不可使用")
	}

	// p1.slice[0] = 100 // 报错 此时 slice为nil 没有数据空间 不能使用 需要先make
	// 使用结构体变量中的slice、map需要先make 才能使用
	p1.slice = make([]int, 10)
	p1.slice[0] = 100
	fmt.Println("p1=", p1)
	// p1= { 0 [0 0 0 0 0] <nil> [100 0 0 0 0 0 0 0 0 0] map[]}  slice 已经有数据空间了 默认值和slice的元素数据类型相关

	// 结构体实例化后, map默认nil 需要先make
	p1.map1 = make(map[string]string)
	p1.map1["key"] = "value"
	fmt.Println("p1=", p1)

	// 4) 不同结构体变量的字段是独立,互不影响,一个结构体变量字段的更改，不影响另外一个。结构体是值类型,默认为值拷贝
	cat2 := cat1
	fmt.Println("cat1=", cat1)
	fmt.Println("cat2=", cat2)
	cat2.Name = "jack"
	fmt.Println("cat1=", cat1)
	fmt.Println("cat2=", cat2)
	// 结构体变量cat2的字段Name更改 不会影响结构体变量cat1的字段Name

	// 创建结构体变量的4种方式
	// 方式1 直接声明 ：var 结构体变量名 结构体名
	var cat3 Cat
	fmt.Println("cat3=", cat3)

	// 方式2 声明并赋值 ： var 结构体变量名 结构体名 = 结构体名{字段值,字段值,...}
	var cat4 Cat = Cat{"mary", 5, "yellow"} // 声明时同时赋值
	fmt.Println("cat4=", cat4)
	// 也可以只声明不赋值 ：var 结构体变量名 结构体名 = 结构体名{}
	var cat5 Cat = Cat{}
	fmt.Println("cat5=", cat5)
	cat5.Name = "jili"

	// 方式3 用new结构体(new(Cat)) 来创建指向结构体的指针变量
	var cat6 *Cat = new(Cat)
	// cat6是一个指针,要给结构体中字段赋值 标准方式用：(*cat6).字段名 简化方式用：cat6.字段名
	(*cat6).Name = "smith"
	// 上面写法等价于 cat6.Name = "smith"
	// 原因:go设计者为了程序员使用方便,简化书写,在底层会对 cat6.Name = "smith" 进行处理,
	//	会给cat6加上取值运算(*cat6) 转化成 (*cat6).Name = "smith"
	cat6.Name = "smith"
	(*cat6).Age = 20
	// 上面写法等价于
	cat6.Age = 20
	(*cat6).Color = "black"
	// 上面写法等价于
	cat6.Color = "black"
	fmt.Printf("cat6 的值(指向的空间地址)= %p \n", cat6)
	fmt.Println("cat6 指向的空间地址的值", *cat6)

	// 方式4 用 &结构体实例(&Cat{}) 来创建指向结构体的指针变量
	var cat7 *Cat = &Cat{}
	// cat7是一个指针,要给结构体中字段赋值 标准方式用：(*cat7).字段名 简化方式用：cat7.字段名
	(*cat7).Name = "scott" // cat7.Name = "scott"
	// 上面写法等价于
	// 原因:go设计者为了程序员使用方便,简化书写,在底层会对 cat7.Name = "scott" 进行处理,
	//	会给cat7加上取值运算(*cat7) 转化成 (*cat7).Name = "scott"
	cat7.Name = "scott"
	cat7.Age = 88
	fmt.Printf("cat7 的值(指向的空间地址)= %p \n", cat7)
	fmt.Println("cat7 指向的空间地址的值", *cat7)
	// 用 &结构体实例(&Cat{}) 来创建指向结构体的指针变量 时 也可以直接对结构体中字段赋值
	var cat8 *Cat = &Cat{"lucy", 29, "red"}
	fmt.Printf("cat8 的值(指向的空间地址)= %p \n", cat8)
	fmt.Println("cat8 指向的空间地址的值", *cat8)

	/**
	创建结构体变量的4种方式说明:
		(1) 方式3和方式4返回的是结构体指针
		(2) 结构体指针访问字段的标准方式是：(*结构体指针).字段名  比如:(*cat7).Name = "scott"
		(3) 但go做了简化,也支持 结构体指针.字段名 比如:cat7.Name = "scott" 更加符合程序员使用的习惯
			go编译器底层对cat7.Name转化成(*cat7).Name
	*/
}
