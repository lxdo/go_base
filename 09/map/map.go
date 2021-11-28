package main

import "fmt"

/**
map
	map介绍:
		map是key-value数据结构,又称为字段或者关联数组
	map的声明:
		var 变量名 map[key type]value type  // var map变量名 map[key数据类型]值数据类型
	key可以是什么数据类型:
		go中 ,map的key可以是很多种数据类型,比如bool,数字,string,指针,channel,还可以是只包含前面几个类型的 接口,结构体,数组
		但通常为int、string类型(重点)
		注意:map的key的数据类型不可以是slice,map,function,因为这几个没法用 == 来判断
	value可以是什么数据类型:
		map的value数据类型和key基本一样
		但通常为数字(整数、浮点数),string,map,struct
	map声明示例:
		var a map[string]string // a map变量名 [key数据类型]值数据类型
		var a map[string]int
		var a map[int]string

		var a map[string]map[string]string
		二维map a->map变量名 [string]->[key数据类型] map[string]string->值数据类型 值还是map数据类型

	注意:声明map是不会分配内存的，初始化需要make，分配内存后才能赋值和使用
         map在使用前一定要make
	map使用细节:
		1) map是引用类型,遵守引用类型传递的机制,在一个函数接收map,修改后,会直接修改原来的map
		2) map的容量达到后,再向map增加元素,会自动扩容,并不会发生panic,也就是说map能动态的增长键值对(key-value)
		3) map的value也经常使用struct类型,更适合管理复杂的数据 [比map的value是一个map(二维map)更好]
*/

func main() {
	// map的声明与定义
	var a map[string]string
	// 声明后map是空的(nil map) 没有数据空间
	fmt.Println("a=", a) // a= map[]
	// 没有make直接赋值会报错
	// a["aa"] = "aa" // 报错 没有make map没有数据空间 是 nil map 无法赋值
	// 在使用map前，需要先make，make的作用就是给map分配数据空间
	a = make(map[string]string, 10) // make(map数据类型,map长度)  10:map可以放10对key-value元素
	a["aa"] = "aa"                  // map 赋值
	a["bb"] = "bb"
	a["aa"] = "cc"       // 当key重复时 后赋值的key->value会把先赋值的key->value覆盖掉
	a["cc"] = "bb"       // 当value重复时，没有影响
	fmt.Println("a=", a) // a= map[aa:cc bb:bb]

	// map的key->value是无序的

	// map使用细节
	// 1) map是引用类型,遵守引用类型传递的机制,在一个函数接收map,修改后,会直接修改原来的map
	m := make(map[int]int, 3)
	m[1] = 90
	m[2] = 88
	m[10] = 20
	fmt.Println("m=", m)
	// 函数内修改map 函数外也改变
	modify(m)
	fmt.Println("m=", m)

	// 2) map的容量达到后,再向map增加元素,会自动扩容,并不会发生panic,也就是说map能动态的增长键值对(key-value)
	// 在上面声明map(m)时,将map(m) len设为3 上面已经赋值3对key-value 在这里继续赋值不会报错,map会自动扩容,动态增加键值对
	m[5] = 50
	fmt.Println("m=", m)

	// 3) map的value也经常使用struct类型,更适合管理复杂的数据 [比map的value是一个map(二维map)更好]
	// map的key为学生的学号,是唯一的
	// map的value为结构体,包含学生的名字,年龄,地址
	// 声明一个map
	students := make(map[string]Stu, 10) // map key->string value->Stu(结构体)
	// 定义并赋值两个结构体
	stu1 := Stu{"tom", 18, "北京"}
	stu2 := Stu{"mary", 20, "丽江"}
	// 把结构体放入map
	students["n1"] = stu1
	students["n2"] = stu2
	fmt.Println("students=", students)

	// 遍历各个学生信息
	for k, v := range students { // v 为结构体
		fmt.Printf("k=%v 学生的名字=%v 学生的年龄=%v 学生的地址=%v \n", k, v.Name, v.Age, v.Address)
		// v.Name 结构体中取值 结构体.属性名

	}

}

// 定义一个学生结构体
type Stu struct {
	Name    string
	Age     int
	Address string
}

// 定义一个函数
func modify(m map[int]int) {
	// 在函数内对传入map进行修改
	m[10] = 900
}
