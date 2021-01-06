// 快捷变量声明
/* 快捷方式声明类型不同的变量 */

package main

import "fmt"

func main() {
	var (
		s string = "foo"
		i int    = 4
	)
	fmt.Println(s)
	fmt.Println(i)

}

// foo
// 4
