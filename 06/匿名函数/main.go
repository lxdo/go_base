package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
匿名函数

介绍：
	go支持匿名函数(没有名字的函数)，如果我们某个函数只是希望使用一次，可以考虑使用匿名函数，匿名函数也可以实现多次调用

匿名函数使用方式：
	1) 在定义匿名函数时就直接调用,这种方式匿名函数只能调用一次
	2) 将匿名函数赋给一个变量 (函数变量)，再通过该变量来调用匿名函数
	3) 全局匿名函数
		如果将匿名函数赋给一个全局变量，那么这个匿名函数，就成为一个全局匿名函数，可以在整个程序有效

*/

// 全局匿名函数：如果将匿名函数赋给一个全局变量，那么这个匿名函数，就成为一个全局匿名函数，可以在整个程序有效(变量名首字母大写)
var (
	// mult就是一个全局匿名函数
	Mult = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	// 在定义匿名函数时就直接调用,这种方式匿名函数只能调用一次
	sum := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20) // 定义时即调用 调用结果返回给变量sum

	fmt.Println("sum=", sum)

	// 将匿名函数赋给一个变量 (函数变量)，再通过该变量来调用匿名函数、
	// 将匿名函数 func(n1 int,n2 int) int 赋给 变量sub 则sub的数据类型就是函数类型，此时我们可以通过sub完成调用
	sub := func(n1 int, n2 int) int {
		return n1 - n2
	}

	// 通过函数变量sub调用匿名函数
	a := sub(30, 30)
	fmt.Println("a=", a)
	// 匿名函数赋给一个变量可实现多次调用
	b := sub(50, 40)
	fmt.Println("b=", b)

	// 调用全局匿名函数
	c := Mult(4, 9)
	fmt.Println("c=", c)



	fmt.Println(buyNum())
}


//比较两个切片 值是否一致
//例1：[1,2,3,3] [2,1,2,3] 返回true
//例2：[1,2,3]   [2,1,2,3] 返回true
//例3：[5,2,3,2] [2,1,2,3] 返回false
func SliceComparison(a, b []string) bool {
	aMap := map[string]bool{}
	bMap := map[string]bool{}
	for _, v := range a {
		aMap[v] = true
	}
	for _, v := range b {
		bMap[v] = true
	}
	if len(aMap) != len(bMap) {
		return false
	}
	for k, _ := range aMap {
		if _, ok := bMap[k]; !ok {
			return false
		}
	}
	return true
}

func buyNum() []string {
	// 红区号码
	nums := bubbleSort(randNum(numSlice(33), 6))
	// 蓝区号码
	nums = append(nums, randNum(numSlice(16), 1)[0])
	return nums
}

/**
随机取数
*/
func randNum(numSlice []string, count int) []string {
	var nums []string
	// 取随机数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

label:
	for {
		// 从数字范围中随机取数
		num := numSlice[r.Intn(len(numSlice))]
		// 验证是否重复
		for _, v := range nums {
			if v == num {
				continue label
			}
		}

		nums = append(nums, num)

		if len(nums) == count {
			break
		}
	}

	return nums
}

/**
冒泡排序
*/
func bubbleSort(nums []string) []string {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

/**
@desc 数字范围
*/
func numSlice(max int) []string {
	var nums []string
	for i := 1; i <= max; i++ {
		nums = append(nums, fmt.Sprintf("%02d", i))
	}
	return nums
}
