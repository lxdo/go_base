package main

import "fmt"

/**
查找
	介绍:
		在go中,我们常用的查找有两种:
			1) 顺序查找
			2) 二分查找(前提:该数组是有序的)

*/
func main() {
	// 1) 顺序查找
	// 案例:有一个数列:小赵,小钱,小孙,小李。从键盘中任意输入一个名称,判断数列中是否包含此名称

	// 定义一个字符串数组
	names := [4]string{"小赵", "小钱", "小孙", "小李"}
	// 定义一个字符串变量接收名称
	var name string = "小张"

	// 顺序查找 把指定名称和数组里的名称循环对比 查到index改为名称对应的下标 未查到index不变
	index := -1
	for i, v := range names {
		if v == name {
			index = i // 将查到的值对应的下标赋值给index
			break
		}
	}
	if index == -1 {
		fmt.Println("未找到", name)
	} else {
		fmt.Printf("找到 %v 下标 %v \n", name, index)
	}

	// 2) 二分查找 前提是被查找的数组是有序的
	// 案例:对一个有序数列进行二分查找 {1,8,10,89,1000,1234} ,输入一个数看看该数组是否存在此数,如果存在,给出对应下标,如果没有,给出提示
	// 思路分析:二分查找分析.png
	arr := [6]int{1, 8, 10, 89, 1000, 1234}
	BinaryFind(&arr, 0, len(arr)-1, 1000)

}

// 二分查找的函数
func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	// 当leftIndex > rightIndex时 就是找不到了，结束递归函数
	if leftIndex > rightIndex {
		fmt.Println("未找到", findVal)
		return
	}

	// 先找到 中间的下标
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		// 说明要查找的数,应该在 leftIndex---middle-1
		BinaryFind(arr, leftIndex, middle-1, findVal) // 这里第1个参数不用传&arr,现在函数内arr本身就是指针变量
	} else if (*arr)[middle] < findVal {
		// 说明要查找的数,应该在 middle+1---rightIndex
		BinaryFind(arr, middle+1, rightIndex, findVal) // 这里第1个参数不用传&arr,现在函数内arr本身就是指针变量
	} else {
		// 找到了
		fmt.Println("下标为", middle, "值为", findVal)
	}
}
