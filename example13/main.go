// ## 函数返回多个值

package main

import "fmt"

func getPrize() (int, string) {
	i := 2
	s := "goldfish"
	return i, s
}

func main() {
	quantity, prize := getPrize()
	fmt.Printf("You win %v %v \n", quantity, prize)
}

// You win 2 goldfish
