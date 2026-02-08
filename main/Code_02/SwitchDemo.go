package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
*
这个是进行switch的练习
switch也可以在前面写赋值
Go 的 switch 语句类似于 C、C++、Java、JavaScript 和 PHP 中的，
不过 Go 只会运行选定的 case，而非之后所有的 case。
在效果上，Go 的做法相当于这些语言中为每个 case 后面自动添加了所需的 break 语句。在 Go 中，
除非以 fallthrough 语句结束，否则分支会自动终止。 Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不限于整数。
*/
func main() {
	fmt.Print("Go 运行的系统环境：")
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("周六是哪天？")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("今天。")
	case today + 1:
		fmt.Println("明天。")
	case today + 2:
		fmt.Println("后天。")
	default:
		fmt.Println("很多天后。")
	}
}
