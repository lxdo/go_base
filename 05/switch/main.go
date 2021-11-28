package main

import "fmt"

/**
switch 分支结构
	基本介绍：
		1) switch 语句用于基于不同条件执行不同动作，每一个case分支都是唯一的，从上到下逐一测试，直到匹配为止
		2) 匹配项后面也不需要再加break
	基本语法:
			switch 表达式 {
				case 表达式1,表达式2,...:  // switch 表达式的值 符合 表达式1或表达式2的值 就执行语句块1
					语句块1

				case 表达式3,表达式4,...:
					语句块2

				// 这里可以有多个case语句

				default:
					语句块
			}
	说明：
		1)：go 的case 后表达式可以有多个，使用逗号间隔
		2): go 中的case语句块不需要写break ，默认会有，
			即在默认情况下，当程序执行完case语句块后，就直接退出该switch控制结构

	细节：
		1) switch/case后是一个表达式 (即：常量值、变量、一个有返回值的函数等都可以)
		2) case 后的各个表达式的值的数据类型，必须和switch的表达式数据类型一致
		3) case 后面可以带多个表达式 ，使用逗号间隔。比如 case 表达式1,表达式2:
		4) case 后面的表达式如果是常量值(字面量)，则要求不能重复
		5) case 后面不需要带break,程序匹配到一个case后就会执行对应的代码块，然后退出switch，
			如果一个都匹配不到，则执行default
		6) default 语句不是必须的
		7) switch 后也可以不带表达式，类似if-else分支来使用
		8) switch 后也可以直接声明/定义一个变量，分号结束  (不推荐使用)
		9) switch 穿透 fallthrough ，如果在case语句块后增加 fallthrough，则会继续执行下一个case,也叫switch穿透
			穿透不会进行判断，直接执行下一个case,默认只能穿透一层
		10) Type Switch :switch 语句还可以被用于type-switch 来判断某个interface变量中实际指向的变量类型

*/

/**
switch和if的比较
	1) 如果判断的具体数值不多，而且符合整数、浮点数、字符、字符串这几张类型，建议使用switch语句，简洁高效
	2) 其他情况：对区间判断和结果为bool类型的判断，使用if，if的使用范围更广
*/
func main() {

	// 基本语法
	var key byte
	fmt.Println("请输入一个字符 a,b,c,d")
	fmt.Scanf("%c", &key)
	var c byte = 'c'
	switch test(key) { // 表达式可以是常量值、变量、一个有返回值的函数等
	case 'a': // 这里是 常量值(字面量)
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case c: // 这里是变量
		// case 后的各个表达式的值的数据类型，必须和switch的表达式数据类型一致
		// 这里变量c数据类型要和test(key)的值的数据类型一致
		fmt.Println("周三")
	default:
		fmt.Println("输入有误")
	}

	var n1 int32 = 5
	var n2 int32 = 20
	switch n1 {
	case n2, 10, 5: // case 后面可以有多个表达式
		fmt.Println("ok1")

		// case 5: // 报错 case 后面的表达式如果是常量值(字面量)，则要求不能重复 5是常量值已经在上面存在
		fmt.Println("ok2")

	default:
		fmt.Println("没有匹配到")

	}

	// switch 后也可以不带表达式，类似if-else分支来使用
	var age int = 10
	switch {
	case age == 10:
		fmt.Println("age=10")
	case age == 20:
		fmt.Println("age=20")
	default:
		fmt.Println("没有值")

	}

	// switch 后也可以不带表达式，类似if-else分支来使用
	// case 中也可以对范围进行判断
	var score int = 90
	switch {
	case score >= 90:
		fmt.Println("score优秀")
	case score >= 80:
		fmt.Println("score良好")
	default:
		fmt.Println("score一般")

	}

	// switch 后也可以直接声明/定义一个变量，分号结束 (不推荐使用)

	switch grade := 90; {
	case grade >= 90:
		fmt.Println("grade优秀")
	case grade >= 80:
		fmt.Println("grade良好")
	default:
		fmt.Println("grade一般")

	}

	// switch 穿透 fallthrough ，如果在case语句块后增加 fallthrough，则会继续执行下一个case,也叫switch穿透
	// 穿透不会进行判断，直接执行下一个case,默认只能穿透一层
	var num int = 10
	switch num {
	case 10:
		fmt.Println(10)
		fallthrough //  穿透不会进行判断，直接执行下一个case,默认只能穿透一层
	case 20:
		fmt.Println(20)
	case 30:
		fmt.Println(30)
	default:
		fmt.Println("。。。")

	}

	// 10) Type Switch :switch 语句还可以被用于type-switch 来判断某个interface变量中实际指向的变量类型
	var x interface{} // 空接口
	var y = 10.0      // float64 (默认)
	x = y

	switch i := x.(type) { // x.(type) 获取x指向的数据类型
	case nil:
		fmt.Printf("x 的类型：%T \n", i)
	case int:
		fmt.Printf("x 是int型 \n")
	case float64:
		fmt.Printf("x 是float64型 \n")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型 \n")
	case bool, string:
		fmt.Printf("x 是 bool 或 string型")

	default:
		fmt.Printf("未知型")
	}
}

func test(key byte) byte {
	return key
}
