package main

import "fmt"

/**
方法练习
*/

// 定义结构体
type Men struct {
}

// 1.编写方法，判断一个数是奇数还是偶数
// 方法和结构体的指针变量绑定
func (me *Men) Judg(num int) string {
	if num%2 == 0 {
		return "偶数"
	}
	return "奇数"
}


func main() {
	// 练习1
	var me Men
	res := (&me).Judg(3) // 标准写法 等价于  me.Judg(3)
	fmt.Println("3是", res)
	res = me.Judg(6) // 简单写法 等价于 (&me).Judg(6)
	fmt.Println("6是", res)

	// 练习2
	/**
	编写方法，使给定的一个二维数组(3*3)转置:
	1 2 3      1 4 7
	4 5 6  ==> 2 5 8
	7 8 9      3 6 9
	 */
}
