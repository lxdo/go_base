package main

import "fmt"

/**
map练习
	1) 使用 map[string]map[string]string 的map类型 (二维map:map的key对应的value也是map类型)
	2) key:表示用户名,是唯一的,不可以重复
	3) 如果某个用户名存在,就将其密码改为"888888" ,如果不存在就增加这个用户信息 (包括昵称nickname和密码pwd)
	4) 编写一个函数 modifyUser(users map[string]map[string]string,name string) 完成上述功能
*/
func main() {
	// 定义map
	users := make(map[string]map[string]string, 10) // make二维map len:10 可以放10对key-value(map)
	modifyUser(users, "tom")
	modifyUser(users, "mary")
	fmt.Println("users=", users) // map是引用类型 函数内改变 函数外也改变
	modifyUser(users, "tom")
	fmt.Println("users=", users) // map是引用类型 函数内改变 函数外也改变
}

func modifyUser(users map[string]map[string]string, name string) () {
	// 判断map users中是否有key 为name(传进函数的name)
	if users[name] != nil { // 说明users中有key为name(传进函数的name)
		users[name]["pwd"] = "888888"
	} else { // 说明users中没有key为name(传进函数的name)
		// 二维map中的一维map也需要make make二维map中的一维map
		// 一维map len为2
		users[name] = make(map[string]string, 2)
		users[name]["nickname"] = name
		users[name]["pwd"] = "999999"
	}
}
