package main

import "fmt"

//这个是变量

/**
var 语句用于声明一系列变量。和函数的参数列表一样，类型在最后。

如例中所示，var 语句可以出现在包或函数的层级。
*/

var c, java, python bool

func main() {
	var i int
	fmt.Println(c, java, python, i)

	/**
	变量声明可以包含初始值，每个变量对应一个。

	如果提供了初始值，则类型可以省略；变量会从初始值中推断出类型
	*/

	var a = 10
	var b, c, d = false, false, 1
	fmt.Println(a, b, c, d)

	/**
	在函数中，短赋值语句 := 可在隐式确定类型的 var 声明中使用。

	函数外的每个语句都 必须 以关键字开始（var、func 等），因此 := 结构不能在函数外使用。
	*/

	aa := 12
	var bb string = "no"

	cc, dd, ee := true, 133, "nodff"

	fmt.Println(aa, bb, cc, dd, ee)

	/**
	bool

	string

	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte // uint8 的别名

	rune // int32 的别名
	     // 表示一个 Unicode 码位

	float32 float64

	complex64 complex128
	*/

	/**
	没有明确初始化的变量声明会被赋予对应类型的 零值。

	零值是：

	数值类型为 0，
	布尔类型为 false，
	字符串为 ""（空字符串）。
	*/

	/**
	表达式 T(v) 将值 v 转换为类型 T。

	一些数值类型的转换：

	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	或者，更加简短的形式：

	i := 42
	f := float64(i)
	u := uint(f)
	与 C 不同的是，Go 在不同类型的项之间赋值时需要显式转换。试着移除例子中的 float64 或 uint 的类型转换，看看会发生什么。
	*/

	/**
	常量的声明与变量类似，只不过使用 const 关键字。

	常量可以是字符、字符串、布尔值或数值。

	常量不能用 := 语法声明。
	*/
	const name = "lisi"

	/**
	数值常量是高精度的 值。

	一个未指定类型的常量由上下文来决定其类型。

	再试着一下输出 needInt(Big) 吧。

	（int 类型可以存储最大 64 位的整数，根据平台不同有时会更小。）
	*/
}
