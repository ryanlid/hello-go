// 理解变量和零值
// 在 Go 语言中声明变量时，如果没有给它指定值，则变量将为默认值，这种默认值被称为零值。

package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// 0 0 false ""
