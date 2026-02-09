package main

import "fmt"

/**
map 映射将键映射到值。

映射的零值为 nil 。nil 映射既没有键，也不能添加键。

make 函数会返回给定类型的映射，并将其初始化备用。
*/

/**
map是 Go 的键值对集合（类似 Python 的字典、Java 的 HashMap）；
string是键的类型（这里用字符串作为地名）；
Vertex是值的类型（这里用结构体存储该地名对应的经纬度）；
*/

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	//m = make(map[string]Vertex)
	//m["Bell Labs"] = Vertex{
	//	40.68433, -74.39967,
	//}
	//fmt.Println(m["Bell Labs"])

	//m := map[string]Vertex{
	//	"Bell Labs": Vertex{40.68433, -74.39967}, // 键名"Bell Labs"是必须的
	//	// 简化写法（值是结构体时，可省略Vertex{}外层），但键名仍不能少
	//	"Google": {37.42202, -122.08408},
	//}
	//
	//fmt.Println(m["Bell Labs"])

	/**
	在映射 m 中插入或修改元素：

	m[key] = elem
	获取元素：

	elem = m[key]
	删除元素：

	delete(m, key)
	通过双赋值检测某个键是否存在：

	elem, ok = m[key]
	若 key 在 m 中，ok 为 true ；否则，ok 为 false。

	若 key 不在映射中，则 elem 是该映射元素类型的零值。

	注：若 elem 或 ok 还未声明，你可以使用短变量声明：

	elem, ok := m[key]
	*/
	m := make(map[string]int)

	m["答案"] = 42
	fmt.Println("值：", m["答案"])

	m["答案"] = 48
	fmt.Println("值：", m["答案"])

	delete(m, "答案")
	fmt.Println("值：", m["答案"])

	v, ok := m["答案"]
	fmt.Println("值：", v, "是否存在？", ok)
}
