package main

import "fmt"

/**
数组：
	介绍：
		数组可以存放多个同一类型数据。数组也是一种数据类型，在go中，数组是值类型

*/
func main() {
	// 数组快速入门

	// 1.定义一个数组
	var hens [6]float64 // 定义一个有6个元素，元素数据类型为float64的数组

	// 2.给数组的每个元素赋值 ,元素的下标从0开始
	hens[0] = 3.0 // hens数组的第一个元素 hens[0] 下标从0开始
	hens[1] = 5.0 // hens数组的第二个元素 hens[1]
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0

	// 3.遍历数组求出元素的和
	total := 0.0 // float64 类型保持一致
	for i := 0; i < len(hens); i++ {
		total += hens[i]
	}

	// 4.求出元素的平均值
	// Sprintf根据format参数生成格式化的字符串并返回该字符串。
	// %.2f   默认宽度，精度2  保留两位小数
	avg := fmt.Sprintf("%.2f", total/float64(len(hens))) // 将结果格式化  保留两位小数 赋值给变量
	fmt.Printf("total= %v avg= %v \n", total, avg)

	/**
	数组的定义
		var 数组名 [数组大小]数据类型  // 数组大小:数组里有几个元素  数据类型:数组元素的数据类型
		var a [5]int
		赋初值 a[0]=1 a[1]=30 ...
	*/
	var intArr [3]int   // 当我们定义完数组后，数组各个元素有默认值 0
	fmt.Println(intArr) // [0 0 0]

	// 数组的地址
	fmt.Printf("intArr地址= %p inArr[0]地址= %p intArr[1]地址= %p intArr[2]地址= %p \n",
		&intArr, &intArr[0], &intArr[1], &intArr[2])
	// intArr地址= 0xc000018100 inArr[0]地址= 0xc000018100 intArr[1]地址= 0xc000018108 intArr[2]地址= 0xc000018110(十六进制 十六进一)
	// 1.数组的地址可以通过数组名来获取 &intArr
	// 2.数组的地址就是数组第一个元素的地址,即数组的首地址
	// 3.数组第二个元素地址就是第一个元素的地址加上数组数据类型(int 占8字节)所占用的字节数 ，依次类推
	//	(即：数组各个元素的地址间隔是依据数组的数据类型决定，比如 int64->8 int32->4)

	/**
	数组的使用
		访问数组元素：
			数组名[下标] ： intArr[2]
	*/

	// 数组元素的访问和赋值
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Println("intArr=", intArr)

	// 四种初始化数组的方式
	var nums1 [3]int = [3]int{1, 2, 3}                      // 定义时即赋值
	var nums2 = [3]int{1, 2, 3}                             // 定义时即赋值
	var nums3 = [...]int{6, 7, 8, 9}                        // 让系统判断 数组元素的个数
	var names1 = [3]string{1: "tom", 0: "jack", 2: "marry"} // 指定数组元素对应的下标
	names2 := [...]string{1: "tom", 0: "jack", 2: "marry"}  // 类型推导
	fmt.Println("nums1=", nums1)                            // nums1= [1 2 3]
	fmt.Println("nums2=", nums2)                            // nums2= [1 2 3]
	fmt.Println("nums3=", nums3)                            // nums3= [6 7 8 9]
	fmt.Println("names1=", names1)                          // names1= [jack tom marry] 会自动将下标排序后输出
	fmt.Println("names2=", names2)                          // names2= [jack tom marry] 会自动将下标排序后输出

	/**
	数组的遍历
		for range结构：
			这是go一种独有的结构，可以用来遍历访问数组的元素
		基本语法
		for index,value:=range arr {
			...
		}

		1) 第一个返回值 index是数组的下标
		2) 第二个value是在该下标位置的值
		3) 他们都是仅在for循环内部可见的局部变量
		4) 遍历数组元素的时候，如果不想使用下标index,可以直接把下标index标为下划线 _
		5) index和value的名称不是固定的，即程序员可以自行指定，一般命名为index和value
	*/

	heroes := [...]string{"宋江", "吴用", "卢俊义"}
	for index, value := range heroes {
		fmt.Printf("index= %v value= %v ", index, value)
	}
	fmt.Println()
	// 遍历数组元素的时候，如果不想使用下标index,可以直接把下标index标为下划线 _
	for _, value := range heroes {
		fmt.Printf(" value= %v ", value)
	}
	fmt.Println()
	/**
	数组使用注意事项和细节
		1) 数组是多个相同类型数据的组合，一个数组一旦声明/定义了，其长度是固定的，不能动态变化
		2) var arr []int ,这时arr就是一个slice切片,不是数组
		3) 数组中的元素可以是任何数据类型，包括值类型和引用类型，但是不能混用。
		4) 数组创建后，如果没有赋值，有默认值(零值 )
			数值类型数组(整数系列、浮点数系统)：默认值为0
			字符串数组：默认值为""
			bool数组：默认值为false
		5) 使用数组的步骤 1.声明数组并开辟空间 2.给数组各个元素赋值 3.使用数组
		6) 数组的下标是从0开始的
		7) 数组下标必须在指定范围内使用，否则报panic:数组越界 ，比如var arr [5]int 则有效下标为0-4
		8) go的数组属值类型，在默认情况下是值传递，因此会进行值拷贝。数组间不会相互影响
		9) 如想在其他函数中，去修改原来的数组，可以使用引用传递(指针方式)
		10) 长度是数组类型的一部分，在传递函数参数时，需要考虑数组的长度
	*/

	// 1) 数组是多个相同类型数据的组合，一个数组一旦声明/定义了，其长度是固定的，不能动态变化
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	// arr1[2] = 1.1 报错  数组元素已定int型  1.1是float64类型 类型不匹配
	// arr1[3] = 5 报错  数组一旦声明/定义了，其长度是固定的，不能动态变化
	fmt.Println("arr1=", arr1)

	// 8) go的数组属值类型，在默认情况下是值传递，因此会进行值拷贝。数组间不会相互影响
	arr := [3]int{11, 22, 33}
	test01(arr)
	fmt.Println("arr=", arr)

	// 9) 如想在其他函数中，去修改原来的数组，可以使用引用传递(指针方式)
	test02(&arr)
	fmt.Println("arr=", arr)

	// 10) 长度是数组类型的一部分，在传递函数参数时，需要考虑数组的长度
	var arr2 = [...]int{1, 2, 3} // 长度为3的数组
	test01(arr2)                 // 因为test01函数要求参数为 数组 [3]int  所以test01只能传长度为3的数组 即arr2长度只能为3
	// 不能传 var arr2=[...]int{1,2}或var arr2=[...]int{1,2,3,4}

}

// 8) go的数组属值类型，在默认情况下是值传递，因此会进行值拷贝。数组间不会相互影响
func test01(arr [3]int) {
	arr[0] = 88
}

// 9) 如想在其他函数中，去修改原来的数组，可以使用引用传递(指针方式)
func test02(arr *[3]int) { // 这个arr就不是数组了  是指向数组的指针
	(*arr)[0] = 88 // (*arr) 取到数组的值
}
