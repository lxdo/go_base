package main

import (
	"fmt"
	"unsafe"
)

/**
小数类型分类

类型        占用存储空间     表数范围

单精度float32   4字节        -3.403E38-3.403E38
双精度float64   8字节        -1.798E308-1.798E308
*/

/**
说明
浮点类型有固定的范围和字段长度  不受具体os(操作系统)的影响
浮点类型默认声明为float64类型
浮点型常量有两种表现形式
	十进制数形式 ：如 5.12     .512 （必须有小数点）
	科学计数法形式：如 5.1234e2 =5.12*10的2次方 5.12E-2=5.12/10的2次方

在开发中 通常 使用float64

*/
func main() {
	var price float32 = 80.12
	fmt.Println(price)

	// 浮点数在机器中存放形式说明:浮点数=符号位+指数位+尾数位
	// 说明浮点数都是有符号的
	var num1 float32 = -0.00089
	var num2 float64 = -7809656.09
	fmt.Println(num1, num2)
	// 尾数部分可能丢失，造成精度损失
	// 说明 float64的精度比float32的要准确 要保存一个精度高的数 ，则应该选用float64位
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println(num3, num4)

	// 浮点类型默认声明为float64类型
	var num5 = 1.1
	fmt.Printf("num5 数据类型 %T 占用字节数 %d", num5, unsafe.Sizeof(num5))
	// 浮点型常量有两种表现形式
	// 十进制数形式 ：如 5.12     .512 （必须有小数点）
	num6 := 5.12
	num7 := .123 // 等价 0.123
	fmt.Println(num6, num7)
	// 科学计数法形式：如 5.1234e2 =5.12*10
	num8 := 5.1234e2   // 等价 5.1234*10的2次方
	num9 := 5.1234E2   // 等价 5.1234*10的2次方
	num10 := 5.1234E-2 // 等价 5.1234/10的2次方

	fmt.Println(num8, num9, num10)

}
