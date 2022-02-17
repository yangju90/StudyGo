package main

import (
	"fmt"
	"os"
	"flag"
)

var(
	source = [5]string{"I","am","stupid","and","weak"}
)

func init(){
	source[0] = "You"
}

// 启动命令 go build
// .\cloudnative01.exe --name abc
func main() {
	name := flag.String("name", "solution1", "2022-02-16 solution1")
	flag.Parse()
	fmt.Println("name:", *name)
	fullString := fmt.Sprintf("Hello World %s ! \n", *name)
	fmt.Print(fullString)

	fmt.Println("启动程序的参数是：" , os.Args)

	fmt.Println(source)
	
	solution()
	fmt.Println(source)
	print(source[0])
	print("\n")


	operate(1, add)
	operate(1, sub)
}

func solution(){
	for index,value := range source {
		fmt.Printf("%s  %d \n", value, index)
		if value == "stupid" {
			source[index] = "smart"
		}else if value == "weak" {
			source[index] = "strong"
		}
	}
}

func operate(x int, f func(int, int)) {
	f(x, 1)
}

func add(x, y int){
	fmt.Println("add:", x + y)
}

func sub(x, y int){
	fmt.Println("sub:", x - y)
}