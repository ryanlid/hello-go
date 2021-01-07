// 比较结构体

// 判断两个结构体是否相等，使用 ==
// 判断两个结构体是否不相等，使用 !=
// 不能对两个类型不同的结构体进行比较，否则将导致编译错误

package main

import "fmt"

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
	if a == b {
		fmt.Println("a and b are the same")
	}
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
}

// a and b are the same
// {Name:Lemonade Ice:true}
// {Name:Lemonade Ice:true}
