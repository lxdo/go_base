package main

import "fmt"

/**
for 循环控制

基本介绍：
	让代码可以循环的执行
基本语法：
	for 循环变量初始化;循环条件;循环变量迭代 {
		循环操作(循环体)
	}
说明:
	1) 循环四要素:
				1.循环变量初始化
				2.循环条件
				3.循环操作(循环体)
				4.循环变量迭代
	2) 循环执行顺序:
				1.执行 循环变量初始化  i := 1
				2.执行 循环条件  i <= 10
				3.执行 循环操作(循环体)  fmt.Println("hello world ", i)
				4.执行 循环变量迭代  i++
				5.反复执行 2 3 4 步骤 直到循环变量不符合循环条件 就退出for循环

细节:
	1) 循环条件是返回一个布尔值的表达式
	2) for 循环的第二种使用方式：

		for 循环判断条件 {
			循环执行语句
		}
		将变量初始化和变量迭代写到其他位置

	3) for 循环的第三种使用方式:

		for {
			循环执行语句
		}
		上面写法等价:
		for ; ; {
			循环执行语句
		}
		是一个无限循环，通常需要配合break语句使用
	4) go 提供 for range的方式，可以方便遍历字符串和数组

*/
func main() {
	//  i := 1 初始化循环变量
	//  i <= 10 循环条件
	//  fmt.Println("hello world ", i) 循环操作(循环体)
	//  i++ 循环变量迭代
	for i := 1; i <= 10; i++ { // for 循环中 定义的变量 只能在for循环中使用
		fmt.Println("hello world ", i)
	}

	// fmt.Println(i) 报错 未定义

	/*
		for 循环的第二种使用方式：

				for 循环判断条件 {
					循环执行语句
				}
				将变量初始化和变量迭代写到其他位置
	*/

	a := 1        // 循环变量初始化 在for循环外定义变量 在for循环里和外都可以使用
	for a <= 10 { // 循环条件
		fmt.Println("a= ", a) // 循环体
		a++                   // 循环变量迭代
	}

	/**
	for 循环的第三种使用方式:

		for {
			循环执行语句
		}
		上面写法等价:
		for ; ; {
			循环执行语句
		}
		是一个无限循环，通常需要配合break(跳出for循环,循环结束)语句使用
	*/

	k := 1
	for { // 这里也等价 for ; ; {
		fmt.Println("hello ", k)
		// 通常会配合break使用
		if k <= 10 {
			fmt.Println("world ", k)
		} else {
			break // break 就是跳出这个for循环,循环结束
		}
		k++
	}

	// go 提供 for range的方式，可以方便遍历字符串和数组
	// 遍历字符串
	// 1) 传统方式 按照字节遍历 (字符串有汉字需转[]rune 切片)
	var str string = "hello,world!"
	// i 在循环中定义 只能在循环中使用
	// len(str) 获取字符串长度
	// str[i] 用下标访问字符串单个字符 从0开始
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
	}

	/**
	注意:
		如果我们的字符串含有中文，那么传统的遍历字符串方式，会出现乱码
		传统的对字符串的遍历是按照字节来遍历，而一个汉字在utf8编码是3个字节(英文和数字是1个字节)
		解决：需要将字符串转成 []rune 切片

	*/
	str = "hello,world!中国" // 字符串中有中文，传统遍历默认按字节遍历，汉字占3个字节，会乱码
	str2 := []rune(str)    // 解决：将字符串转成 []rune 切片
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c \n", str2[i])
	}

	// 2) for range方式遍历 按照字符遍历 (字符串有汉字无需转换)
	str = "abcdef"
	str = "abcdef中国"
	for index, val := range str {
		fmt.Printf("index= %d val= %c \n", index, val)
	}

	// 练习
	// 1) 打印1~100之间所有是9的倍数的整数的个数及总和
	var max int = 100
	var count int = 0
	var sum int = 0
	for i := 1; i <= max; i++ {
		if i%9 == 0 {
			count++
			sum += i
		}
	}
	fmt.Printf("count=%v sum=%v \n",count,sum)
}
