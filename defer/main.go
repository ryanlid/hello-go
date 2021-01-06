package main

import (
	"fmt"
)

// // defer 语句，将在它所在的函数返回前执行
// func main() {
// 	defer fmt.Println("I am run after the function completes")
// 	fmt.Println("hello world")
// }

// 多个defer时，按语句出现的相反的顺序执行它们指定的函数
func main() {
	defer fmt.Println("I am the first defer statement")
	defer fmt.Println("I am the second defer statement")
	defer fmt.Println("I am the third defer statement")
	fmt.Println("Hello World")
}

/*
输出

Hello World
I am the third defer statement
I am the second defer statement
I am the first defer statement
*/
