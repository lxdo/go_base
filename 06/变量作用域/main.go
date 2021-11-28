package main

import "fmt"

/**
变量作用域
	说明：
		1) 函数内部声明/定义的变量叫局部变量，作用域仅限于函数内部 (局部变量作用域和变量首字母大小写无关)
		2) 函数外部声明/定义的变量叫全局变量，作用域在整个包都有效，
			如果其首字母为大写，则作用域在整个程序有效
		3) 如果变量是在一个代码块，比如for/if中，那么这个变量的作用域就在该代码块
		4) 当全局变量和局部变量名相同,并且都在作用域内，会采用就近原则

*/

//  函数外部声明/定义的变量叫全局变量
// 变量首字母小写在整个包可用
// 变量首字母大写在所有包都可用
var age int = 30
var Name string = "jack"

// 函数
func test() (int, string) {
	age := 20     // 局部变量 作用域只在函数内部
	Name := "tom" // 局部变量 作用域只在函 数内部  (和变量首字母大小写无关   )
	return age, Name
}
func main() {
	fmt.Println("age=", age)
	fmt.Println("Name=", Name)

	//3) 如果变量是在一个代码块，比如for/if中，那么这个变量的作用域就在该代码块
	for i := 1; i <= 3; i++ {
		fmt.Println("i=", i)
	}

	// 再次在for循环定义也不会报重复定义
	for i := 1; i <= 3; i++ {
		fmt.Println("i=", i)
	}

	// fmt.Println("i=", i) 报错 未定义i 在for循环中定义的变量只在for循环中有效
}


