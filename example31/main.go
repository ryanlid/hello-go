// 自定义结构体数据字段的默认值
// 默认情况下，Go 给数据字段指定相应数据类型的零值
// Go 语言中没有提供自定义默认值的内置方法，但可使用构造函数来实现这个目标。
// 通过构造函数创建结构体将没有指定值的数据字段设置默认值。
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
