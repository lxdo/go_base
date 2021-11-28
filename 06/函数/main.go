package main

import "fmt"

/**
函数
	概念:
		为完成某一功能的程序指令(语句)的集合，称为函数
		在go中，函数分为:自定义函数、系统函数
	基本语法:

			func 函数名 (形参列表) (返回值类型列表) {
				执行语句...
				return 返回值列表
			}

	1) 形参列表：表示函数的输入
	2) 函数中的语句：表示为了实现某一功能代码块
	3) 函数可以有返回值，也可以没有

return 语句
	基本语法:
		go函数支持返回多个返回值
	细节:
		1) 如果返回多个值时，在接收时，希望忽略某个返回值，则使用 _ 符合表示占位忽略
		2) 如果返回值只有一个，(返回值类型列表) 可以不写 () 括号


*/

/**
函数注意事项和细节讨论
	1) 函数的形参列表可以是多个，返回值列表也可以是多个
	2) 形参列表和返回值列表的数据类型可以是值类型和引用类型
	3) 函数的命名遵循标识符命名规范，首字母不能是数字，首字母大写该函数可以被本包文件和其他包文件使用，类似public,
		首字母小写，只能被本包文件使用，其他包文件不能使用 ，类似private
	4) 函数中的变量是局部的，函数外不生效
	5) 基本数据类型和数组默认都是值传递，即进行值拷贝，在函数内修改，不会影响到原来的值
	6) 如果希望函数内的变量能修改函数外的变量，可以传入变量的地址&，函数内以指针的方式操作变量，从效果上类似引用传递
	7) go函数不支持函数 重载 // 重载即为有些语言，函数名相同，参数列表 数据类型或数量不同 被认定为不同函数 go不支持同一个包下函数名相同
	8) 在go中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了，通过该变量可以对函数调用
	9) 函数既然是一种数据类型，因此在go中，函数可以作为形参，并且调用
	10) 为了简化数据类型定义，go支持自定义数据类型
		基本语法:type 自定义数据类型名 数据类型 // 相当于一个别名
	11) 支持对函数返回值命名
	12) 使用 _ 标识符 ，忽略返回值
	13) go支持可变参数

		支持0到多个参数
		func sum (args...int) sum int {
		}

		支持1到多个参数
		func sum (n1 int,args...int) sum int {

		}

		说明：(1) args 是slice(切片)，通过 args[index],可以访问到各个值
             (2) 如果函数的形参列表有可变参数，则可变参数需要放在形参列表最后

*/

// go支持可变参数 args是自定义名称 通常用args
// 案例演示:编写一个函数Sum，可以求出 ，1到多个int的和
// 支持1到多个参数
func Sum1(n1 int, args ...int) int {
	sum := n1
	// 遍历args
	for i := 0; i < len(args); i++ {
		sum += args[i] // args[0] 表示取出args切片的第一个元素值，其他依次类推
	}

	return sum
}

// 函数中的变量是局部的，函数外不生效
func ab(a int) int {
	// num 是ab函数的局部变量，只能在ab函数中使用
	num := 10
	a += num // 基本数据类型和数组默认都是值传递，即进行值拷贝，在函数内修改，不会影响到原来的值
	fmt.Println("ab函数中 a=", a)
	return a
}

// 如果希望函数内的变量能修改函数外的变量，可以传入变量的地址&，函数内以指针的方式操作变量，从效果上类似引用传递
// n 就是 *int类型
func test(n *int) bool {
	*n = *n + 10
	fmt.Println("test函数 n=", *n)
	return true
}

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

// 函数既然是一种数据类型，因此在go中，函数可以作为形参，并且调用
func myFun(fun func(int, int) int, n1 int, n2 int) int {
	return fun(n1, n2)
}

/**
为了简化数据类型定义，go支持自定义数据类型
基本语法: type 自定义数据类型名 数据类型 // 相当于一个别名
*/
// 给函数类型取别名
type myFuncType func(int, int) int // 这时 myFuncType 就是 func(int ,int) int 类型

// 给函数类型取别名 在函数作为形参时，可以简化编写
// 上面函数可以简化为
func my(fun myFuncType, n1 int, n2 int) int {
	return fun(n1, n2)
}

func main() {
	// fmt.Println("ab函数 num=", num) // 这里不能使用num num是ab函数的局部变量
	a := 3
	ab(a) // ab函数内修改a的值 不会影响main函数中a的值
	fmt.Println("main函数中 a=", a)

	// 如果希望函数内的变量能修改函数外的变量，可以传入变量的地址&，函数内以指针的方式操作变量，从效果上类似引用传递
	n := 20
	test(&n)
	fmt.Println("main函数 n=", n)

	// 在go中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了，通过该变量可以对函数调用
	c := cal
	fmt.Printf("c的类型是 %T cal的类型是 %T \n", c, cal)
	c(10, 40, '*') // 等价 cal(10,40,'*')

	// 函数既然是一种数据类型，因此在go中，函数可以作为形参，并且调用
	he := myFun(getSum, 10, 20) // 函数作为参数
	fmt.Println("he=", he)

	/**
	为了简化数据类型定义，go支持自定义数据类型
	基本语法: type 自定义数据类型名 数据类型 // 相当于一个别名
	*/

	// 给int取了别名 在go中 myInt和int虽然都是int类型，但是go认为myInt和int是两个不同类型
	type myInt int
	var numx myInt
	numx = 40
	var numy int
	numy = int(numx) // 如果直接赋值会报错(无法将 'numx' (类型 myInt) 用作类型 int) 所以还需要转换 因为go认为myInt和int是两个不同类型
	fmt.Printf("numx 的类型是 %T ,值为 %v \n", numx, numx)
	fmt.Printf("numy 的类型是 %T ,值为 %v \n", numy, numy)

	// 给函数类型取别名 函数调用
	m := my(getSum, 500, 600)
	fmt.Println("m=", m)

	// 支持对函数返回值命名 在返回值类型列表对返回值进行命名 在函数体直接使用 return时无需返回
	sm, sb := getSumAndSub2(50, 60)
	// 使用 _ 标识符 ，忽略返回值
	sm, _ = getSumAndSub2(20, 30)
	fmt.Printf("sm= %v sb= %v \n", sm, sb)

	// 可变参数的函数使用
	totala := Sum1(10, 0, -1, 90, 10)
	fmt.Println("totala=", totala)

	var n1 float64 = 3
	var n2 float64 = 4
	var op byte = '+'

	// 函数调用
	// 			实参
	res := cal(n1, n2, op)
	fmt.Println("res =", res)

	// 调用返回多个返回值函数
	sum, sub := getSumAndSub(n1, n2)
	fmt.Printf("sum = %v sub = %v \n", sum, sub)
	// 调用返回多个返回值函数 希望忽略某个值 使用 _ 符号表示占位忽略
	_, sub2 := getSumAndSub(n1, n2)
	fmt.Println("sub2=", sub2)
}

// 函数定义 案例

//	函数名			形参 形参类型		 返回值类型
func cal(n1 float64, n2 float64, op byte) float64 {
	var res float64
	switch op {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("op error")
	}

	// 当函数有return 时，就是将结果返回给调用者
	return res

}

//
// go函数支持返回多个返回值
// 返回多个返回值 案例
// 编写一个函数，计算两个数的和和差
func getSumAndSub(n1 float64, n2 float64) (float64, float64) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

//  支持对函数返回值命名 在返回值类型列表对返回值进行命名 在函数体直接使用 return时无需返回  上面函数可以简化为
func getSumAndSub2(n1 float64, n2 float64) (sum float64, sub float64) {
	sum = n1 + n2
	sub = n1 - n2
	return // 不需要再return 返回值
}
