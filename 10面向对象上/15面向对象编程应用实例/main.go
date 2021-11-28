package main

import "fmt"

/**
面向对象编程应用实例

步骤
1) 声明(定义)结构体 ,确定结构体
2) 编写结构体的字段
3) 编写结构体的方法
*/

// 1) 声明(定义)结构体 ,确定结构体
type Student struct {
	// 2) 编写结构体的字段
	name  string
	age   int
	score float64
}

// 3) 编写结构体的方法
func (stu *Student) say() string {
	info := fmt.Sprintf("stu的信息 name=%v age=%v score=%v",
		stu.name, stu.age, stu.score) // 格式化字符串
	return info
}

func main() {
	stu := Student{"tom", 18, 90.5}
	fmt.Println(stu.say())

	var box Box
	box.len = 1.1
	box.width = 1.2
	box.height = 1.3

	volumn := box.volumn()
	fmt.Printf("volumn =  %.2f", volumn)
}

// 定义box结构体
type Box struct {
	len    float64
	width  float64
	height float64
}

// 定义求box体积方法
func (box *Box) volumn() float64 {
	return box.len * box.width * box.height
}
