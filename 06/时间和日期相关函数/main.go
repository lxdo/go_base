package main

import (
	"fmt"
	"strconv"
	"time"
)

/**
时间和日期相关函数

1) 时间和日期相关的函数，需要导入time包
2) time.Time 类型，用于表示时间
*/
func main() {
	// time.Time 类型，用于表示时间
	// 1.获取当前本地时间
	now := time.Now() // 类型为time.Time 类型
	fmt.Printf("now= %v   now type= %T \n", now, now)

	// 获取到其他的日期信息
	// 2.通过now可以获取到年月日、时分秒数据
	fmt.Printf("年= %v \n", now.Year())
	fmt.Printf("月= %v \n", now.Month())      // 返回英文月份
	fmt.Printf("月= %v \n", int(now.Month())) // int转为中文月份
	fmt.Printf("日= %v \n", now.Day())
	fmt.Printf("时= %v \n", now.Hour())
	fmt.Printf("分= %v \n", now.Minute())
	fmt.Printf("秒= %v \n", now.Second())

	// 格式化日期时间 返回当前时间
	// 方式1：就是使用 Printf或者Sprintf

	// 格式化输出 Printf根据format参数生成格式化的字符串并写入标准输出
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	// 格式化转换 Sprintf根据format参数生成格式化的字符串并返回该字符串 需要用变量接收
	date := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("date=%v \n", date)

	// 方式2:使用now.Format("2006/01/02 15:04:05")
	// 返回当前时间 但"2006/01/02 15:04:05" 这个字符串的各个数字值是固定的，各个数字都代表其含义
	// 2006->年 01->月 02->日  15->时 04->分 05->秒 中间的拼接符号可以自定义
	/**
	注意："2006/01/02 15:04:05" 这个字符串的各个数字是固定的，必须是这样写
	        年 月  日  时 分 秒
		 "2006/01/02 15:04:05" 这个字符串各个数字可以自由的组合，这样可以按程序需求来返回时间和日期
	*/

	fmt.Println(now.Format("2006/01/02 15:04:05")) // 返回当前时间 2021/03/01 22:03:38
	fmt.Println(now.Format("2006-01-02"))          // 2021-03-01
	fmt.Println(now.Format("15:04:05"))            // 22:11:34
	fmt.Println(now.Format("01"))                  // 03

	/**
	时间的常量
		const (
			Nanosecond Duration = 1 // 纳秒
			Microsecond         = 1000*Nanosecond // 微妙
			Millisecond 		= 1000*Microsecond // 毫秒
			Second				= 1000*Millisecond // 秒
			Minute				= 60*Second // 分钟
			Hour				= 60*Minute // 小时
		)
	时间常量的作用：在程序中可用于获取指定时间单位的时间，比如想得到100毫秒	100*time.Millisecond

	休眠
		使用到time中func Sleep (d Duration)方法
		案例：time.Sleep(100*time.Millisecond) // 休眠100毫秒
	*/

	// 结合Sleep方法使用时间常量
	i := 0
	for {
		i++
		fmt.Println(i)
		// 休眠
		time.Sleep(time.Second) // 休眠1秒

		// 休眠0.1秒
		// time.Sleep(time.Second * 0.1) // 报错 时间常量不能用除法
		time.Sleep(time.Millisecond * 100) // 时间常量可以用乘法

		if i == 2 {
			break
		}
	}

	/**
	获取当前unix时间戳和unix纳秒时间戳
		作用：可以用来获取随机数(纳秒时间戳获取随机粒度更细)
	*/

	fmt.Println("unix时间戳=", now.Unix())
	fmt.Println("unix纳秒时间戳=", now.UnixNano())

	// 统计函数test执行时间
	// 在执行test前，先获取到当前unix时间戳
	start := time.Now().Unix()
	fmt.Println("start=", start)
	test()
	// 在执行test后，再获取到当前unix时间戳
	end := time.Now().Unix()
	fmt.Println("end=", end)
	fmt.Printf("执行test函数耗费时间为 %v 秒 \n", end-start)
}

// 练习 编写一段代码来统计函数test执行时间

func test() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i) // 把整数转成字符串
	}
}
