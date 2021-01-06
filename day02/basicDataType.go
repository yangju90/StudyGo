package main

import (
	"fmt"
	"math"
)

var (
	var32int int32 = 1
	var64int int64 = 1
	a        int   = 10
	b        int   = 077
	c        int   = 0xff

	d bool = true

	varComplex complex64 = 1 + 2i

	s string = "张三"

	ss string = `张三
				李四`
)

func main() {
	pi := math.Pi

	fmt.Println(pi)

	fmt.Printf("%f \n", pi)
	fmt.Printf("%.2f \n", pi)

	fmt.Println(varComplex)
	fmt.Println(s)
	fmt.Println("字符串长度", len(s))

	fmt.Println(ss)
	fmt.Println("字符串长度", len(ss))

	fmt.Printf("%d \n", a) // 10进制输出
	fmt.Printf("%d \n", b)
	fmt.Printf("%d \n", c)
	fmt.Printf("%b \n", a) // 2进制输出
	fmt.Printf("%o \n", b) // 8进制输出
	fmt.Printf("%x \n", c) // 16进制输出

	fmt.Printf("%t \n", d)

	traverseString()
	changeString()
	sqrtDemo()
	manCount()
	gotoDemo()
}

// 字符串底层是一个byte数组，所以可以和[]byte类型相互转换
// UTF8编码下一个中文汉字由3~4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串

func traverseString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c) \n", s[i], s[i])
	}
	// 104(h) 101(e) 108(l) 108(l) 111(o) 230(æ) 178(²) 153() 230(æ) 178(²) 179(³)

	fmt.Println()

	for index, r := range s {
		fmt.Printf("%d :", index)
		fmt.Printf("%v(%c) \n", r, r)
	}
	// 104(h) 101(e) 108(l) 108(l) 111(o) 27801(沙) 27827(河)

	fmt.Println()
}

// Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。
// 强制类型转换的基本语法如下：
// T(表达式) 表达式包括变量、复杂算子和函数返回值等

func changeString() {
	s1 := "big"
	// 强制转换类型
	byteS1 := []byte(s1)
	byteS1[0] = 'p'

	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '紫'

	fmt.Println(string(runeS2))
	fmt.Println(int(runeS2[0]))
	fmt.Println(int(runeS2[1]))
	fmt.Println(int(runeS2[2]))
	fmt.Println()
}

func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
	fmt.Println()
}

func manCount() {
	words := "hello沙河小王子"
	count := 0
	const startCode = 0x2E80
	const endCode = 0x9FFF
	for _, w := range words {
		if startCode < w && endCode > w {
			count++
		}
	}
	fmt.Println("中文字符数:", count)
}

func gotoDemo() {
	//breakTag:
forloop1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 5 && j == 2 {
				//goto gotoTag
				//break breakTag
				continue forloop1
			}

			fmt.Printf("%d ..... %d", i, j)
			fmt.Println()
		}
	}

	return

	//gotoTag:
	fmt.Println("结束循环体")

	fmt.Println("继续执行程序")
}