// 创建方法集

package main

import (
	"fmt"
	"math"
)

// Sphere 球
// 声明结构体
type Sphere struct {
	Radius float64
}

// SurfaceArea 计算球体表面积
// 声明将结构体 Sphere 作为接收者的方法
func (s *Sphere) SurfaceArea() float64 {
	return float64(4) * math.Pi * (s.Radius * s.Radius)
}

// Volume 计算球体体积
func (s *Sphere) Volume() float64 {
	radiusCubed := s.Radius * s.Radius * s.Radius
	return (float64(4) / float64(3)) * math.Pi * radiusCubed
}

func main() {
	s := Sphere{
		Radius: 5,
	}
	fmt.Println(s.SurfaceArea())
	fmt.Println(s.Volume())
}
