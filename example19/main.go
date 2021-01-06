// 使用 defer 语句
// defer 语句可以在函数返回前，执行另一个函数
package main

import (
	"fmt"
)

// defer 语句，将在它所在的函数返回前执行
func main() {
	defer fmt.Println("I am run after the function completes")
	fmt.Println("hello world")
}

// hello world
// I am run after the function completes
