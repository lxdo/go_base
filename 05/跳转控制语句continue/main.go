package main

import "fmt"

/**
跳转控制语句 continue

基本介绍：
	1) continue 语句用于结束本次循环，继续执行下一次循环 (当执行到continue时，就不再执行下面的代码，而是直接下一次循环)
	2) continue 语句出现在多层嵌套的循环语句体中时，可以通过标签指明要跳过的是哪一层循环，这个和前面的break标签的使用规则一样
		(continue默认结束最近for循环的本次循环，进入最近for循环的下次循环)
		(continue 标签： 结束 标签 对应for循环的本次循环 进入标签对应for循环的下次循环)

基本语法：
	{
		...
		continue
	}
*/
func main() {
label2: // 设置一个标签，名称是自定义的 绑定循环
	for i := 0; i < 4; i++ {
		//label1: // 设置一个标签，名称是自定义的 绑定循环
		for j := 0; j < 4; j++ {
			if j == 2 {
				// continue    // 默认 结束本次循环 继续进入最近for循环下一次循环
				//continue label1 // 结束label1标签对应的for循环的本次循环 ，进入label1标签对应的for循环的下次循环
				continue label2 // 结束label2标签对应的for循环的本次循环 ，进入label2标签对应的for循环的下次循环
			}
			fmt.Println("j=", j)
		}
	}

	// 练习
	// 1) 打印1-10之内的奇数
	for i := 0; i <= 10; i++ {
		if i%2 == 0 { // %2=0是偶数
			continue
		}
		fmt.Printf(" 奇数是%v", i)
	}

	fmt.Println() // 输出换行
	// 2) 从键盘读入个数不确定的整数，并判断读入的正数和负数的个数，输入为0时结束程序
	var positiveCount int // 正数的个数
	var negativeCount int // 负数的个数
	var num int
	for {
		fmt.Println("请输入一个整数")
		fmt.Scanln(&num)
		if num == 0 {
			break // 终止for循环
		}
		if num > 0 {
			positiveCount++
			continue // 结束本次循环，进入下次循环
		}
		negativeCount++
	}
	fmt.Printf("正数个数是%v 负数个数是%v \n", positiveCount, negativeCount)

}
