package main

import (
	"fmt"
)

// Slice 相当与ArrayList
// 切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合
//要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断

// var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
// s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
// s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil

func main() {
	fmt.Println("======= start Slice Demo ==========")
	demo2()
}

func demo1() {
	var (
		a []string
		b = []int{}
		c = []bool{false, true}
		d = []bool{false, true}
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a == nil)
	fmt.Println(b == nil)
	fmt.Println(c == nil)
	fmt.Println(d == nil)
}

func demo2() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	fmt.Println("\n", s[0], s[1])
	s1 := s[3:4]
	fmt.Println("\n", s1[0], cap(s1))

	m := make([]string, 2, 10)
	fmt.Println(m)
	fmt.Println(cap(m))
	fmt.Println(len(m))
	fmt.Println(m[1])

}
