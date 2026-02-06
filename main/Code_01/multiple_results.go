package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	s, s2 := swap("你好", "golang")
	fmt.Println("进行string交换：" + s + "  " + s2)
}
