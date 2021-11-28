package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
字符串中常用的系统函数

	1) 统计字符串的长度，按字节 len(str)  // 内建函数
	2) 字符串遍历，同时处理有中文的问题  r:=[]rune(str)
	3) 字符串转整数 n,err:=strconv.Atoi("12")
	4) 整数转字符串 str=strconv.ltoa(12345)
	5) 字符串转 []byte(byte切片):var bytes=[]byte("hello go")
	6) []byte(byte切片)转字符串：str=string([]byte{97,98,99})
	7) 把10进制转2，8，16进制:str=strconv.FormatInt(123,2) 返回对应的字符串
	8) 查找子串是否在指定的字符串中：strings.Contains("seafood","food") // 如果有返回true
	9) 统计一个字符串中有几个指定的子串 ： strings.Count("chinese","e") //2
	10) 不区分大小写的字符串比较(==是区分字母大小写的) ：strings.EqualFold("abc","ABC") // true
	11) 返回子串在字符串第一次出现的index值，如果没有返回-1 ：strings.Index("NLT_abc","abc") //4
	12) 返回子串在字符串最后一次出现的index，如没有返回-1 ：strings.LastIndex("go goland","go")
	13) 将指定的子串替换成另外一个子串：strings.Replace("go go hello","php",n) n可以指定你希望替换几个，如果n=-1表示全部替换
		不会改变原字符串，需要用变量接收替换后的新字符串
	14) 按照指定的某个字符为分割标识，将一个字符串拆分成字符串数组：strings.Split("hello,world,ok",",")
		不会改变原字符串
	15) 将字符串的字母进行大小写的转换：
			转成小写 strings.ToLower("Go") // go
			转成大写 strings.ToUpper("Go") // GO
		不会改变原字符串
	16) 将字符串左右两边的空格去掉：strings.TrimSpace(" tn a lone go   ")
		不会改变原字符串
	17) 将字符串左右两边指定的字符去掉 ： strings.Trim("! hello! "," !") // "hello"  将左右两边!和" "(空格)都去掉
		不会改变原字符串
	18) 将字符串左边指定的字符去掉： strings.TrimLeft("! hello! "," !") // "hello! " 将左边!和" "(空格)去掉
		不会改变原字符串
	19) 将字符串右边指定的字符去掉： strings.TrimRight("! hello! "," !") // "! hello" 将右边!和" "(空格)去掉
		不会改变原字符串
	20) 判断字符串是否以指定的字符串开头：strings.HasPrefix("ftp://192.168.10.1","ftp") // true
	21) 判断字符串是否以指定的字符串结束：strings.HasSuffix("NLT_abc.jpg","abc") // false
*/
func main() {
	// 1) 统计字符串的长度，按字节 len(str)  // 内建函数
	str := "hello"
	str2 := "hello中"                    // go编码统一为utf8  字母和数字(ASCII的字符)占一个字节，汉字占3个字节
	fmt.Println("str len=", len(str))   // 5
	fmt.Println("str2 len=", len(str2)) // 8 按字节返回

	// 2) 字符串遍历，同时处理有中文的问题  r:=[]rune(str)
	r := []rune(str2)             // 可以解决带中文的字符串遍历乱码
	for i := 0; i < len(r); i++ { // len(str)按字节返回 汉字占3个字节会乱码 需要用r:=[]rune(str)转换
		fmt.Printf(" %c ", r[i])
	}

	fmt.Println()
	// 3) 字符串转整数 n,err:=strconv.Atoi("12")

	// 转换成功的返回
	n1, err1 := strconv.Atoi("123")
	fmt.Println("n1=", n1, "err1=", err1)
	// 可以这样判断转换是否成功
	if err1 != nil {
		fmt.Println("转换错误", err1)
	} else {
		fmt.Println("转换结果是", n1)
	}
	// 转换失败的返回
	n2, err2 := strconv.Atoi("hello")
	fmt.Println("n2=", n2, "err2=", err2)
	// 可以这样判断转换是否成功
	if err2 != nil {
		fmt.Println("转换错误", err2)
	} else {
		fmt.Println("转换结果是", n2)
	}

	// 4) 整数转字符串 str=strconv.ltoa(12345)
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v ,str type %T \n", str, str)

	// 5) 字符串转 []byte(byte切片):var bytes=[]byte("hello go") // 应用场景：把字符串写入到二进制文件中
	var bytes = []byte("hello go") // 把字符串的每个字符(包括空格)转换为对应的ASCII码
	// 把字符串每个字符(包括空格)转换为对应的ASCII码
	fmt.Printf("bytes=%v \n", bytes) // bytes=[104 101 108 108 111 32 103 111]

	// 6) []byte(byte切片)转字符串：str=string([]byte{97,98,99})
	str = string([]byte{97, 98, 99}) // 把每个ASCII码转换成对应的字符，并且组成一个字符串返回
	fmt.Println("str=", str)         // str= abc

	// 7) 把10进制(int64)转2，8，16进制:str=strconv.FormatInt(123,2) 返回对应的字符串
	str = strconv.FormatInt(123, 2)  // 转成二进制
	fmt.Println("123对应的二进制 =", str)  // 123对应的二进制 = 1111011
	str = strconv.FormatInt(123, 16) // 转成十六进制
	fmt.Println("123对应的十六进制 =", str) // 123对应的十六进制 = 7b

	// 8) 查找子串是否在指定的字符串中：strings.Contains("seafood","food")  // 如果有返回true
	b := strings.Contains("seafood", "food") // 有返回true
	fmt.Printf("b= %v \n", b)                // b= true
	b = strings.Contains("seafood", "mary")
	fmt.Printf("b= %v \n", b) // b= false 没有返回false

	// 9) 统计一个字符串中有几个指定的子串 ： strings.Count("chinese","e") //2
	count := strings.Count("chinese", "e")
	fmt.Printf("chinese中有 %v 个e \n", count) // 2

	// 10) 不区分大小写的字符串比较(==是区分字母大小写的) ：strings.EqualFold("abc","ABC") // true
	b = strings.EqualFold("abc", "ABC") // 不区分字母大小写
	fmt.Println("b=", b)                //b= true
	// ==区分字母大小写
	fmt.Println("abc" == "ABC")

	// 	11) 返回子串在字符串第一次出现的index值，如果没有返回-1 ：strings.Index("NLT_abc","abc") //4
	index := strings.Index("NLT_abc", "abc") // index从0开始
	fmt.Println("index=", index)             // index= 4 子串首次出现的index值
	index = strings.Index("NLT_abc", "he")   // 没有出现返回-1
	fmt.Println("index=", index)             // index= -1

	// 12) 返回子串在字符串最后一次出现的index，如没有返回-1 ：strings.LastIndex("go goland","go")
	index = strings.LastIndex("go golang", "go") // index 从0开始 空格也算
	fmt.Println("index=", index)                 // index= 3

	// 13) 将指定的子串替换成另外一个子串：strings.Replace("go go hello","php",n) n可以指定你希望替换几个，如果n=-1表示全部替换
	// 不会改变原字符串
	str = strings.Replace("go go hello", "go", "php", 1)  // 只会替换第一个
	fmt.Println("str=", str)                              // str= php go hello
	str = strings.Replace("go go hello", "go", "php", -1) // 全部替换
	fmt.Println("str=", str)                              // str= php php hello
	s := "go go hello"
	str = strings.Replace(s, "go", "php", -1)
	// 不会改变原来字符串 ，会将替换后的新字符串返回，需要用变量接受新字符串
	fmt.Println("s=", s)     // s= go go hello
	fmt.Println("str=", str) // str= php php hello

	// 14) 按照指定的某个字符为分割标识，将一个字符串拆分成字符串数组：strings.Split("hello,world,ok",",")
	// 不会改变原字符串
	arr := strings.Split("hello,world,ok", ",")    // 用" ," 将字符串分割成字符串数组
	fmt.Printf("arr=%v  arr type=%T \n", arr, arr) // arr=[hello world ok]  arr type=[]string

	/**
	15) 将字符串的字母进行大小写的转换：
			转成小写 strings.ToLower("Go") // go
			转成大写 strings.ToUpper("Go") // GO

		不会改变原字符串
	*/
	// 全部改为小写
	s = "goLang Hello"
	str = strings.ToLower(s)
	fmt.Println("str=", str) // str= golang hello
	// 全部改为大写
	str = strings.ToUpper(s)
	fmt.Println("str=", str) // str= GOLANG HELLO
	fmt.Println("s=", s)     // s= goLang Hello 不会改变原字符串

	// 16) 将字符串左右两边的空格去掉：strings.TrimSpace(" tn a lone go   ")
	// 不会改变原字符串
	s = " tn a lone go   "
	str = strings.TrimSpace(s)
	fmt.Printf("str=%q \n", str) // str="tn a lone go"    将字符串左右两边的空格去掉
	fmt.Printf("s=%q \n", s)     //  s=" tn a lone go   "   不会改变原字符串

	// 17) 将字符串左右两边指定的字符去掉 ： strings.Trim("! hello! "," !") // "hello"  将左右两边!和" "(空格)都去掉
	// 不会改变原字符串
	s = "! hello! "
	str = strings.Trim(s, " !")  // 将字符串s左右两边" "(空格)和!都去掉 中间不会去掉
	fmt.Printf("str=%q \n", str) // str="hello"
	fmt.Printf("s=%q \n", s)     //  s="! hello! " 不会改变原字符串

	// 18) 将字符串左边指定的字符去掉： strings.TrimLeft("! hello! "," !") //  将左边!和" "(空格)去掉
	// 不会改变原字符串
	str = strings.TrimLeft(s, " !")
	fmt.Printf("str=%q \n", str) // str="hello! "

	// 19) 将字符串右边指定的字符去掉： strings.TrimRight("! hello! "," !") // 将右边!和" "(空格)去掉
	// 不会改变原字符串
	str = strings.TrimRight(s, " !")
	fmt.Printf("str=%q \n", str) // str="! hello"

	// 20) 判断字符串是否以指定的字符串开头：strings.HasPrefix("ftp://192.168.10.1","ftp") // true
	b = strings.HasPrefix("ftp://192.168.10.1", "ftp")
	fmt.Println("b=", b) // b= true

	// 21) 判断字符串是否以指定的字符串结束：strings.HasSuffix("NLT_abc.jpg","abc") // false
	b = strings.HasSuffix("NLT_abc.jpg", "abc")
	fmt.Println("b=", b) // b= false
}
