package main

import (
	"fmt"
	"time"
)

var messages = make(chan int, 10)
var done = make(chan bool)


func consumer(message chan int, done chan bool){
	ticker := time.NewTicker(time.Second)

	for _ = range ticker.C {
		select{
		case <- done:
			fmt.Println("child process interrupt")
			return
		default:
			fmt.Println("receive message ", <- message)
		}
	}
}


func prod(msg <- chan int){
	for i := 0 ; i< 20; i++ {
		messages <- i
	}
}

func main(){
	go prod(messages)

	go consumer(messages, done)

	time.Sleep(time.Second * 14)
	close(done)
	time.Sleep(time.Second * 1)
	fmt.Println("Main process exit")
}