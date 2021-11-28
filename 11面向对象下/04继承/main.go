package main

import "fmt"

/**
继承(下)
	深入讨论:
			1) 结构体可以使用嵌套匿名结构体所有的字段和方法。即:首字母大写或者小写的字段、方法，都可以使用
			2) 匿名结构体字段访问可以简化
				结构体变量.匿名结构体.匿名结构体字段=>结构体变量.匿名结构体字段
				结构体变量.匿名结构体.匿名结构体方法=>结构体变量.匿名结构体方法
				b.A.Name = "tom" => b.Name="tom"
				b.A.Say()		 => b.Say()

				执行流程
				1) 当我们直接通过b访问字段或方法时，其执行流程如下 比如b.Name ，
				2) 编译器会先看b对应结构体有没有Name字段，如果有，则直接调用结构体B的Name字段，
				3) 如果没有就去看B中嵌入的匿名结构体A中有没有Name字段,如果有，就调用，如果没有,继续查找A中嵌入的匿名结构体有没有Name字段
					如果都找不到就报错

			3) 当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则访问，
				如希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分	=> 结构体变量.匿名结构体.匿名结构体字段(方法)
*/

// 结构体A
type A struct {
	Name string
	age  int
}

func (a *A) Say() {
	fmt.Println("A say", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
}

// 继承A  也有一个Name字段
type C struct {
	A
	Name string
}

func (c *C) Say() {
	fmt.Println("C say", c.Name)
}

func main() {
	// 1) 结构体可以使用嵌套匿名结构体所有的字段和方法。即:首字母大写或者小写的字段、方法，都可以使用
	var b B
	b.A.Name = "tom"
	b.A.age = 19
	b.A.Say()   // A say tom
	b.A.hello() // A hello tom

	// 2) 匿名结构体字段访问可以简化
	b.Name = "smith" // 结构体B没有Name字段 就近访问原则 访问其嵌套的匿名结构体A的Name字段 即给A的Name字段赋值
	b.age = 20       // 结构体B没有age字段 就近访问原则 访问其嵌套的匿名结构体A的age字段 即给A的age字段赋值
	b.Say()          // A say smith
	b.hello()        // A hello smith

	// 3) 当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则访问，
	var c C
	c.Name = "jack" // C结构体和A结构体都有Name 字段 就近访问原则  优先访问结构体C中Name
	c.age = 100     //   C结构体中没有age字段 就会访问其嵌套的匿名结构体C的age字段
	c.Say()         //  C say jack  C结构体和A结构体都有Say方法，就近访问原则 优先访问结构体C中Say  此时c.Name正好有值jack
	c.hello()       // A hello  C结构体没有hello方法 ，就会访问其嵌套的匿名结构体C的hello方法 此时a.Name 没有值

	// 当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则访问，如果就要访问匿名结构体中字段或者方法
	// 结构体变量.嵌套匿名结构体.字段/方法
	b.A.Name = "sock"
	b.A.Say() // A say sock
}
