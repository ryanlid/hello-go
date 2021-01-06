// 使用递归函数
package main

import "fmt"

func feedMe(portion int, eaten int) int {
	eaten = portion + eaten
	if eaten >= 5 {
		fmt.Printf("I'm full! I 've eaten %d\n", eaten)
		return eaten
	}
	fmt.Printf("I'm still hungry! I've eaten %d\n", eaten)
	return feedMe(portion, eaten)
}
func main() {
	feedMe(1, 0)
}

// I'm still hungry! I've eaten 1
// I'm still hungry! I've eaten 2
// I'm still hungry! I've eaten 3
// I'm still hungry! I've eaten 4
// I'm full! I 've eaten 5
