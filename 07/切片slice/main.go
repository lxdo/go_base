package main

import "fmt"

/**
切片 slice
	基本介绍：
		1) 切片的英文是slice
		2) 切片是数组的一个引用,因此切片是引用类型，在进行传递时，遵守引用传递的机制
		3) 切片的使用和数组类似，遍历切片、访问切片的元素和求切片长度 len(slice) 都一样
		4) 切片的长度是可以变化的，因此切片是一个可以动态变化数组
		5) 切片定义的基本语法：
			var 变量名 []类型 比如：var a []int


*/
func main() {
	// 切片基本使用

	// 声明数组
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	// 声明/定义一个切片
	// 1.slice就是切片名
	// 2.intArr[1:3] 表示slice引用到intArr数组的第2个元素(起始下标为1)到第4个元素(不包含第4个元素)(结束下标为3，不包含3)
	// 包前不包后
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	// slice切片的元素
	fmt.Println("slice=", slice)
	// slice切片的元素个数
	fmt.Println("slice len=", len(slice))
	// slice切片的容量 (切片目前可以存放最多元素的个数 容量可以动态变化)
	fmt.Println("slice cap=", cap(slice))
	// 切片是引用类型  slice引用 intArr[1:3](不包含下标3)  slice[0]的地址和intArr[1]的地址相同
	fmt.Printf("intArr[1]地址= %p slice第一个元素地址= %p \n", &intArr[1], &slice[0])
	// intArr[1]地址= 0xc00001a0f8 slice第一个元素地址= 0xc00001a0f8

	// slice是引用类型，改变slice的元素就会改变intArr对应的元素 ，因为他们是同一个地址
	slice[1] = 35
	fmt.Println("intArr=", intArr)
	// slice切片的元素
	fmt.Println("slice=", slice)

	/**
	切片的使用:
		方式1：  定义一个切片，然后让切片去引用一个已经创建好的数组，上面的slice就是引用的intArr数组
				案例：
					var intArr [5]int = [...]int{1, 22, 33, 66, 99}
					slice := intArr[1:3]

		方式2:  通过make来创建切片
				基本语法:var 切片名 []type = make([]type,len,[cap])
				参数说明:[]type就是切片数据类型 ,len:大小, cap:切片容量(可选)
				如果加cap参数 需要 cap >= len

				通过make方式创建切片可以指定切片的大小和容量
				如果没有给切片的各个元素赋值，那么就会使用默认值 [int->0,float->0,string->"",bool->false]
				通过make方式创建的切片对应的数组是由make底层维护，对外不可见，即只能通过slice去访问各个元素
		方式3：	定义一个切片，直接就指定具体数组，使用原理类似make的方式
				var slice3 []string = []string{"tom", "jack", "mary"}

		方式1和方式2的区别(面试题)：
				方式1是直接引用数组，这个数组是事先存在的，程序员是可见的
				方式2是通过make来创建切片，make也会创建一个数组，是由切片在底层进行维护，程序员是看不见的

	*/

	// 方式2: 通过make来创建切片
	var slice2 []int = make([]int, 4, 10)                                     // 参数: 切片数据类型 切片大小 切片容量 (切片容量>=切片大小)
	fmt.Println("slice2=", slice2, " len=", len(slice2), "cap=", cap(slice2)) // 定义未赋值 默认值为0
	slice2[0] = 100                                                           // 访问切片
	slice2[2] = 200
	fmt.Println("slice2=", slice2, " len=", len(slice2), "cap=", cap(slice2))

	// 方式3
	var slice3 []string = []string{"tom", "jack", "mary"}
	fmt.Println("slice3=", slice3, "len=", len(slice3), "cap=", cap(slice3))

	/**
	切片的遍历
		1) for循环常规方式遍历
		2) for range 结构遍历切片
	*/

	// 用方式1定义一个切片
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice4 := arr[1:4] // [20 30 40]
	// 1) for循环常规方式遍历
	for i := 0; i < len(slice4); i++ {
		fmt.Printf("slice4[%v]=%v ", i, slice4[i])
	}

	fmt.Println()
	// 2) for range 方式遍历切片
	for i, v := range slice4 {
		fmt.Printf("i=%v v=%v ", i, v)
	}
	fmt.Println()
	/**
	切片注意事项和细节说明

		1) (方式1定义切片) 切片初始化时 var slice = arr[startIndex:endIndex]
			说明：从arr数组下标为startIndex，取到下标为endIndex的元素（不包含arr[endIndex]）
		2) (方式1定义切片) 切片初始化时，仍然不能越界，范围在 [0-len(arr)]之间 ，但是可以动态增长
			(1) var slice = arr[0:end] 可以简写 var slice = arr[:end]
			(2) var slice = arr[start:len(arr)]可以简写：var slice = arr[start:]
			(3) var slice = arr[0:len(arr)] 可以简写 var slice = arr[:]
		3) cap是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素
		4) 切片定义完后，还不能使用，因为本身是一个空的，需要让其引用到一个数组，或者make一个空间供切片来使用
		5) 切片可以继续切片
		6) 用append内置函数，可以对切片进行动态追加
	*/

	// var slice = arr[0:len(arr)] 可以简写 var slice = arr[:]
	slice5 := arr[:]
	fmt.Println("arr=", arr)
	fmt.Println("slice5=", slice5)

	// 5) 切片可以继续切片
	// 切片slice6引用切片slice5 startIndex=1 endIndex=2 (startIndex、endIndex是slice5的下标) 包含 slice5[1]不包含slice5[2]
	slice6 := slice5[1:2]
	fmt.Println("slice6=", slice6)

	// arr、slice5、slice6 中相对应的元素都指向同一数据空间，一个当中元素变化，另外两个相对应的元素也会变化
	slice6[0] = 1000
	fmt.Println("arr=", arr)
	fmt.Println("slice5=", slice5)
	fmt.Println("slice6=", slice6)

	// 6) 用append内置函数，可以对切片进行动态追加
	var slice7 []int = []int{100, 200, 300}
	// 通过append直接给slice7追加具体的元素 (元素数据类型要和切片数据类型一致)
	slice8 := append(slice7, 400, 500, 600) // 追加完slice7本身没有变化，需要重新接收(给slice7)或赋值(给slice8)
	fmt.Println("slice7=", slice7)
	fmt.Println("slice8=", slice8)
	// append是对切片进行动态追加，所以一般用原来切片名(slice7)接收追加后的切片
	slice7 = append(slice7, 400, 500, 600)
	fmt.Println("slice7=", slice7)
	// 在切片的基础上追加切片 第二个参数必须是切片加"..."
	slice7 = append(slice7, slice7...) // 将切片slice7(可以是自己或其他切片)追加给切片slice7

	fmt.Println("slice7=", slice7)

	/**
	切片append操作底层原理分析
		1) 切片append操作的本质就是对数组扩容(是对切片指向的数组扩容)
		2) go底层会创建一下新的数组newArr(安装扩容后大小)
		3) 将slice原来包含的元素拷贝到新的数组newArr
		4) slice 重新引用到newArr(原变量接收)
		5) 注意newArr是在底层来维护的，程序员不可见

		// 相当于创建新的数组，拷贝旧值，追加新值，让切片指向新的数组
	*/

	/**
	切片的拷贝操作
		切片使用copy内置函数完成拷贝
		copy(slice2,slice1)  // slice2和slice1数据类型都要是切片 copy(拷贝切片,被拷贝切片)
		拷贝后两个切片数据空间是独立的，互相不影响，修改一个数据，不会影响另一个数据
		当len(slice1被拷贝切片) > len(slice2拷贝切片) 只会拷贝 slice2(拷贝切片)长度能装下的元素
	*/
	var slice9 []int = []int{1, 2, 3, 4, 5}
	var slice10 = make([]int, 10)
	fmt.Println("slice10=", slice10)
	copy(slice10, slice9) // 直接拷贝即可 copy(拷贝切片,被拷贝切片)
	fmt.Println("slice10=", slice10)
	fmt.Println("slice9=", slice9)
}
