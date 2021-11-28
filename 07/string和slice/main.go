package main

import "fmt"

/**
string和slice
	1) string底层是一个byte数组，因此string也可以进行切片处理
	2) string和切片在内存的形式 (见 string和slice内存分析图)
	3) string是不可变的，也就说不能通过str[0]="Z" 方式来修改字符串
	4) 如果需要修改字符串，可以 先将string转换成[]byte或者[]rune  然后再修改  修改完重新转成string
		string-> []byte/[]rune->修改->string

*/

func main() {
	// 1) string底层是一个byte数组，因此string也可以进行切片处理
	str := "hello@world"
	// 使用切片获取到world
	slice := str[6:] // 从str下标为6(str[6])元素引用到str最后
	fmt.Println("slice=", slice)
	// 	3) string是不可变的，也就说不能通过str[0]='H' 方式来修改字符串
	// str[0] = 'H' 报错 string是不可变的,不能通过str[0]='H' 方式修改字符串

	// 4) 如果需要修改字符串，可以 先将string转换成[]byte或者[]rune  然后再修改  修改完重新转成string

	// string->[]byte
	byteSlice := []byte(str)
	// 修改
	byteSlice[0] = 'Z' // 这里如果是汉字就会报错 如： byteSlice[0] = '北'
	// 再转成string
	str = string(byteSlice)
	fmt.Println("str=", str)

	// 我们修改字符串时，把string转换成[]byte后，可以处理英文和数字，但是不能处理中文
	// 原因：[]byte是按字节来处理，而一个汉字是3个字节，因此就会出现乱码
	// 解决方法：在修改字符串时，把string转换成 []rune ,因为[]rune是按字符处理，兼容汉字

	RuneSlice := []rune(str) // string->[]rune
	RuneSlice[0] = '北'       // 修改
	str = string(RuneSlice)  // 再转成string
	fmt.Println("str=", str)

	/**
	为了兼容中文，在修改字符串时，最好都把string转换成[]rune修改 修改完再转回字符串
	*/
}
