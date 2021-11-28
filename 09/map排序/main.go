package main

import (
	"fmt"
	"sort"
)

/**
map排序
	基本介绍:
		1) go中没有一个专门的方法针对map的key进行排序
		2) go中的map默认是无序的,注意也不是按照添加的顺序存放的,每次遍历,得到的输出可能不一样
		3) go中map的排序,是先将key进行排序,然后根据key遍历输出value
*/
func main() {
	// 2) go中的map默认是无序的,注意也不是按照添加的顺序存放的,每次遍历,得到的输出可能不一样
	map1 := make(map[int]int)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	fmt.Println("map1=", map1)
	// 3) go中map的排序,是先将key进行排序,然后根据key遍历输出value
	// 思路:1.先将map的key放入到切片 2.对存放map key的切片进行排序 3.遍历切片,按照key输出value
	// 1.先将map的key放入到切片
	var keys []int           // 定义 int切片
	for k, _ := range map1 { // 遍历map
		keys = append(keys, k) // 将map的key放入到切片
	}
	fmt.Println("keys=", keys)
	// 2.对存放map key的切片进行排序
	sort.Ints(keys) // sort.Ints(a []int) 参数为int型切片 将参数(切片a)排序为递增顺序
	fmt.Println("keys=", keys)
	//  3.遍历切片keys,按照key输出value
	for _, v := range keys {
		fmt.Printf("map[%v]=%v ", v, map1[v])
	}
}
