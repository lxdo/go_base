package main

import "fmt"

/**
面向对象编程-继承

继承的深入讨论
	4) 结构体嵌入两个(或多个)匿名结构体，如两个匿名结构体有相同的字段和方法(同时结构体本身没有同名的字段和方法)
	在访问时，就必须明确指定匿名结构体名字，否则编译报错

	5) 如果一个struct嵌套了一个有名结构体，这种模式就是组合，如果是组合关系，那么在访问组合的结构体的字段或方法时，必须带上结构体的名字

	6) 嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值
*/

// 结构体A
type A struct {
	Name string
	age  int
}

// 结构体B
type B struct {
	Name  string // 和结构体A有相同的字段
	score float64
}

// 结构体C继承A和B (嵌入A和B)
type C struct {
	A
	B
}

// 组合
// 5) 如果一个struct嵌套了一个有名结构体，这种模式就是组合，如果是组合关系，那么在访问组合的结构体的字段或方法时，必须带上结构体的名字
type D struct {
	a A // 有名结构体
}

// 结构体 商品
type Goods struct {
	Name  string
	Price float64
}

// 结构体 品牌
type Brand struct {
	Name    string
	Address string
}

// 结构体 Tv
type Tv struct {
	Goods // 嵌套匿名结构体
	Brand
}

type Tv2 struct {
	*Goods // 嵌套指向匿名结构体数据空间的指针变量
	*Brand
}

func main() {
	// 4) 结构体嵌入两个(或多个)匿名结构体，如两个匿名结构体有相同的字段和方法(同时结构体本身没有同名的字段和方法)
	//	在访问时，就必须明确指定匿名结构体名字，否则编译报错
	var c C
	// c.Name="tom" // 报错 未明确指定匿名结构体名字 导致Name 无法确定
	// 如果在结构体C中也有Name string字段  ，则不会报错，因为按照就近访问原则 已经确定了Name 字段(即结构体C中的Name)
	c.A.Name = "tom"

	// 5) 如果一个struct嵌套了一个有名结构体，这种模式就是组合，如果是组合关系，那么在访问组合的结构体的字段或方法时，必须带上结构体的名字
	var d D
	// d.Name // 直接访问有名结构体的字段或属性 报错
	d.a.Name = "jack" // 结构体变量.有名结构体的名字.有名结构体属性(方法)
	// 如果结构体D中有Name string 字段 则 d.Name 不会报错 访问的是结构体D中的Name

	// 总结:结构体D嵌套一个有名结构体A,在访问有名结构体的字段Name时，就必须带上有名结构体的名字 如 d.a.Name

	// 6) 嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值
	tv := Tv{
		Goods{"电视机001", 5000.99},
		Brand{"海尔", "山东"},
	}

	tv2 := Tv{
		Goods{Name: "电视机002", Price: 3000.99},
		Brand{Name: "夏普", Address: "北京"}, // 字段名:字段值 不依赖于字段顺序
	}
	fmt.Println(tv)
	fmt.Println(tv2)

	tv3 := Tv2{
		&Goods{"电视机003", 7000.99}, // 嵌套指向匿名结构体Goods的指针类型
		&Brand{"创维", "合肥"},
	}

	tv4 := Tv2{
		&Goods{Name: "电视机004", Price: 3001.99}, // 嵌套指向匿名结构体Goods的指针类型
		&Brand{Name: "美的", Address: "安徽"},      // 字段名:字段值 不依赖于字段顺序
	}
	fmt.Println(tv3)
	fmt.Println(tv4)
	fmt.Println(*tv3.Goods,*tv3.Brand)
	fmt.Println(*tv4.Goods,*tv4.Brand)
}
