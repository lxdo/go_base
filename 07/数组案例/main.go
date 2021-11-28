package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
数组案例
*/
func main() {
	// 1) 创建一个byte类型的26个元素的数组，分别放置"A"-"Z"。使用for循环访问所有元素并打印出来(提示：字符数据运算 "A"+1="B")

	// 声明一个数组
	var chars [26]byte

	// 使用for循环，利用字符可以进行运算的特点来赋值 (提示：字符数据运算 "A"+1="B")
	for i := 0; i < 26; i++ {
		chars[i] = 'A' + byte(i) // 'A'是byte类型 i是int类型 需要强转
	}

	fmt.Println("chars=", chars) // 直接输出 数组元素是字符对应的ASCII码
	for _, value := range chars {
		fmt.Printf(" %c ", value) // %c 按照字符输出 输出ASCII码对应的字符
	}

	fmt.Println()
	// 2) 请求出一个数组的最大值，并得到对应的下标

	// 声明一个数组
	nums := [...]int{1, -1, 9, 90, 13}

	// 假设第一个元素就是最大值，下标为0
	maxVal := nums[0]
	maxValIndex := 0

	// 然后从第二个元素开始循环比较，如果发现有更大的，则交换
	for i := 1; i < len(nums); i++ {
		if maxVal < nums[i] {
			maxVal = nums[i]
			maxValIndex = i
		}
	}
	fmt.Printf("maxIndex= %v maxVal= %v \n", maxValIndex, maxVal)

	// 3) 求出一个数组的和、平均值

	// 声明一个数组
	// 用nums数组

	// 累计求和
	sum := 0
	for _, val := range nums {
		sum += val
	}
	// sum和 len(nums)两个都为int型，相除结果只保留整数，小数会被去掉 需要都转成float64相除
	fmt.Printf("sum=%v avg=%v \n", sum, float64(sum)/float64(len(nums)))

	// 4) 随机生成5个数，并将其反转打印

	// 声明一个数组
	var arr1 [5]int
	// 随机生成5个数
	// 为了每次生成的随机数不一样，我们需要给一个seed值
	len := len(arr1) // 经常需要用到计算数组len,用变量接收
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		arr1[i] = rand.Intn(100) // 取随机数范围 0<=n<100
	}
	fmt.Println("交换前 arr1=", arr1)
	// 反转打印，交换的次数是 len(arr1)/2 ，第一个和倒数第一个交换 第二个和倒数第二个交换
	for i := 0; i < len/2; i++ {
		// a,b=b,a 实现 变量值交换
		arr1[i], arr1[len-1-i] = arr1[len-1-i], arr1[i]
	}
	fmt.Println("交换后 arr1=", arr1)

}
