package day05

import (
	"fmt"
)

func Day05() {
	fmt.Println("=============day05===========")
	ptrDemo()
	demoMakeNew()
}

func ptrDemo() {
	a := 10
	moditfyV(a)
	fmt.Println(a)
	moditfyPtrV(&a)
	fmt.Println(a)
}

func moditfyV(x int) {
	x = 100
}

func moditfyPtrV(x *int) {
	*x = 1000
}

// make 与new 用来分配内存
//二者都是用来做内存分配的。
//make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
//而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
func demoMakeNew() {
	var c *int
	c = new(int)
	fmt.Println(c)

	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false

	var m map[string]int
	m = make(map[string]int, 10)
	m["张三"] = 20
	fmt.Println(m["张三"])

	m1 := make(map[string]int, 20)
	m1["李四"] = 30
	fmt.Println(m1)

	fmt.Println("===============指针地址=================")
	p := 10
	q := &p
	fmt.Printf("value: %d ptr: %p \n", p, &p)
	fmt.Printf("value: %p ptr: %p \n", q, &q)
	fmt.Println(&q)
}
