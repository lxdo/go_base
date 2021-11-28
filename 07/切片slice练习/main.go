package main

import "fmt"

/**
切片练习
*/
func main() {
	// 1)切片拷贝练习
	var a []int = []int{1, 2, 3, 4, 5} // 被拷贝切片有5个元素
	var slice = make([]int, 2)         // 而定义时拷贝切片只有2个元素
	fmt.Println("slice=", slice)       // slice= [0 0]
	copy(slice, a)                     // 不会报错 当len(a)>len(slice) 只会拷贝slice能装下的元素
	fmt.Println("slice=", slice)       // slice= [1 2]

	// 切片是引用类型，所以在传递时，遵守引用传递机制
	// arr slice2 slice3都指向同一个数据空间
	var slice2 []int
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	slice2 = arr[:]
	var slice3 = slice2
	slice3[0] = 10

	fmt.Println("slice2=", slice2)
	fmt.Println("slice3=", slice3)
	fmt.Println("arr=", arr)

	// 切片是引用类型，所以在传递时，遵守引用传递机制
	var s = []int{1, 2, 3, 4}
	fmt.Println("s=", s)
	// 切片是引用类型，在函数内修改切片会同步修改函数外对应的实参
	test(s)
	fmt.Println("s=", s)

	/**
	编写一个函数 fbn(n int),要求完成
		1) 可以接收一个 n int // n为需要获取的斐波那契数列的元素个数
		2) 能够将斐波那契的数列放到切片中
		3) 提示，斐波那契的数列形式：
		arr[0]=1 arr[1]=1 arr[2]=2 arr[3]=3 arr[4]=5 arr[5]=8

	思路：
		1) 声明一个函数 fbn(n int) []uint64  函数返回一个切片
		2) 在fbn函数进行for循环来存放斐波那契数列 0->1 1->1
	*/
	fnbSlice := fbn(10)
	fmt.Println("fnbSlice=", fnbSlice)

}

// 切片是引用类型，在函数内修改元素，函数外的切片也同步变化
func test(slice []int) {
	slice[0] = 100
}

// 因为斐波那契数列会越来越大,所以返回值数据类型给到最大 []uint64
// 参数n为获取的斐波那契数列的元素个数
func fbn(n int) []uint64 {
	// 声明一个切片,切片大小n
	fbnSlice := make([]uint64, n)
	// 第一个数和第二个数的斐波那契为1
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	// 进行for循环来存放斐波那契的数列
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}

	return fbnSlice
}
