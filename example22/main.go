// 使用切片

// 在 Go 语言中，使用数组存在一定的局限性。
// 在数组中无法添加元素，在切片可以添加、删除元素，复制切片中的元素

package main

import "fmt"

func main() {
	// 使用内置函数make() 声明一个长度为2的空切片
	// 第一个参数为数据类型，第二个参数为长度
	var cheeses = make([]string, 2)
	cheeses[0] = "Alice"
	cheeses[1] = "Bob"
	cheeses = append(cheeses, "Carol")
	fmt.Println(cheeses[2])
}

// Carol
