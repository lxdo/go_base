

// 要求开发一个hello.go 输出"hello,world"
/**
【规范】【执行】

1.go文件的后缀为 .go
2.package main 表示该hello.go文件所在包是main,在go中每个文件都必须归属于一个包
3.import "fmt" 表示：引入一个包，包名fmt,引入该包后，就可以使用fmt包的函数，比如：fmt.Println
4.func main(){

}// func 是关键字，表示一个函数  即我们程序的入口
5.go build 源文件 编译    （正式环境） 执行 ./执行文件  go build -o h（执行文件） hello.go(源文件) 编译自定义执行文件名称
6.go run  源文件 直接执行 （开发环境）

	说明:两种执行流程方式的区别
		1.如果我们先编译生成了可执行文件，那么我们可以将改执行文件拷贝到没有go开发环境的机器上，仍然可以运行
		2.如果我们是直接go run 源代码，那么如果要在另外机器上这么运行，也需要go的开发环境，否则无法运行
		3.在编译时，编译器会将程序运行依赖的库文件包含在可执行文件中，所以，可执行文件变大了很多
*/

package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
