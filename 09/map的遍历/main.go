package main

import "fmt"

/**
map的遍历
	map的遍历只能用for range结构遍历

*/

func main() {
	// 一维map遍历
	cities := make(map[string]string)
	cities["c1"] = "北京"
	cities["c2"] = "天津"
	cities["c3"] = "上海"

	for key, value := range cities {
		fmt.Printf("key=%v value=%v ", key, value)
	}

	fmt.Println()

	// 二维map遍历
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

	// 双层for range遍历二维map
	for key, value := range studentMap {
		fmt.Printf("key=%v", key)
		for k, v := range value {
			fmt.Printf(" k=%v v=%v", k, v)
		}
		fmt.Println()
	}

	/**
	map的长度
		len(map) // 获取map里有多少对key-value
	*/
	// 一维map
	fmt.Println("cities len =", len(cities)) //  map cities 长度为3 有3对key-value
	// 二维map 会获取到有多少对一维map(key-map)
	fmt.Println("studentMap len =",len(studentMap)) // map studentMap 长度为2 有2对key-map(value)



}
