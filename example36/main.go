// 使用方法

// 方法类似于函数
// 在关键字func 后面添加了另一个参数部分，用于接受单个参数

package main

import (
	"fmt"
	"strconv"
)

type Movie struct {
	Name   string
	Rating float64
}

// 函数 summary 和结构体 Movie 相互依赖，但它们之间没有直接关系，
// 如果不能访问结构体 Movie 定义，就无法声明函数 summary

func (m *Movie) summary() string {
	// 将 float64 类型转换为字符串类型
	r := strconv.FormatFloat(m.Rating, 'f', 1, 64)
	return m.Name + "," + r
}

func main() {
	m := Movie{
		Name:   "Spiderman",
		Rating: 3.2,
	}
	fmt.Println(m.summary())
}

// Spiderman,3.2
