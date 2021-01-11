package fmtDemo

import (
	"errors"
	"fmt"
	"os"
)

//实现了C语言中的printf 和 scanf
func PrintDemo() {
	fmt.Print("在终端打印该信息。")
	name := "瓦力王"
	fmt.Printf("我是：%s \n", name)
	fmt.Println("我要啃大瓜")
}

func FprintDemo(file string) {
	fmt.Fprintln(os.Stdout, "项标准输出中输入")
	fileObj, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		fmt.Println("打开文件错, err = ", err)
		return
	}

	name := "我要啃大瓜"
	fmt.Fprintln(fileObj, "项文件中写入:", name)
}

func SprintDemo() {
	name := "瓦力王"
	age := 18

	s1 := fmt.Sprintf("name : %s , age : %d", name, age)
	s2 := fmt.Sprintln("啃大瓜瓦力王")

	fmt.Println(s1, s2)
}

func ErrorDemo(e string) {
	e0 := fmt.Errorf("打印错误输出.....", e)
	fmt.Println(e0)

	e1 := errors.New("自定义基本错误！")
	fmt.Println(e1)
	//Go1.13版本为fmt.Errorf函数新加了一个%w占位符用来生成一个可以包裹Error的Wrapping Error。
	w := fmt.Errorf("Warp了一个错误 %w", e1)
	fmt.Println(w)
	w1 := fmt.Errorf("打印错误%s", w)
	fmt.Println(w1)
}

func Print() {
	fmt.Printf("%v\n", 100)    // 100
	fmt.Printf("%+v\n", 100)   // 100
	fmt.Printf("%v\n", false)  // false
	fmt.Printf("%+v\n", false) // false
	o := struct{ name string }{"小王子"}
	fmt.Printf("%v\n", o)  // {小王子}
	fmt.Printf("%#v\n", o) // struct { name string }{name:"小王子"}
	fmt.Printf("%T\n", o)  // struct { name string }
	fmt.Printf("%+v\n", o) // {name:"小王子"}
	fmt.Printf("100%%\n")  // 100%
}
