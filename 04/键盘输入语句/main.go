package main

import "fmt"

/**
介绍：
	在编程中，需要接收用户输入的数据，就可以使用键盘输入语句来获取
*/

/**
步骤：
	导入fmt包
	调用fmt包的fmt.Scanln() 或者 fmt.Scanf()

*/

func main() {
	// 案例：要求可以从控制台接收用户信息 (姓名，年龄，薪水,是否通过考试)

	// 方式1 fmt.Scanln() 一条一条输入
	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Println("请输入姓名 ")
	// 当程序执行到  fmt.Scanln(&name),程序会停止在这里，等待用户输入，并回车
	fmt.Scanln(&name)

	fmt.Println("请输入年龄 ")
	// 当程序执行到  fmt.Scanln(&age),程序会停止在这里，等待用户输入，并回车
	fmt.Scanln(&age)

	fmt.Println("请输入薪水 ")
	// 当程序执行到  fmt.Scanln(&sal),程序会停止在这里，等待用户输入，并回车
	fmt.Scanln(&sal)

	fmt.Println("请输入是否通过考试 ")
	fmt.Scanln(&isPass)

	fmt.Printf("名字是 %v  年龄是 %v  薪水是 %v  是否通过考试 %v \n", name, age, sal, isPass)

	// 方式2 fmt.Scanf 可以按指定的格式输入
	fmt.Println("请输入你的姓名,年龄,薪水,是否通过考试,使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)

	fmt.Printf("名字是 %v  年龄是 %v  薪水是 %v  是否通过考试 %v \n", name, age, sal, isPass)
}
