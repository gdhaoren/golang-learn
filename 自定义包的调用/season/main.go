// season project main.go
package main

// import的是文件夹名字，使用的是文件中定义的package name 二者不一定要同名，但是应当同名以免出现理解混乱，降低程序可读性
// import 的是seasonFunc 而调用使用的是 seasonFunc_.Season
// seasonFunc 是文件夹名字，seasonFunc_则是定义在.go文件中的package seasonFunc_

import (
	"fmt"
	"seasonFunc"
	"strconv"
	//"time"
)

/*
// 格式化字符串的函数后面都会有个f
// 注意Sprintf不能用于直接输出到标准输出显示，他只能用来返回格式化好的字符串,Sprint不能用来格式化字符串，但是可以用来连接字符串
// Printf可以同时格式化字符串并输出到显示器上,但是他没有自动换行的能力，需要自行添加换行符号
// Println只能直接输出结果不能格式化字符串,输出后会换行
func main() {
	var monthStr string
	fmt.Scanln(&monthStr)
	month, _ := strconv.Atoi(monthStr)
	switch {
	case month >= 3 && month <= 5:
		fmt.Printf("%d月是春季\n", month)
	case month >= 6 && month <= 8:
		fmt.Printf("%d月是夏季\n", month)
	case month >= 9 && month <= 11:
		fmt.Printf("%d月是秋季\n", month)
	case month <= 2 || month == 12:
		fmt.Printf("%d月是冬季\n", month)
	}
	//time.Sleep(5 * time.Second)
}
*/
func main() {
	var monthStr string
	fmt.Scanln(&monthStr)
	month, _ := strconv.Atoi(monthStr)
	fmt.Println(seasonFunc_.Season(month))
}
