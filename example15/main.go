// 使用具名返回值

// 具名返回值让函数能够在返回前将值赋值给具名变量，有助于提升函数的可读性。
// 使用具名返回值，在函数签名的返回值

package main

import "fmt"

func sayHi() (x, y string) {
	x = "Hello"
	y = "World"
	return
}

func main() {
	fmt.Println(sayHi())
}

// Hello World
