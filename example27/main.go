// 接 example17
// 函数append 是一个不定参数函数

package main

import "fmt"

// 在切片末尾添加多个元素

func main() {
	var cheeses = make([]string, 2)
	cheeses[0] = "Alice"
	cheeses[1] = "Bob"
	cheeses = append(cheeses, "Carol", "Reblochon")
	fmt.Println(cheeses)
}

// [Alice Bob Carol Reblochon]
