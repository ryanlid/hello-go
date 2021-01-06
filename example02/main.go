// 类型转换

/*
strconv 包提供了一整套类型转换方法，可用于转换为字符串或将字符串转换为其他类型
*/

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var b bool = true
	fmt.Println(reflect.TypeOf(b))
	var s string = strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(s))
}

// bool
// string
