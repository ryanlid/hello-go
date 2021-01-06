// 将变量作为值传递

package main

import "fmt"

func showMemoryAddress(x int) {
	fmt.Println(&x)
	return
}

func main() {
	i := 1
	fmt.Println(&i)
	showMemoryAddress(i)
}

// 0xc000018098
// 0xc0000180b0
