// 包含range子句的for 语句

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	for i, n := range numbers {
		fmt.Println("index: ", i)
		fmt.Println("value: ", n)
		fmt.Println("-----------")
	}
}

// index:  0
// value:  1
// -----------
// index:  1
// value:  2
// -----------
// index:  2
// value:  3
// -----------
// index:  3
// value:  4
// -----------
