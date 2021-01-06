// 使用数组

// 创建数组：声明一个数组变量，并指定其长度和数据类型
// 声明数组的长度后，就不能给它增加元素

package main

import "fmt"

func main() {
	var cheeses [2]string
	cheeses[0] = "Alice"
	cheeses[1] = "Bob"
	fmt.Println(cheeses[0])
	fmt.Println(cheeses[1])
	fmt.Println(cheeses)
}

// Alice
// Bob
// [Alice Bob]
