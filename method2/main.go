package main

import "fmt"

// type Triangle struct{
// 	width float64
// 	height float64
// }

// Triangle 三角形
type Triangle struct {
	base   float64
	height float64
}

func (t *Triangle) area() float64 {
	return 0.5 * (t.base * t.height)
}

// 修改三角形的底值
func (t Triangle) changeBase(f float64) {
	t.base = f
	return
}

func main() {
	t := Triangle{base: 3, height: 1}
	fmt.Println(t.area())
}
