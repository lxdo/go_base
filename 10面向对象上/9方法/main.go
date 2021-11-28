package main

import "fmt"

/**
方法的声明(定义)
	func (recevier type) 方法名 (参数列表) (返回值列表) {
		方法体
		return 返回值
	}

	1) 参数列表:表示方法输入
	2) recevier type:表示这个方法和type这个类型进行绑定,或者说该方法作用于type类型
	3) recevier type: type可以是结构体,也可以其他的自定义类型
	4) recevier:就是type类型的一个变量(实例) 比如Person结构体的一个变量(实例)

	5) 返回值列表:表示返回的值,可以多个
	6) 方法主体:表示为了实现某一功能代码块
	7) return 语句不是必须的

*/
/**
方法注意事项和细节讨论
	1) 结构体类型是值类型,在方法调用中,遵守值类型的传递机制,是值拷贝传递方式
	2) 如程序员希望在方法中，修改方法外结构体变量的值,可以通过结构体指针的方式来处理
	3) go中的方法作用在指定的数据类型上的(即:和指定的数据类型绑定),因此自定义类型,都可以有方法,
       而不仅仅是struct,比如int,float32等都可以有方法
	4) 方法的访问范围控制的规则,和函数一样。方法名首字母小写,只能在本包访问,方法首字母大写,可以在本包和其他包访问
	5) 如果一个类型实现了String()这个方法,那么fmt.Println默认会调用这个类型的String()进行输出
	6) 为了提高效率,通常我们方法和结构体的指针类型(指向结构体地址的指针变量)绑定 func (c *Circle) area() float64 {...}
*/

// 2) 如程序员希望在方法中，修改方法外结构体变量的值,可以通过结构体指针的方式来处理
// 为了提高效率,通常我们方法和结构体的指针类型(指向结构体地址的指针变量)绑定
type Circle struct {
	radius float64
}

// 方法area和Circle结构体指针变量(指向结构体地址的指针变量)绑定
// 此时c是指针变量(指向结构体变量的数据空间)
func (c *Circle) area() float64 {
	// 因为c是指向Circle结构体地址的指针变量 在方法内改变结构体字段值 方法外也同样改变
	c.radius = 10
	// 指向 结构体变量数据空间 的 指针变量c的值 = 结构体变量c的地址
	fmt.Printf("area中 指向结构体的指针变量c 的值= %p \n", c)
	return 3.14 * c.radius * c.radius // 简写方式 c.radius等价于 标准方式 (*c).radius
}
func main() {
	// 创建Circle结构体实例
	var c Circle
	c.radius = 5.0
	res := (&c).area() // 在调用area时 会将结构体变量c的地址(&c)传入area方法中
	// 标准方式 (&c).area() 等价于 c.area()  为了使用方便,编译器底层做了优化,会自动的加上&c
	res2 := c.area() // 等价 (&c).area()
	fmt.Println("res=", res)
	fmt.Println("res2=", res2)
	// 方法内改变结构体变量c的字段radius 方法外也同步改变
	fmt.Println("c.radius=", c.radius)

	// 结构体变量c的地址
	fmt.Printf("main中 结构体变量c 地址= %p \n", &c)
}
