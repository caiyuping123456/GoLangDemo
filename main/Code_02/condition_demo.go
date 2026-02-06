package main

import (
	"fmt"
	"math"
)

func main() {
	/**
	使用if条件,这个小括号就是可以去掉，{}不可以去掉
	*/
	i := 1
	if i < 100 {
		fmt.Println("哈哈哈哈")
	}

	/**
	和 for 一样，if 语句可以在条件表达式前执行一个简短语句。
	该语句声明的变量作用域仅在 if 之内。
	就是把括号里面的提前了
	*/
	if v := math.Pow(10, 2); v < 100 {
		fmt.Println(v)
	}

	/**
	使用else
	*/
	a := 100
	if a++; a < 100 {
		fmt.Println("你好")
	} else {
		fmt.Println("哈哈哈")
	}
}
