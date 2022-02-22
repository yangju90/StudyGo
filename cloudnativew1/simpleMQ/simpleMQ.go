package main

import(
	"fmt"
	"time"
)

/**
* 通道无需每次关闭
* 关闭的作用时告诉接受者该通道再无新数据发送
* 只有发送方需要关闭通道
*/
var mq = make(chan int)

func prod(p chan <- int){

	defer close(p)
	for i := 0; i< 10; i++ {
		p <- i
	}
}

func consumer(c <- chan int){
	for i := range c {
		fmt.Println(i)
	}
}

func main(){
	go consumer(mq)
	go prod(mq)

	

	time.Sleep(time.Second)

	if v, notClose := <- mq; notClose{
		fmt.Println(v)
	}
	time.Sleep(time.Second)


}