package main

import "fmt"

/**
while 和 do...while 的实现

go 没有while和do...while语法
如果需要使用类似其他语言的while和do...while
可以通过for循环来实现
*/

/**
for循环实现while的效果:

			循环变量初始化
			for {
				if 循环条件表达式 {
					break // 跳出for循环
				}
				循环操作(循环体)
				循环变量迭代
			}

说明：
	1) 上面for循环是一个无限循环
	2) break 语句就是跳出for循环
*/

/**
for循环实现do...while的效果:

			循环变量初始化
			for {
				循环操作(循环体)
				循环变量迭代
				if 循环条件表达式 {
					break // 跳出for循环
				}
			}

说明:
	1) 上面的循环是先执行，再判断，因此至少执行一次
	2) 当循环条件成立后，break结束循环
*/

func main() {
	// for循环实现while
	var i int = 1 // 循环变量初始化
	for {
		if i > 10 { // 循环条件
			break // 跳出for循环，结束循环
		}
		fmt.Println("hello ", i)
		i++ // 循环变量的迭代
	}
	fmt.Println("i= ", i)

	// for循环实现do...while
	i = 1
	for {
		fmt.Println("world ", i)
		i++
		if i > 10 {
			break // 跳出for循环
		}
	}
	fmt.Println("i= ", i)
}
