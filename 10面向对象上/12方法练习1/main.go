package main

import "fmt"

/**
方法的练习
*/

// 编写结构体(MethodUtils)
type MethodUtils struct {
}

// 1.编写一个方法，方法不需要参数，在方法中打印10*8的矩形,在main方法中调用该方法
// 给MethodUtils编写方法
func (mu MethodUtils) print() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// 2.编写一个方法，提供m和n两个参数，方法中打印一个m*n的矩形
func (mu MethodUtils) print2(m int, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// 3.编写一个方法算该矩形的面积(可以接收长len,和宽width)。将其作为方法返回值，在main方法中调用该方法，接收返回的面积值并打印
func (mu MethodUtils) area(len float64, width float64) float64 {
	return len * width
}
func main() {
	// 练习1
	var mu MethodUtils
	mu.print()
	// 练习2
	mu.print2(3, 10)
	// 练习3
	area := mu.area(5.2, 1.8)
	fmt.Println("area=", area)
}
