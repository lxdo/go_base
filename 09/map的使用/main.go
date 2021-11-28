package main

import "fmt"

/**
map的使用

map使用的方式
	方式1)
		// 声明 ,这时map=nil
		var a map[string][string]
		// make分配数据空间
		a = make(map[string][string],10)  // map数据类型 map长度(能存放几对key-value)

	方式2)
		// 声明,就直接make
		var a = make(map[string][string]) // 不用约定map长度,可以赋值多个key-value
	方式3)
		// 声明,直接赋值
		var a map[string]string = map[string]string{"name":"tom"}
*/
func main() {
	// map使用的方式
	// 方式1 声明后进行make
	var a map[string]string
	a = make(map[string]string, 10)
	a["aa"] = "aa"
	fmt.Println("a=", a)

	// 方式2 声明,就直接make
	// 不用约定map长度,可以赋值多个key-value
	var b = make(map[string]string)
	b["c1"] = "北京"
	b["c2"] = "上海"
	b["c3"] = "南京"
	fmt.Println("b=", b)

	// 方式3 声明,直接赋值,可以赋值多个key->value
	c := map[string]string{
		"what":  "what",
		"why":   "why",
		"where": "where",
		"how":   "how",
	}
	c["who"] = "who" // 可以继续赋值
	fmt.Println("c=", c)

	// 案例:用二维map存放2个学生信息,每个学生有name和sex信息
	// 思路: map[string]map[string]string

	studentMap := make(map[string]map[string]string) // make 二维map 不设长度
	// 这个make不能少
	studentMap["stu01"] = make(map[string]string, 3) // make 二维map中一维map,每个一维map设长度3 (存放3对key-value)
	studentMap["stu01"]["name"] = "tom"              // 给一维map赋值
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "beijing"

	// 这个make不能少
	studentMap["stu02"] = make(map[string]string, 3) // make 二维map中一维map,每个一维map设长度3 (存放3对key-value)
	studentMap["stu02"]["name"] = "mary"             // 给一维map赋值
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "nanjing"

	fmt.Println("studentMap=", studentMap)

	/**
	map的增删改查操作
		map增加和更新:
			map["key"]=value // 如果key还没有,就是增加 如果key存在就是修改(覆盖key-value)
		map删除:
			delete(map,"key") , delete是一个内置函数,如果key存在,就删除该key-value,如果key不存在,不操作,但是也不会报错
			细节说明:
				1) 如果我们要删除map的所有key,没有一个专门的方法一次删除,可以遍历一下key,逐个删除key
				2) 或者 map = make(...), make一个新的,让原来的成为垃圾,被gc回收 // map的名称不变  (推荐使用)
		map查找:
			val,res:=cities["n1"]
			1) 如果map cities中存在key n1,那么 val就是key n1对应的value 。
			2) 如果map cities中存在key n1,那么res就会返回true,否则返回false
			3) val,res为自定义变量

	*/
	// map增加和更新
	cities := make(map[string]string)
	cities["n1"] = "北京" // key没有,就是增加
	cities["n2"] = "南京" // key没有,就是增加
	cities["n3"] = "深圳" // key没有,就是增加
	cities["n1"] = "上海" // key存在,就是修改 后面key-value 覆盖前面key-value
	fmt.Println("cities=", cities)

	// map删除
	delete(cities, "n1") // key存在,就删除
	fmt.Println("cities=", cities)
	delete(cities, "n10") // key不存在,不操作,不报错

	// 一次性删除map的所有key
	// 1) 方法一 遍历所有key,逐一删除
	// 2) 方法二 直接make一个新的空间 map的名称不变 (推荐)
	cities = make(map[string]string)
	fmt.Println("cities=", cities)

	cities["n1"] = "苏州"
	// map查找
	val, res := cities["n1"] // 自定义val,res
	// 如果map cities存在key n1  val为n1对应的value res为true 否则res为false
	if res {
		fmt.Println("map cities key n1 = value", val)
	} else {
		fmt.Println("not found")
	}

}
