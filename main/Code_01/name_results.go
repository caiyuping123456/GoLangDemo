package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum / 10
	y = sum - (x * 10)
	return
}

func main() {
	fmt.Println("这个是裸返回值")
	fmt.Println(split(17))
}
