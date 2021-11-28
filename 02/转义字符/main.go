package main

import "fmt"

/**


转义字符的使用
1. \t  一个制表位，实现对齐的功能 tab键
2. \n  换行符
3. \\  一个\
4. \"  一个"
5. \r  一个回车
*/
// fmt包中提供格式化，输出，输入的函数

func main() {
	fmt.Println("hello\tworld")
	fmt.Println("he\nwo")
	fmt.Println("hello\\world")
	fmt.Println("tom说\"hello world\"")
	// \r     回车 从当前行的最前面开始输出，覆盖掉以前的内容
	fmt.Println("天龙八部\r张飞")
}

//注释
// 行注释 （官方推荐）
/**
块注释
*/


/*
代码格式化命令 gofmt -w main.go
该指令可以将格式化后的内容重写入到文件
 */


/*
go环境变量配置及作用
GOROOT 指定 go sdk 安装目录 自生成
GOPATH 就是golang工作目录 我们所有项目源码都在这个目录下
*/
