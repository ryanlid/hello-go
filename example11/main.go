// 将变量作为指针传递

package main

import "fmt"

func showMemoryAddress(x *int) {
	fmt.Println(x)
	return
}

func main() {
	i := 1
	fmt.Println(&i)
	showMemoryAddress(&i)
}

// 0xc0000b4008
// 0xc0000b4008
