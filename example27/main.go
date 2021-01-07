// 使用结构体
// 一系列具有指定数据类型的数据字段，可以通过单一变量引用一系列相关的值，可在单个变量中存储众多类型不同的数据字段。
package main

import "fmt"

// 关键字 type 指定一种新类型，新类型名称为 Movie
// 大括号内，使用名称和类型指定一系列数据字段，请注意，此时没有给数据字段复制，将结构体视为模板
type Movie struct {
	Name   string
	Rating float32
}

// 使用简短变量赋值声明并初始化变量m，给数据字段指定的值，为相应的数据类型。
func main() {
	m := Movie{
		Name:   "Citizen Kane",
		Rating: 10,
	}

	// 使用点表示法，访问数据字段并打印到控制台
	fmt.Println(m.Name, m.Rating)
}

// Citizen Kane 10
