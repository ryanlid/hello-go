// 定义不定参数函数
// 要指定不定参数，可使用3个点(...)，表示可以接收可变数量的参数
package main

import "fmt"

// 表示函数可以接受任意数量的 int 参数
func sumNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total

}

func main() {
	result := sumNumbers(1, 2, 3, 4)
	fmt.Printf("The result is %v\n", result)
}
