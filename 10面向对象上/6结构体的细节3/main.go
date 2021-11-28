package main

import (
	"encoding/json"
	"fmt"
)

/**
结构体的注意事项和使用细节:
	1) 结构体的所有字段在内存中是连续的
	2) 结构体是用户单独定义的类型,和其他结构体类型进行转换时需要有完全相同的字段(名字、个数和类型)
	3) 结构体进行type重新定义(相当于取别名),go认为是新的数据类型,但是相互间可以强转
	4) struct的每个字段上,可以写上一个tag,该tag可以通过反射机制获取,常见的使用场景就是序列化和反序列化
*/

// 定义Monster结构体
type Monster struct {
	Name  string `json:"name"` // 给struct每个字段写上tag,该tag可以通过反射机制获取
	Age   int    `json:"age"`  // `json:"age"` 就是struct字段Age的tag
	Skill string `json:"skill"`
}

func main() {
	// 4) struct的每个字段上,可以写上一个tag,该tag可以通过反射机制获取,常见的使用场景就是序列化和反序列化

	// 1,创建一个Monster结构体变量
	monster := Monster{"tom", 50, "eat"}
	// 2,将monster结构体变量序列化为json格式字符串
	jsonStr, err := json.Marshal(monster) // 返回 1.序列化后的json格式字符串(byte类型 需要用string()转换)  2.err
	if err != nil {
		fmt.Println("json 处理错误", err)
	}

	fmt.Println("jsonStr=", string(jsonStr))
	// jsonStr= {"Name":"tom","Age":50,"Skill":"eat"}
	// 问题:因为序列化为json格式字符串通常是返回给前端,而很多程序员不习惯字段首字母为大写
	// 分析:如果把结构体中字段首字母改为小写的话,这些字段在其他包将无法访问,也就无法在json包中序列化为json格式字符串了
	// 解决:在struct的每个字段上,可以写上一个tag,该tag可以通过反射机制获取  如上 Name  string `json:"name"`
	// 实现了结构体变量序列化后字段首字母通过反射机制都变成了小写,利于前端使用
	// jsonStr= {"name":"tom","age":50,"skill":"eat"}

}
