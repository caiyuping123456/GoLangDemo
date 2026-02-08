package main

import "fmt"

/*
*
结构体
*/

// 定义结构体
type Demo struct {
	x int
	y int
}

func main() {
	result := Demo{1, 2}
	fmt.Println(result)

	//结构体字段可通过点号 . 来访问。
	//y的值
	fmt.Println(result.y)
	//修改
	result.y = 14
	fmt.Println(result.y)

	//定义指针
	p := &result
	p.y = 100
	fmt.Println(result.y)

	var (
		v1 = Demo{1, 2}  // 创建一个 Vertex 类型的结构体
		v2 = Demo{x: 1}  // Y:0 被隐式地赋予零值
		v3 = Demo{}      // X:0 Y:0
		p1 = &Demo{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
	)
	fmt.Println(v1, p1, v2, v3)
}
