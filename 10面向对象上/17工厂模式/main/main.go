package main

import (
	"fmt"
	"go_base/10面向对象上/17工厂模式/model"
)


/**
工厂模式
说明：
	go的结构体没有构造函数，通常可以使用工厂模式来解决这个问题
*/
/**
需求背景：
	1.如果 package model包有个Student结构体，如何在main包中使用该结构体
    2.如果 package model包有个student结构体，如何在main包中使用该结构体(结构体首字母小写)

*/
/**
需求解决:跨包创建结构体实例(变量)的方法
	1.如果model包的Student结构体首字母大写，在main包中引入model包后，直接使用，没有问题
	2.如果model包的student结构体首字母小写, 在main包中引入model包后，不能直接使用，可以用工厂模式解决

*/
func main() {
	// 跨包创建Student结构体实例

	// model包的Student结构体首字母大写 引入包后 直接使用
	var stu = model.Student{
		Name:  "tom",
		Score: 78.09,
	}

	fmt.Println("stu=", stu)

	// model包的student结构体首字母小写 用工厂模式 调用返回student结构体实例的函数
	var stu2 = model.NewStudent("tom", 88.8) // 因为在model包NewStudent函数返回的是指向结构体变量数据空间的指针变量
	// 所以这里stu2是指向结构体变量数据空间的指针变量
	fmt.Println("stu2=", stu2)
	fmt.Println("stu2=", *stu2)
	fmt.Println("stu2.Score=", stu2.Score)

	/**
	思考:如果model包的student的结构体的字段Score改成score(首字母小写),我们还能正常访问吗，如果让其可以访问
		Score改成score无法在其他包直接访问score，可以在model包中给结构体student绑定一个方法用来返回score的值
	*/
	fmt.Println("stu2.score=",stu2.GetScore())

}
