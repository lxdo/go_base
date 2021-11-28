package main

import "fmt"

/**
函数练习

总结：
	1) 函数多个形参数据类型一致，可以简写为: n1,n2,n3 数据类型
*/

// 函数多个形参数据类型一致，可以简写为: n1,n2 数据类型
func sum(n1, n2 float32) float32 {
	fmt.Printf("n1 type= %T \n", n1)
	return n1 + n2

}

// 编写函数，可以交换n1和n2的值
func swap(n1 *int, n2 *int) {
	*n1, *n2 = *n2, *n1 // 指针变量交换值
}

func main() {
	fmt.Println("sum=", sum(1, 2))

	n1 := 50
	n2 := 20
	fmt.Printf("n1=%v n2=%v \n", n1, n2)
	swap(&n1, &n2) // 以值传递方式 改变函数外变量值
	fmt.Printf("n1=%v n2=%v \n", n1, n2)

	// 调用打印金字塔函数
	printPyramid(9)

	// 调用打印乘法表函数
	printMulti(9)
}

// 用函数封装打印金字塔

func printPyramid(totalLevel int) {
	// i 表示层数
	for i := 1; i <= totalLevel; i++ {
		// 在打印*前先打印空格
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}
		// j 表示每层打印多少 *
		for j := 1; j <= (2*i - 1); j++ {
			if j == 1 || j == 2*i-1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}
}

// 用函数封装打印乘法表
func printMulti(num int) {
	// i 表示层数
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, j*i)
		}
		fmt.Println()
	}
}


