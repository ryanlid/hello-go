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

```go example08
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

```go example09
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

```go example10
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

```go example11
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

```go example12
package main

import "fmt"

const greeting string = "Hello World"

func main() {
	fmt.Println(greeting)
}
```

## 函数返回多个值

```go example13
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

```go example14
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

```go example15
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

```go example16
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

```go example17
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

```go example18
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

```go example19
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

```go example20
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

## 使用数组

声明一个数组变量，需要指定其长度和数据类型

声明数组的长度后，就不能修改它的长度

```go example21
package main

import "fmt"

func main() {
	var cheeses [2]string
	cheeses[0] = "Alice"
	cheeses[1] = "Bob"
	fmt.Println(cheeses[0])
	fmt.Println(cheeses[1])
	fmt.Println(cheeses)
}

// Alice
// Bob
// [Alice Bob]
```

## 使用切片

在 Go 语言中，使用数组存在一定的局限性。

在数组中无法添加元素，在切片可以添加、删除元素，复制切片中的元素

```go example22
package main

import "fmt"

func main() {
	// 使用内置函数make() 声明一个长度为 2 的空切片
	// 第一个参数为数据类型，第二个参数为长度
	var cheeses = make([]string, 2)
	cheeses[0] = "Alice"
	cheeses[1] = "Bob"
	cheeses = append(cheeses, "Carol")
	fmt.Println(cheeses[2])
}

// Carol
```

## 在切片末尾添加多个元素

函数 append 是一个不定参数函数，通过 append 可以在切片末尾添加很多值

```go example23
package main

import "fmt"

func main() {
	var cheeses = make([]string, 2)
	cheeses[0] = "Alice"
	cheeses[1] = "Bob"
	cheeses = append(cheeses, "Carol", "Reblochon")
	fmt.Println(cheeses)
}

// [Alice Bob Carol Reblochon]
```

## 在切片中删除元素

从切片中删除元素，也可以使用内置函数 append

```go example24
package main

import "fmt"

func main() {
	var cheeses = make([]string, 3)
	cheeses[0] = "Alice"
	cheeses[1] = "Bob"
	cheeses[2] = "Carol"
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)

	// 删除索引为2 的元素
	cheeses = append(cheeses[:2], cheeses[2+1:]...)
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
}

// 3
// [Alice Bob Carol]
// 2
// [Alice Bob]
```

## 复制切片中的元素

要复制切片的全部或部分元素，使用内置函数 copy

复制前，必须先声明一个类型相同的切片

函数 copy 在新切片中创建元素的副本，修改一个切片的元素不会影响另一个切片

```go example25
package main

import "fmt"

func main() {
	var cheeses = make([]string, 3)
	cheeses[0] = "Alice"
	cheeses[1] = "Bob"
	var smellyCheeses = make([]string, 2)
	copy(smellyCheeses, cheeses)
	fmt.Println(smellyCheeses)
}

// [Alice Bob]
```

## 使用映射

映射是通过键，来访问的无序元素编组

```go example26
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
```

## 使用结构体

结构体是一系列具有指定数据类型的数据字段，可以通过单一变量引用一系列相关的值，可在单个变量中存储众多类型不同的数据字段。

```go example27
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
```

声明一个类型为结构体的变量

```go example28
package main

import "fmt"

type Movie struct {
	Name   string
	Rating float32
}

func main() {
	var m Movie
	// 创建结构体时，如果没有初始化，则 Go 将把每个数据字段设置为相应数据类型的零值
	fmt.Printf("%+v\n", m)
	m.Name = "Alice"
	m.Rating = 0.998
	fmt.Printf("%+v\n", m)
}

// {Name: Rating:0}
// {Name:Alice Rating:0.998}
```

使用 关键字 new 来创建结构体实例

```go example29
package main

import "fmt"

type Movie struct {
	Name   string
	Rating float32
}

func main() {
	m := new(Movie)
	// var m Movie
	m.Name = "Alice"
	m.Rating = 0.998
	// fmt.Printf("%+v\n", m)
	fmt.Printf("%+v\n", m)
}
// TODO: 实际打印会多出一个&，待确认原因
// &{Name:Alice Rating:0.998}
```
## 嵌套结构体

为建立较复杂的数据结构，在一个结构体中嵌套另一个结构体的方式很有用

```go example30
package main

import "fmt"

type Superhero struct {
	Name    string
	Age     int
	Address Address
}
type Address struct {
	Number int
	Street string
	City   string
}

func main() {
	m := Superhero{
		Name: "Batman",
		Age:  32,
		Address: Address{
			Number: 1007,
			Street: "Mountain",
			City:   "Gotham",
		},
	}
	fmt.Printf("%+v\n", m)
}

// {Name:Batman Age:32 Address:{Number:1007 Street:Mountain City:Gotham}}
```

## 自定义结构体数据字段的默认值

默认情况下，Go 给数据字段指定相应数据类型的零值

Go 语言中没有提供自定义默认值的内置方法，但可使用构造函数来实现这个目标。

通过构造函数创建结构体将没有指定值的数据字段设置默认值。

```go example31
package main

import "fmt"

type Alarm struct {
	Time  string
	Sound string
}

func NewAlarm(time string) Alarm {
	a := Alarm{
		Time:  time,
		Sound: "Klaxon",
	}
	return a

}
func main() {
	fmt.Printf("%+v\n", NewAlarm("8:00"))
}

// {Time:8:00 Sound:Klaxon}
```

## 比较结构体

// 判断两个结构体是否相等，使用 ==
// 判断两个结构体是否不相等，使用 !=
// 不能对两个类型不同的结构体进行比较，否则将导致编译错误

```go example32
package main

import "fmt"

type Drink struct {
	Name string
	Ice  bool
}

func main() {
	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	b := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	if a == b {
		fmt.Println("a and b are the same")
	}
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
}

// a and b are the same
// {Name:Lemonade Ice:true}
// {Name:Lemonade Ice:true}
```

## 检查结构体的类型

```go example33
package main

import (
	"fmt"
	"reflect"
)

type Drink struct {
	Name string
	Ice  bool
}

func main() {
	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	b := Drink{
		Name: "Lemonade",
		Ice:  true,
	}

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
}

// main.Drink
// main.Drink
```

## 区分指针引用和值引用

以值引用的方式复制结构体

将指向结构体的变量赋给另一个变量时，a与b相同，但它是b的副本，而不是指向b的引用，修改b不会影响a，反之亦然。

```go example34
package main

import "fmt"

type Drink struct {
	Name string
	Ice  bool
}

func main() {
	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	b := a
	b.Ice = false
	fmt.Printf("%+v\n", b)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
}

// {Name:Lemonade Ice:false}
// {Name:Lemonade Ice:true}
// 0xc00000c080
// 0xc00000c0a0
```

要修改原始结构体实例包含的值，必须使用指针。

指针是指向内存地址的引用，因此使用它的操作的不是结构体的副本，而是其本身。

```go example35
package main

import "fmt"

type Drink struct {
	Name string
	Ice  bool
}

func main() {
	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	// 将指向a的指针（而不是a本身）赋值给b
	b := &a
	// 修改b，将修改分配给a的内存，因为a和b指向相同的内存
	b.Ice = false
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%+v\n", *b)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", &a)
}

// {Lemonade false}
// &{Lemonade false}
// {Name:Lemonade Ice:false}
// {Name:Lemonade Ice:false}
// 0xc00000c080
// 0xc00000c080
```

## 使用方法

方法类似于函数

在关键字func 后面添加了另一个参数部分，用于接受单个参数

```go example36
package main

import (
	"fmt"
	"strconv"
)

type Movie struct {
	Name   string
	Rating float64
}

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
```


```go example37

```


```go example38

```


```go example39

```

```go example40

```



