package day05

import (
	"fmt"
)

func Day05() {
	fmt.Println("=============day05===========")
	ptrDemo()
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
