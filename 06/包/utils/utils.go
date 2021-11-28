package utils

import "fmt"

/**
为了让其他包的文件使用Cal函数，需要将c大写 类型其他语言的public
go中叫该函数可导出
*/

//	函数名			形参 形参类型		 返回值类型
func Cal(n1 float64, n2 float64, op byte) float64 {
	var res float64
	switch op {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("op error")
	}

	return res
}

var Num int = 3  // 变量首字母大写 可以在其他包被调用
