// 检查变量的值是否为零值

// 在 Go 语言中，为确定变量是否赋值，不能检测它是否为nil,而必须检查它是否为默认值

package main

import "fmt"

func main() {
	var s string
	if s == "" {
		fmt.Println("s has not been assigned a value and is zero valued")
	}
}

// s has not been assigned a value and is zero valued
