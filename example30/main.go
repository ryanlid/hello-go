// 嵌套结构体
// 为建立较复杂的数据结构，在一个结构体中嵌套另一个结构体的方式很有用

package main

import "fmt"

type Superhero struct {
	Name    string
	Age     int
	Address Address
}
type Address struct {
	Number int
	Street string
	City   string
}

func main() {
	m := Superhero{
		Name: "Batman",
		Age:  32,
		Address: Address{
			Number: 1007,
			Street: "Mountain",
			City:   "Gotham",
		},
	}
	fmt.Printf("%+v\n", m)
}

// {Name:Batman Age:32 Address:{Number:1007 Street:Mountain City:Gotham}}
