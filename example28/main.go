package main

import "fmt"

type Movie struct {
	Name   string
	Rating float32
}

func main() {
	var m Movie
	// 创建结构体时，如果没有初始化，则 Go 将把每个数据字段设置为相应数据类型的零值
	fmt.Printf("%+v\n", m)
	m.Name = "Alice"
	m.Rating = 0.998
	fmt.Printf("%+v\n", m)
}

// {Name: Rating:0}
// {Name:Alice Rating:0.998}
