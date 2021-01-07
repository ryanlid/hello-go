// 使用映射

// 映射是通过键，来访问的无序元素编组

package main

import "fmt"

func main() {
	// 通过 make 创建一个空映射，其键的类型是字符串，其值的类型是整数
	var players = make(map[string]int)

	// 在映射中动态添加元素
	players["cook"] = 32
	players["bairstow"] = 27
	players["stokes"] = 26

	fmt.Println(players["cook"])
	fmt.Println(players["bairstow"])
	fmt.Println(players)

	// 从映射中删除元素
	delete(players, "cook")
	fmt.Println(players)
}

// 32
// 27
// map[bairstow:27 cook:32 stokes:26]
// map[bairstow:27 stokes:26]
