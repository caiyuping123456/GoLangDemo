package main

import (
	"fmt"
)

func main() {
	fmt.Println("使用for循环")
	fmt.Println("打印100遍：你好")
	s := "你好"

	for i := 1; i < 100; i++ {
		fmt.Println(s)
	}

	/**
	注意，初始化和后置是可以不用的
	*/
	i := 1
	for i < 100 {
		i++
		fmt.Println(s)
	}

	/**
	注意，go中的for是c中的while，所以；可以去掉
	*/

	for i < 100 {
		i++
		fmt.Println(s)
	}

	/**
	无限循环
	*/
	for {
		fmt.Println("这个是无限循环")
	}

}
