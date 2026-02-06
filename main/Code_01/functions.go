package main

import "fmt"

/**
函数可接受零个或多个参数。

在本例中，add 接受两个 int 类型的参数。

注意类型在变量名的 后面。
*/

// 这是一个加函数
func add(x int, y int) int {
	return x + y
}

/**
当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。
被简写为

x, y int

*/
// 这是一个乘函数
func matifiy(x, y int) int {
	return x * y
}

func main() {
	//测试加函数
	fmt.Println("测试加函数：", add(3, 5))
}
