package main

import "fmt"

/**
	3) go中的方法作用在指定的数据类型上的(即:和指定的数据类型绑定),因此自定义类型,都可以有方法,
       而不仅仅是struct,比如int,float32等都可以有方法
*/
// 自定义类型integer(给int起别名为integer)
type integer int

// 给integer绑定一个方法print
// 这里i是integer类型的变量
func (i integer) print() { // 值传递 方法内改变值方法外不变
}

// 给integer类型变量的指针变量(指向integer类型变量的数据空间)绑定一个方法change
func (i *integer) change() { // 引用传递 方法内改变值方法外也改变
	*i += 1

}
func main() {
	/**
		3) go中的方法作用在指定的数据类型上的(即:和指定的数据类型绑定),因此自定义类型,都可以有方法,
	       而不仅仅是struct,比如int,float32等都可以有方法
	*/
	// 定义integer类型变量
	var i integer = 10
	// 用integer变量调用绑定的方法print
	i.print() // 会把integer变量i当做参数传入print方法(默认值传递)
	// 用integer变量i的指针变量(指向变量i的空间地址)调用change方法(引用传递)
	i.change() // 会把&i(指向变量i的空间地址)当做参数传入change方法(引用传递)  // 简写方式
	// i.change() 等价于 (&i).change()
	(&i).change()
	fmt.Println("i=", i)

	/**

	 */
}
