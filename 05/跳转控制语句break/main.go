package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
跳转控制语句break
基本介绍：
	break语句用于终止某个语句块的执行，用于中断当前for循环或跳出switch语句
	(当执行到break时，就中断跳出当前的for循环)

基本语法:
		{
			...
			break
		}

细节：
	1) break 语句出现在多层嵌套的语句块中时，可以通过标签指明要跳出哪一层循环
	(嵌套循环中break默认跳出最近的for循环)
	2) 标签的基本使用

*/
func main() {
	// 例：随机生成1-100的一个数，直到生成了99这个数，看看一共用了几次
	// 分析：写一个无限循环控制，不停生成随机数，当生成了99时，退出循环(break)

	/**
	补充：如何随机生成随机数
			1) 为了生成一个随机数，需要先给rand设置一个种子(在go中，需要先生成一个随机种子，否则返回的值总是固定的)
				rand.Seed(time.Now().Unix())  // 1秒随机1次
				rand.Seed(time.Now().UnixNano()) // 1纳秒随机1次
				// time.Now().Unix() 返回一个unix时间戳=php中的time() (为1970 01-01 00:00:00到现在的秒数)
				// time.Now().UnixNano() 返回一个unix时间戳 (为1970 01-01 00:00:00到现在的纳秒数)
			2) rand.Intn(n)  获取随机数  范围:  [0,n)


	*/
	/**
	如何随机生成1-100的一个数
		1) 先给rand设置一个种子 	rand.Seed(time.Now().UnixNano()) // 1纳秒随机一次 粒度更细
		2) rand.Intn(100) 在 [0,100) 生成随机数(不包括100)
			随机生成1-100的一个数 即rand.Intn(100)+1  范围为: [1,101)
	*/

	rand.Seed(time.Now().UnixNano()) // 给rand设置种子 1纳秒随机一次 粒度更细
	n := rand.Intn(100) + 1          // 生成 [1,101)即 [1,100]随机数
	fmt.Println("随机数 n = ", n)
	fmt.Println("当前时间戳 秒 ", time.Now().Unix())
	fmt.Println("当前时间戳 纳秒 ", time.Now().UnixNano())

	// 随机生成1-100的一个数，直到生成了99这个数，看看一共用了几次
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano()) // 给rand设置种子 1纳秒随机一次 粒度更细
		num := rand.Intn(100) + 1        // 生成 [1,101)即 [1,100]随机数
		count++
		if num == 99 {
			break // 跳出循环
		}
	}
	fmt.Println("生成99一共使用了 ", count)

	/**
	1) break 语句出现在多层嵌套的语句块中时，可以通过标签指明要跳出哪一层循环
	(嵌套循环中break默认跳出最近的for循环)
	(break 标签 ：直接结束标签对应的for循环)
	2) 标签的基本使用
	*/

label2: // 设置一个标签，名称是自定义的 绑定循环
	for i := 0; i < 4; i++ {
		//label1: // 设置一个标签，名称是自定义的 绑定循环
		for j := 0; j < 10; j++ {
			if j == 2 {
				// break // 默认跳出最近的for循环
				//break label1  // 跳出 lable1标签 对应的for循环
				break label2 // 跳出 label2标签 对应的for循环
			}
			fmt.Println("j=", j)
		}
	}

	/**
	说明：
		1) break 默认会跳出最近的for循环
		2) break 后面可以指定标签 ，跳出标签对应的for循环
	*/

}


