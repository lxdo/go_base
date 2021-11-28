package main

import "fmt"


/**
map切片
	基本介绍:
		切片的数据类型如果是map，则我们称为 map切片(slice of map)
		map切片([]map[key数据类型]value数据类型)可以存放多个map
		这样使用 map切片里的map个数就可以动态变化了

		map切片可以简单理解为:由多个map组成的切片

*/
func main() {
	// 声明一个map切片
	var students []map[string]string // map切片中元素都是map(map[string]string)  map个数可以动态变化
	// make切片 声明的切片make后才可使用
	students = make([]map[string]string, 2) // make切片(切片声明后需要make) students切片 len为2 即可以存放2个map

	// 给map切片中赋值map
	if students[0] == nil {
		students[0] = make(map[string]string, 2) // make map(用make声明map) map切片中元素都是map ,所以需要先make map
		students[0]["name"] = "tom"
		students[0]["age"] = "18"
	}
	fmt.Println("students=", students) // students= [map[age:18 name:tom] map[]]
	// map切片中第二个map未赋值 所以为map[]

	// 给map切片中赋值map
	if students[1] == nil {
		students[1] = make(map[string]string, 2) // make map(用make声明map) map切片中元素都是map ,所以需要先make map
		students[1]["name"] = "jack"
		students[1]["age"] = "20"
	}
	fmt.Println("students=", students) // students= [map[age:18 name:tom] map[age:20 name:jack]]

	// 因为students切片在make时设定len为2,如果需要动态增加students切片的map个数,不能像上面直接给map切片赋值map,不然会报错
	// 可以用append函数动态给map切片增加map

	// 定义1个map
	student := map[string]string{
		"name": "mary",
		"age":  "25",
	}
	// 通过append把 student(map) 加入 students(切片)
	students = append(students, student)
	fmt.Println("students=", students)
	// students= [map[age:18 name:tom] map[age:20 name:jack] map[age:25 name:mary]]

}
