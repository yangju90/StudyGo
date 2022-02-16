package main

import "fmt"

var(
	source = [5]string{"I","am","stupid","and","weak"}
)

func main() {
	fmt.Println(source)
	solution()
	fmt.Println(source)
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