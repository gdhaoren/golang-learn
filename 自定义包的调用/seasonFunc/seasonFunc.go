// seasonFunc project seasonFunc.go
package seasonFunc_

// import使用的是包所在文件夹的名字。import过后在程序中进行调用的时候，是使用包文件中写好的package name来进行调用，这个名字会在两个.go文件中，一个是name.go 一个是doc.go
// 由于不包含package main 和 mian()因此对于这种包文件进行编译是不会生成可执行文件的 .exe
import "fmt"

func Season(month int) string {
	var result string
	switch month {
	case 1, 2, 12:
		result = fmt.Sprintf("%d是冬季", month)
		return result
	case 3, 4, 5:
		return fmt.Sprintf("%d是春季", month)
	case 6, 7, 8:
		return fmt.Sprintf("%d是夏季", month)
	case 9, 10, 11:
		return fmt.Sprintf("%d是秋季", month)
	default:
		return "未知季节月份"
	}
}
