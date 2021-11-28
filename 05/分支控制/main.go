package main

import "fmt"

/**
分支控制 if else
	介绍：让程序有选择的执行，分支控制有三种
		1) 单分支
		2) 双分支
		3) 多分支
*/

/**
1)单分支
	基本语法：
		if 条件表达式 {
			执行代码块
		}
	说明：当条件表达式为 true 时，就会执行 { } 的代码
	注意：{} 是必须有的,就算只写一行代码
	细节：go的if还有一个强大的地方就是条件判断语句里面允许声明一个变量，
         这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了


*/

/**
2)双分支
	基本语法:
		if 条件表达式 {
			执行代码块1
		} else {
			执行代码块2
		}
	说明:
		当条件表达式成立(true)，即执行代码块1，否则执行代码块2  { }也是必须有的
		双分支只会执行其中一个分支代码块

*/

/**
3) 多分支
	基本语法：
		if 条件表达式1 {
			执行代码块1
		} else if 条件表达式2 {
			执行代码块2
		} else {
			执行代码块n
		}
	说明：
		1.else {} 不是必须的
		2.多分支只会执行其中一个分支代码块

*/

/**
4)嵌套分支：
	基本介绍:
		在一个分支结构中又完整的嵌套了另一个完整的分支结构，里面的分支的结构称为内层分支 外面的分支结构称为外层分支
	基本语法:
		if 条件表达式 {
			if 条件表达式 {
			}else{
			}
		}
	说明：嵌套分支不宜过多，建议控制在3层内

*/
/**
编写规范：
	1.	if 条件表达式 {
			执行代码块
		}
	{} 是必须有的,就算只写一行代码

	2. if (条件表达式) {
		执行代码块
	}
	条件表达式 加() 也不会报错，但是建议不要加()

	3.  if 条件表达式 {
			执行代码块1
		}
		else{     // 编译错误 else 换行写是错误的
			执行代码块2
		}
		else 换行写是错误的，会报错

	4.	if a=4 {  // 编译错误 if的条件表达式不能是赋值语句
			xxx
		}


*/
func main() {
	// 1) 单分支 基本语法
	var age byte
	fmt.Println("请输入年龄:")
	fmt.Scanln(&age)

	// 当条件表达式为 true 时，就会执行 { } 的代码
	if age > 18 {
		fmt.Println("你年龄大于18岁")
	}

	// if条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了
	if a := 20; a > 18 { // 可以在条件表达式和执行代码块中使用 a
		fmt.Println("a>18 a=", a)
	}
	// fmt.Println("a = ", a)  // 报错 a 未定义

	// 2) 双分支 基本语法
	if age > 18 {
		fmt.Println("大于18岁")
	} else { // else 不能换行
		fmt.Println("你还年轻")
	}

	// 3) 多分支 基本语法
	var score int
	fmt.Println("请输入成绩:")
	fmt.Scanln(&score)

	if score == 100 {
		fmt.Println("奖励一辆BMW")
	} else if score > 80 && score <= 99 {
		fmt.Println("奖励一台iPhone7plus")
	} else if score >= 60 && score <= 80 {
		fmt.Println("奖励一个ipad")
	} else {
		fmt.Println("什么都没有")
	}

	// 多分支只会执行其中一个分支代码块
	var n int = 10
	if n > 9 { // 该条件表达式成立 执行该代码块 即使后面条件表达式也成立 但不会执行
		fmt.Println("ok1")
	} else if n > 6 {
		fmt.Println("ok2")
	} else if n > 3 {
		fmt.Println("ok3")
	} else {
		fmt.Println("ok4")
	}

	// 4) 嵌套分支
	var second float64 = 5

	if second <= 8 {
		var gender string = "男"
		if gender == "男" {
			fmt.Println("进入决赛男子组")
		} else {
			fmt.Println("进入决赛女子组")
		}
	} else {
		fmt.Println("淘汰")
	}
	// 练习
	// 1. 定义两个变量int32，判断二者的和，是否能被3又能被5整除 (被x整除:取余为0)
	var num1 int32 = 10
	var num2 int32 = 5
	if (num1+num2)%3 == 0 && (num1+num2)%5 == 0 {
		fmt.Println("能被3又能被5整除")
	}

}
