package main

import (
	"fmt"
	"unsafe"
)

/**
布尔类型 ： bool
	基本介绍：
		布尔类型也叫bool类型，bool类型数据只允许取值true和false
		bool类型占1个字节
		bool类型适于逻辑运算，一般用于程序流程控制
			if条件
			for循环
 */
func main()  {
	var b =false
	fmt.Println("b=",b)

	// bool类型占存储空间1个字节
	fmt.Println("b 的占用空间=",unsafe.Sizeof(b))

	// bool类型数据只允许取值true和false
	// b=1 报错
}
