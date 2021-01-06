// 检查变量的类型

// 通过标准库中的 reflect 包，访问变量的底层类型

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s string = "string"
	var i int = 10
	var f float32 = 1.2
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(f))
}

// string
// int
// float32
