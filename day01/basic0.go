package main

import "fmt"

// 标准赋值
var (
	name string
	age  int
	isOk bool
)

// 推导声明赋值
var age1 = 20

// 定义为常量后，不能在赋值
const SIGN = "DAY01"

const (
	statusOK = 200
	notFound = 404
)

// n1、n2、n3  值相同，不写与上行一致
const (
	n1 = 100
	n2
	n3
)

// iota常量计数器，在const关键字出现时将被重置为0，const中每新增一行常量声明将使iota计数一次
const (
	_ = iota
	a1 = iota
	a2
	a3
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

const (
	p1 = iota //0
	p2        //1
	_
	p4        //3
)

// 多个 iota 定义在一行
const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

 
func main() {
	isOk = true
	name = "mat"
	age = 30

	// 简单打印
	fmt.Print(isOk)
	fmt.Println()
	fmt.Printf("this is mime name: %s \n", name)
	fmt.Printf("I'm %d year old", age)
	fmt.Println()
	fmt.Println(age1)


	// 函数内部局部变量赋值
	studentName := "王鹏博"

	fmt.Println(studentName)

	// 函数返回值（匿名变量 _）
	a,_ := foo();
	_,b := foo();
	
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)


	fmt.Println(a1 , " ",  a2, " ", a3)
	fmt.Println(p1 , " ",  p2, " ", p4)
}

func foo()(int, string)  {
	return 10,"func returned"
}
