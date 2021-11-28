package main

import (
	"fmt"
	"strconv"
)

/**
基本数据类型和string的转换

介绍
	在程序开发中，我们经常需要将基本数据类型转成string类型。或者将string类型转成基本数据类型
*/
func main() {
	/**
	基本类型转string类型
		1.方式1：fmt.Sprintf("%参数",表达式)
				func Sprintf(format string, a ...interface{}) string
				Sprintf根据format参数生成格式化的字符串并返回该字符串。
			1) 参数需要和表达式的数据类型相匹配
			2) fmt.Sprintf()会返回转换后的字符串
	*/
	var n1 int = 99
	var n2 float64 = 23.456
	var n3 bool = true
	var n4 byte = 'h'
	var str string // 空的str

	str = fmt.Sprintf("%d", n1)
	fmt.Printf("str type %T str = %q \n", str, str)

	str = fmt.Sprintf("%f", n2)
	fmt.Printf("str type %T str = %q \n", str, str)

	str = fmt.Sprintf("%t", n3)
	fmt.Printf("str type %T str = %q \n", str, str)

	str = fmt.Sprintf("%c", n4)
	fmt.Printf("str type %T str = %q \n", str, str)

	/**
		2.方式2：使用strconv包的函数
	 */
	// int->string
	// 1) 把int64 转为string
	// 说明             转换成int64    10：进制（这里是10进制）
	str=strconv.FormatInt(int64(n1),10)
	fmt.Printf("str type %T str = %q \n", str, str)
	// 2) 把int转为string
	str=strconv.Itoa(n1)
	fmt.Printf("str type %T str = %q \n", str, str)


	// float->string
	// 说明                  'f':  输出格式  10: 精度(小数位保留位数 这里表示小数位保留10位)  64:表示这个小数是float64
	str=strconv.FormatFloat(n2,'f',10,64)
	fmt.Printf("str type %T str = %q \n", str, str)

	// bool->string
	str=strconv.FormatBool(n3)
	fmt.Printf("str type %T str = %q \n", str, str)


}
