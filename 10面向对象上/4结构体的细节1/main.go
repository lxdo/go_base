package main

import "fmt"

/**
结构体的注意事项和使用细节:
	1) 结构体的所有字段在内存中是连续的
	2) 结构体是用户单独定义的类型,和其他类型进行转换时需要有完全相同的字段(名字、个数和类型)
*/

// 定义结构体
type Point struct {
	x int
	y int
}

// 定义结构体
type Rect struct {
	leftUp, rightDown Point // leftUp、rightDown 字段数据类型 为Point结构体
}
type Rect2 struct {
	leftUp, rightDown *Point // leftUp、rightDown 字段数据类型 为指向Point结构体数据空间的指针变量
}

func main() {
	// 1) 结构体的所有字段在内存中是连续的

	// 定义Rect结构体变量
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	// r1有四个int,在内存中是连续分布的 (int占8字节)
	fmt.Printf("r1.leftUp.x的地址=%p r1.leftUp.y的地址=%p r1.rightDown.x的地址=%p r1.rightDown.y的地址=%p \n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)
	// r1.leftUp.x的地址=0xc0000b8000 r1.leftUp.y的地址=0xc0000b8008 r1.rightDown.x的地址=0xc0000b8010 r1.rightDown.y的地址=0xc0000b8018
	// r1结构体变量所有字段在内存中是连续的(int占8字节)

	// r2有两个*Point类型(指向Point结构体数据空间的指针变量),这两个*Point类型的本身地址在内存中是连续的,
	// 他们指向的Point结构体数据空间的地址不一定是连续的
	// 定义Rect2结构体变量
	r2 := Rect2{&Point{1, 2}, &Point{3, 4}}
	fmt.Printf("r2.leftUp的本身地址=%p r2.rightDown的本身地址=%p \n",
		&r2.leftUp, &r2.rightDown)
	// r2.leftUp的本身地址=0xc00008e1e0 r2.rightDown的本身地址=0xc00008e1e8
	// r2的两个字段(指针变量)本身的地址是连续的 (int占8字节)
	fmt.Printf("r2.leftUp的指向地址=%p r2.rightDown的指向地址=%p",
		r2.leftUp, r2.rightDown)
	// r2.leftUp的指向地址=0xc000016080 r2.rightDown的指向地址=0xc000016090
	// 他们指向的Point结构体数据空间的地址不一定是连续的



}
