// 使用 关键字 new 来创建结构体实例
package main

import "fmt"

type Movie struct {
	Name   string
	Rating float32
}

func main() {
	m := new(Movie)
	// var m Movie
	m.Name = "Alice"
	m.Rating = 0.998
	// fmt.Printf("%+v\n", m)
	fmt.Printf("%+v\n", m)
}

// TODO: 实际打印会多出一个&，待确认原因
// &{Name:Alice Rating:0.998}
