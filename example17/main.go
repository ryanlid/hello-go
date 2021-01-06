// 将函数作为值传递
// Go 将函数视为一种类型，可将函数赋值给变量，以后再通过变量来调用它们。

package main

import "fmt"

func anotherFunction(f func() string) string {
	return f()
}

func main() {
	fn := func() string {
		return "function called"
	}
	fmt.Println((anotherFunction(fn)))
	// function called
}
