package main

import "fmt"

// 数组
/**
类型 [n]T 表示一个数组，它拥有 n 个类型为 T 的值。

表达式

var a [10]int
会将变量 a 声明为拥有 10 个整数的数组。

数组的长度是其类型的一部分，因此数组不能改变大小。 这看起来是个限制，不过没关系，Go 拥有更加方便的使用数组的方式。
*/
func main() {

	//申明数组
	var a [3]int
	a[0] = 1
	a[2] = 3
	fmt.Println(a[0], a[2])
	fmt.Println(a)

	b := []string{"a", "b", "c", "d"}
	fmt.Println(b)

}
