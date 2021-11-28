package main

import "fmt"

/**
面向对象编程-继承
	基本介绍:
		1) 继承可以解决代码复用
		2) 当多个结构体存在相同的属性(字段)和方法时，可以从这些结构体中抽象出结构体(比如Student),在该结构体中定义这些相同的属性和方法
		3) 其他的结构体不需要重新定义这些属性(字段)和方法，只需嵌套一个Student匿名结构体即可
		4) 也就是说:在go中，如果一个struct嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的字段和方法，从而实现了继承特性

	基本语法:
		type Goods struct {
			Name string
			Price int
		}

		type Book struct {
			Goods // 这里就是嵌套匿名结构体 Goods (继承)
			Writer string
		}
	继承优点:
		1) 代码的复用性提高了
		2) 代码的扩展性和维护性提高了


*/
// 案例:使用嵌套匿名结构体的方式来实现继承特性 ，编写一个学生考试系统

// 基础结构体 Student
type Student struct {
	// 共有字段
	Name  string
	Age   int
	Score float64
}

// 共有方法

// 显示成绩
func (stu *Student) ShowInfo() {
	fmt.Printf("Name=%v Age=%v Score=%v \n", stu.Name, stu.Age, stu.Score)
}

// 评分
func (stu *Student) SetScore(score float64) {
	// 业务判断分数的合法性

	stu.Score = score
}

// 给*Student结构体增加一个方法，那么Pupil和Graduate都可以使用该方法
func (stu *Student) GetSum(n1, n2 int) int {
	return n1 + n2
}

// 小学生
type Pupil struct {
	Student // 嵌入了Student匿名结构体 (继承)
}

// 这是Pupil结构体特有方法 方法中内容不一样
func (p *Pupil) testing() {
	fmt.Println("Pupil testing")
}

// 大学生
type Graduate struct {
	Student // 嵌入了Student匿名结构体 (继承)
}

// 这是Graduate结构体特有方法 方法中内容不一样
func (g *Graduate) testing() {
	fmt.Println("Graduate testing")
}

func main() {
	// 当我们对结构体嵌入了匿名结构体使用方法会发生变化
	pupil := &Pupil{}           // pupil为指向结构体变量数据空间的指针变量
	pupil.Student.Name = "tome" // 结构体变量.匿名结构体.匿名结构体字段
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(70) // 结构体变量.匿名结构体.匿名结构体方法
	pupil.Student.ShowInfo()

	fmt.Println(pupil.Student.GetSum(1,2))
	// 大学生
	// 当我们对结构体嵌入了匿名结构体使用方法会发生变化
	graduate := &Graduate{}        // graduate为指向结构体变量数据空间的指针变量
	graduate.Student.Name = "mary" // 结构体变量.匿名结构体.匿名结构体字段
	graduate.Student.Age = 22
	graduate.testing()
	graduate.Student.SetScore(90)
	graduate.Student.ShowInfo()

}
