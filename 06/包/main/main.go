package main

/**
引入包方法：
		包的绝对路径：/Users/lixiaodong/m2odata/www/go/src/go_code/06/包/utils
	1) 先配置项目GOPATH路径   指到src上层目录  /Users/lixiaodong/m2odata/www/go
	2)
		import (
			"包路径"   从src 下层目录写起  (默认从src下定位)   go_code/06/包/utils
		)
*/
import (
	"fmt"
	"go_base/06/包/utils"
	//ut "go_code/06/包/utils"  // 给包起别名 在调用时用 ut.Cal(别名.函数) ut.Num (别名.变量)  原来的包名就不能使用了
)

/**
包

介绍：
	在实际的开发中，往往需要在不同的文件中，去调用其他文件的定义的函数，比如main.go中，去使用utils.go文件中的函数

包的原理：
	包的本质实际上就是创建不同的文件夹，来存放程序文件

包的演示：包 文件夹

基本概念：
	go的每一个文件都是属于一个包的，也就是说go是以包的形式来管理文件和项目目录结构的

包的作用：
	1)	区分相同名字的函数、变量等标识符
	2)  当程序文件很多时，可以很好的管理项目
	3) 	控制函数、变量等访问范围，即作用域

包的相关说明：
	打包基本语法：package 包名  例 package utils
	引入包的基本语法：import "包的路径"

包使用的快速入门：
		文件夹：包 文件夹
		实现： go相互调用函数
		操作：将func Cal定义到文件utils.go，将utils.go放到一个包(utils)中
		当其他文件需要使用到utils.go的方法时，可以import该包，就可以使用了

包的使用细节：
	1) 在给一个文件打包时，该包对应一个文件夹，比如utils文件夹对应的包名就是utils
		文件的包名通常和文件所在的文件夹名一致，一般为小写字母

	2) 当一个文件要使用其他包函数或变量时，需要先引入对应的包
		(1) 引入方式1: import "包名"  // 引入单个包
		(2) 引入方式2:  引入多个包
						import (
							"包名"
							"包名"
						)
	3) package 指令在文件第一行 ，然后是import指令

	4) 在import包时，路径从$GOPATH的src下开始，不用带src，编译器会自动从src下开始引入

	5) 为了让其他包的文件，可以访问到本包的函数，则该函数名的首字母需要大写，类似其他语言的public，这样才能跨包访问
		比如 utils.go的Cel函数

	6) 在访问其他包函数、变量时，其语法是 包名.函数名  包名.变量名

	7) 如果包名比较长，go支持给包取别名，取别名后，原来的包名就不能使用了

	8) 在同一包下，不能有相同的函数名(也不能有相同的全局变量名)，否则报重复定义
		(1) 同一包下，同一文件下函数不能重名  即utils文件夹下 utils.go文件不能有同名函数
		(2)	同一包下, 不同文件下函数不能重名  即utils文件夹下 多个文件不能有同名函数

	9) 如果要编译成一个可执行程序文件，就需要将这个包声明为main(main包只能有一个)，即package main,这个就是一个语法规范
	如果是写一个库，包名可以自定义
		详见：图片 将包编译成可执行文件
			(1) 在GOPATH目录(src上层目录)下  执行编译命令: go build  src(不需要带)下main包路径
				默认在GOPATH目录下生成一个可执行文件
			(2) 指定可执行文件名称: 在GOPATH目录(src上层目录)下 执行编译命令: go build -o bin/my  src(不需要带)下main包路径
				会在GOPATH目录下生成 bin(目录)/my(可执行文件)
*/

func main() {

	var n1 float64 = 1.3
	var n2 float64 = 1.4
	var op byte = '+'

	// 引用utils包中Cal函数 包.函数名
	res := utils.Cal(n1, n2, op)
	fmt.Println("res =", res)

	// 引用utils包中变量 (变量名首字母要大写)  格式：包.变量名
	fmt.Println("utils包中变量Num", utils.Num)
}
