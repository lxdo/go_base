package main

import "fmt"

/**
字符串类型：string
基本介绍
字符串就是一串固定长度的字符连接起来的字符序列，Go的字符串是由单个字节连接起来的
go语言的字符串的字节使用utf8编码标识Unicode文本
*/
func main() {
	// string 基本使用
	var address string = "北京长城 hello world"
	fmt.Println(address)

	// 使用细节
	// 1.go语言字符串的字节使用utf8编码标识Unicode文本 ，不会出现乱码问题
	// 2.字符串一旦赋值了，字符串就不能修改了，在go中字符串是不可变的

	// address[0] = "a"  // 报错 这里就不能去修改address的内容 即go中的字符串是不可变的
	/**
	3.字符串的两种表现形式
	1).双引号 会识别转义字符
	2).反引号 以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果
	*/
	// 双引号
	str := "abc\nabc"
	fmt.Println(str)
	// 反引号
	str2 := `
		func main() {
			// string 基本使用
			var address string = "北京长城 hello world"
			fmt.Println(address)
		}
	`
	fmt.Println(str2)

	// 4.字符串拼接方式 +
	str3 := "hello" + " world"
	str3 += " haha"
	fmt.Println(str3)

	// 当一个拼接的操作很长时，可以分行写。但是注意，需要将+保留在上一行

	var str4 = "hello" + " world" +
		"hello" + " world" +
		"hello" + " world" +
		"hello" + " world" +
		"hello" + " world"
	fmt.Println(str4)

	/**
	基本数据类型默认值
	在go中，数据类型都有一个默认值 ，当程序员没有赋值时，就会保留默认值，
	在go中，默认值又叫零值
	*/
	/**
	数据类型         默认值

	整型              0
	浮点型            0
	字符串            ""
	布尔类型          false
	*/

	var a int       // 0
	var b float32   // 0
	var c float64   // 0
	var sex bool    // false
	var name string // ""

	// 这里的%v 表示按照变量的值输出
	fmt.Printf("a=%d,b=%v,c=%v,sex=%v,name=%v", a, b, c, sex, name)

}
