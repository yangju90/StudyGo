package main

import (
	"fmt"

	"mat.indi.com/StudyGo/day03/fmtDemo"
)

var (
	nums [10]int
)

func main() {
	fmt.Println("数组示例学习")
	fmt.Print(nums[0])

	initA()
	ptrDemo()

	fmt.Println("========================*****************==================")
	fmtDemo.PrintDemo()
	fmtDemo.FprintDemo("./abc.txt")
	fmtDemo.SprintDemo()
	fmtDemo.ErrorDemo("原始错误")
	fmtDemo.Print()

	// 标准输入Scan、bufio、Sscan、Fscan
	// fmtDemo.ScanDemo()
	// fmtDemo.ScanfDemo()
	// fmtDemo.ScanlnDemo()
	fmtDemo.BufioDemo()
	fmtDemo.SscanDemo()
}

func initA() {
	numsArray := [3]int{1, 2, 3}
	numsArrayIndex := [...]int{3: 1, 4: 2, 10: 3}

	strArray := [...]string{"aaa", "bbb"}
	strArrayIndex := [...]string{2: "aaa", 6: "bbb"}

	fmt.Println(numsArray)
	fmt.Println(numsArrayIndex)
	fmt.Println("int数组长度:", len(numsArrayIndex))
	fmt.Println(strArray)
	fmt.Println("字符串数组长度:", len(strArray))
	fmt.Println(strArrayIndex)
	fmt.Println("字符串数组长度:", len(strArrayIndex))

	//注意： 多维数组只有第一层可以使用...来让编译器推导数组长度。
	//数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值
	multiArray := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}

	fmt.Println(multiArray)

	changeArray(multiArray)

	fmt.Println(multiArray)
}

func changeArray(a [3][2]string) {
	a[2][1] = "xxx"

	fmt.Println(a)
}

// [n]*T表示指针数组，*[n]T表示数组指针
func ptrDemo() {
	n := [5]int{5, 6, 7, 8, 9}
	var ptra *[5]int = &n

	fmt.Println()
	fmt.Print(*ptra)
	fmt.Printf("%d", &ptra)
	fmt.Println()
	fmt.Printf("%d", &n[0])
	fmt.Println()

	for i, v := range ptra {
		fmt.Printf("索引:%d 值:%d 值内存地址:%d\n", i, &(ptra[i]), &v)
		fmt.Printf("索引:%d 值:%d 值内存地址:%d\n", i, v, &v)
	}

	var ptrs [5]*int
	for i, _ := range n {
		ptrs[i] = &n[i]
	}
	for i, x := range ptrs {
		fmt.Printf("索引：%d 值a:%d 内存地址：%p\n", i, *x, x)
	}
}
