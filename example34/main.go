// 区分指针引用和值引用
// 以值引用的方式复制结构体
// 将指向结构体的变量赋给另一个变量时，a与b相同，但它是b的副本，而不是指向b的引用，修改b不会影响a，反之亦然。
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
	b := a
	b.Ice = false
	fmt.Printf("%+v\n", b)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
}

// {Name:Lemonade Ice:false}
// {Name:Lemonade Ice:true}
// 0xc00000c080
// 0xc00000c0a0
