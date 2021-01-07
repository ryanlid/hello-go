// 复制切片中的元素

// 要复制切片的全部或部分元素，使用内置函数 copy
// 复制前，必须先声明一个类型相同的切片
// 函数 copy 在新切片中创建元素的副本，修改一个切片的元素不会影响另一个切片

package main

import "fmt"

func main() {
	var cheeses = make([]string, 3)
	cheeses[0] = "Alice"
	cheeses[1] = "Bob"
	var smellyCheeses = make([]string, 2)
	copy(smellyCheeses, cheeses)
	fmt.Println(smellyCheeses)
}

// [Alice Bob]
