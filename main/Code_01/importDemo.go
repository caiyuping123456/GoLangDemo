package main

/**
此代码用圆括号将导入的包分成一组，这是“分组”形式的导入语句。

当然你也可以编写多个导入语句，例如：

import "fmt"
import "math"
不过使用分组导入语句要更好。
*/

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("这个是测试包导入", rand.Int()*10)
}
