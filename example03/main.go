// 快捷变量声明
/* 可以在一行内声明多个类型相同的变量，并给它们赋值 */

package main

import "fmt"

func main() {
	var s, t string = "foo", "bar"
	fmt.Println(s)
	fmt.Println(t)
}

// foo
// bar
