// 在切片中删除元素

// 从切片中删除元素，也可以使用内置函数 append

package main

import "fmt"

func main() {
	var cheeses = make([]string, 3)
	cheeses[0] = "Alice"
	cheeses[1] = "Bob"
	cheeses[2] = "Carol"
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)

	// 删除索引为2 的元素
	cheeses = append(cheeses[:2], cheeses[2+1:]...)
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
}

// 3
// [Alice Bob Carol]
// 2
// [Alice Bob]
