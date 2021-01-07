// 要修改原始结构体实例包含的值，必须使用指针。
// 指针是指向内存地址的引用，因此使用它的操作的不是结构体的副本，而是其本身。

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
	// 将指向a的指针（而不是a本身）赋值给b
	b := &a
	// 修改b，将修改分配给a的内存，因为a和b指向相同的内存
	b.Ice = false
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%+v\n", *b)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", &a)
}

// {Lemonade false}
// &{Lemonade false}
// {Name:Lemonade Ice:false}
// {Name:Lemonade Ice:false}
// 0xc00000c080
// 0xc00000c080
