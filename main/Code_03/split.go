package main

import "fmt"

//切片
/**
每个数组的大小都是固定的。而切片则为数组元素提供了动态大小的、灵活的视角。 在实践中，切片比数组更常用。

类型 []T 表示一个元素类型为 T 的切片。.

切片通过两个下标来界定，一个下界和一个上界，二者以冒号分隔：

a[low : high]
它会选出一个半闭半开区间，包括第一个元素，但排除最后一个元素。

以下表达式创建了一个切片，它包含 a 中下标从 1 到 3 的元素：

a[1:4]
*/
func main() {
	a := []int{1, 2, 3, 4, 5, 6, 78}

	var result []int = a[2:4]

	fmt.Println(result)

	/**
	切片类似数组的引用
	切片就像数组的引用 切片并不存储任何数据，它只是描述了底层数组中的一段。

	更改切片的元素会修改其底层数组中对应的元素。

	和它共享底层数组的切片都会观测到这些修改。
	*/
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a1 := names[0:2]
	b := names[1:3]
	fmt.Println(a1, b)

	b[0] = "XXX"
	fmt.Println(a1, b)
	fmt.Println(names)

	/**
	切片字面量类似于没有长度的数组字面量。

	这是一个数组字面量：

	[3]bool{true, true, false}
	下面这样则会创建一个和上面相同的数组，然后再构建一个引用了它的切片：

	[]bool{true, true, false}
	*/
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	/**
	在进行切片时，你可以利用它的默认行为来忽略上下界。

	切片下界的默认值为 0，上界则是该切片的长度。

	对于数组

	var a [10]int
	来说，以下切片表达式和它是等价的：

	a[0:10]
	a[:10]
	a[0:]
	a[:]
	*/
	s1 := []int{2, 3, 5, 7, 11, 13}

	s1 = s1[1:4]
	fmt.Println(s1)

	s1 = s1[:2]
	fmt.Println(s)

	s1 = s1[1:]
	fmt.Println(s1)

	/**
	切片拥有 长度 和 容量。

	切片的长度就是它所包含的元素个数。

	切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。

	切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取。

	你可以通过重新切片来扩展一个切片，给它提供足够的容量。 试着修改示例程序中的切片操作，向外扩展它的长度，看看会发生什么。
	*/
	s2 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s2)

	// 截取切片使其长度为 0
	s2 = s2[:0]
	printSlice(s2)

	// 扩展其长度
	s2 = s2[:4]
	printSlice(s2)

	// 舍弃前两个值
	s2 = s2[2:]
	printSlice(s2)

	/**
	nil 切片
	切片的零值是 nil。

	nil 切片的长度和容量为 0 且没有底层数组。
	*/
	var s3 []int
	fmt.Println(s3, len(s3), cap(s3))
	if s3 == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
