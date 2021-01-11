package fmtDemo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。
func ScanDemo() {
	var (
		name    string
		age     int
		married bool
	)

	fmt.Scan(&name, &age, &married)
	fmt.Printf("扫描结果 name: %s age: %d married: %t \n", name, age, married)
}

func ScanfDemo() {
	var (
		name    string
		age     int
		married bool
	)

	fmt.Scanf("1: %s 2: %d 3: %t", &name, &age, &married)
	fmt.Printf("扫描结果 name: %s age: %d married: %t \n", name, age, married)
}

//Scanln类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置
func ScanlnDemo() {
	var (
		name    string
		age     int
		married bool
	)

	fmt.Scanln(&name, &age, &married)
	fmt.Printf("扫描结果 name: %s age: %d married: %t \n", name, age, married)
}

/*
* bufio
* 有时候我们想完整获取输入的内容，而输入的内容可能包含空格，这种情况下可以使用bufio包来实现
 */

func BufioDemo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("请输入内容: ")

	text, _ := reader.ReadString('\n') // 读到换行
	fmt.Println(text)
	text = strings.TrimSpace(text)
	fmt.Println(text)
}

// Fscan

// Sscan
func SscanDemo() {
	s := "我叫瓦利王，来找我的儿子"
	var (
		name string
		work string
	)
	fmt.Sscanf(s, "我叫%s，来找我的%s", &name, &work)

	fmt.Println("谁来了？")
	fmt.Println(name)
	fmt.Println("找谁？")
	fmt.Println(work)
}
