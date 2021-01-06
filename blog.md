# Go 语言基本语法

## 检查变量的类型

通过标准库中的 reflect 包，访问变量的底层类型

```go example01
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s string = "string"
	var i int = 10
	var f float32 = 1.2
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(f))
}

// string
// int
// float32

```

## 类型转换

`strconv` 包提供了一整套类型转换方法，可用于转换为字符串或将字符串转换为其他类型

```go example02
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var b bool = true
	fmt.Println(reflect.TypeOf(b))
	var s string = strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(s))
}

// bool
// string
```

## 快捷变量声明

可以在一行内声明多个类型相同的变量，并给它们赋值

```go example03
package main

import "fmt"

func main() {
	var s, t string = "foo", "bar"
	fmt.Println(s)
	fmt.Println(t)
}

// foo
// bar
```

快捷方式声明类型不同的变量

```go example04
package main

import "fmt"

func main() {
	var (
		s string = "foo"
		i int    = 4
	)
	fmt.Println(s)
	fmt.Println(i)

}

// foo
// 4
```

## 理解变量和零值

在 Go 语言中声明变量时，如果没有给它指定值，则变量将为默认值，这种默认值被称为零值。

```go example05
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// 0 0 false ""
```

## 检查变量的值是否为零值

在 Go 语言中，为确定变量是否赋值，不能检测它是否为 nil,而必须检查它是否为默认值

```go example06
package main

import "fmt"

func main() {
	var s string
	if s == "" {
		fmt.Println("s has not been assigned a value and is zero valued")
	}
}

// s has not been assigned a value and is zero valued
```

## 简短变量声明

`:=` 表明 使用的是简短变量声明，不需要使用关键字 var,不用指定变量的类型，将右边的值赋给变量。

不能在函数外使用简短变量声明

```go example06
package main

import "fmt"

func main() {
	s := "Hello World"
	fmt.Println(s)
}

// Hello World
```

## 惯用的变量声明

在函数内使用简短变量声明，在函数外省略类型

> 所有的变量声明方式
>
> ```go
> var s string = "Hello World"
> var s = "Hello World"
> var t string
> t = "Hello World"
> u := "Hello World"
> ```

```go exmaple08
package main

import "fmt"

var s = "Hello World"

func main() {
	i := 42
	fmt.Println(s)
	fmt.Println(i)
}

```

## 使用指针

打印变量在内存中的地址

```go exmaple09
package main

import "fmt"

func main() {
	s := "Hello World"
	fmt.Println(&s)
}

// 可能每次运行输出的值都不一样
// 0xc000010200
```

## 将变量作为值传递

```go exmaple10
package main

import "fmt"

func showMemoryAddress(x int) {
	fmt.Println(&x)
	return
}

func main() {
	i := 1
	fmt.Println(&i)
	showMemoryAddress(i)
}

// 0xc000018098
// 0xc0000180b0

```

## 将变量作为指针传递

```go exmaple11
package main

import "fmt"

func showMemoryAddress(x *int) {
	fmt.Println(x)
	return
}

func main() {
	i := 1
	fmt.Println(&i)
	showMemoryAddress(&i)
}

// 0xc0000b4008
// 0xc0000b4008

```

## 声明常量

在程序中试图修改常量将导致错误

```go exmaple12
package main

import "fmt"

const greeting string = "Hello World"

func main() {
	fmt.Println(greeting)
}
```

## 函数返回多个值

```go exmaple13
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
```

## 定义不定参数函数

要指定不定参数，可使用 3 个点(...)，表示可以接收可变数量的参数

```go exmaple14
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

// The result is 10
```

## 使用具名返回值

具名返回值让函数能够在返回前将值赋值给具名变量，有助于提升函数的可读性。

使用具名返回值，在函数签名的返回值

```go exmaple15
package main

import "fmt"

func sayHi() (x, y string) {
	x = "Hello"
	y = "World"
	return
}

func main() {
	fmt.Println(sayHi())
}

// Hello World
```

## 使用递归函数

不断的调用自己，直到满足特定条件的函数

实现递归，可调用自己的代码，作为终止语句的返回值

```go exmaple16
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
```

## 将函数作为值传递

Go 将函数视为一种类型，可将函数赋值给变量，以后再通过变量来调用它们。

```go exmaple17
package main

import "fmt"

func anotherFunction(f func() string) string {
	return f()
}

func main() {
	fn := func() string {
		return "function called"
	}
	fmt.Println((anotherFunction(fn)))
	// function called
}
```

## 包含 range 子句的 for 语句

```go exmaple18
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

```

## 使用 defer 语句

`defer` 语句可以在函数返回前，执行另一个函数

```go exmaple19
package main

import (
	"fmt"
)

// defer 语句，将在它所在的函数返回前执行
func main() {
	defer fmt.Println("I am run after the function completes")
	fmt.Println("hello world")
}

// hello world
// I am run after the function completes
```

多个 defer 时，按语句出现的相反的顺序执行它们指定的函数

```go exmaple20
package main

import (
	"fmt"
)

// 多个defer时，按语句出现的相反的顺序执行它们指定的函数
func main() {
	defer fmt.Println("I am the first defer statement")
	defer fmt.Println("I am the second defer statement")
	defer fmt.Println("I am the third defer statement")
	fmt.Println("Hello World")
}

/*
输出

Hello World
I am the third defer statement
I am the second defer statement
I am the first defer statement
*/

```
