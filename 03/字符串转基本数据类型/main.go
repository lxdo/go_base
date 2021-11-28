package main

import (
	"fmt"
	"strconv"
)

/**
string 转成基本数据类型
*/
func main() {
	// string -> bool
	var str string = "true"
	var b bool
	// 1.strconv.ParseBool(str) 返回两个值 (value bool,err error)
	// 2.只获取value bool，不想获取err 所以使用_忽略
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v \n", b, b)

	// string -> int64
	var str2 string = "1234590"
	var n1 int64
	//                       字符串    进制        整数类型
	n1, _ = strconv.ParseInt(str2, 10, 64) // int64 error

	fmt.Printf("n1 type %T n1=%v \n", n1, n1)

	// 注意：因为返回的是int64或float64 ，如果希望得到int32,float32等需要转换类型
	var n2 int
	n2 = int(n1)
	fmt.Printf("n2 type %T n2=%v \n", n2, n2)

	// string -> float64
	var str3 string = "123.456"
	var f1 float64
	// 字符串  精度32/64
	f1, _ = strconv.ParseFloat(str3, 64) // float64 error
	fmt.Printf("f1 type %T f1=%v \n", f1, f1)

	// 注意：因为返回的是int64或float64 ，如果希望得到int32,float32等需要转换类型

	/**
	在将string类型转换成基本数据类型时，要确保string类型能够转成有效的数据，
	比如我们可以把"123"，转成一个整数，但是不能把"hello"转成一个整数，如果这样做
	go直接将其转成0
	*/

	str = "hello"
	var num int64
	num = 12
	num, _ = strconv.ParseInt(str, 10, 64)
	// 如果没有转换成功，即使num之前有值，也会被转换为0
	fmt.Printf("num type %T num=%v \n", num, num)
	// bool型转换失败默认false
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v \n", b, b)
	// float -> 0
}
