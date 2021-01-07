// 检查结构体的类型

package main

import (
	"fmt"
	"reflect"
)

type Drink struct {
	Name string
	Ice  bool
}

func main() {
	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	b := Drink{
		Name: "Lemonade",
		Ice:  true,
	}

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
}

// main.Drink
// main.Drink
