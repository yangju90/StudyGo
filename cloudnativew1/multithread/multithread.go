package main

import(
	"fmt"
	"time"
	"math/rand"
)

// channel 的通信是同步的
// channel 当缓冲区满时，数据发送是阻塞的
// 通过make关键字创建通道时可定义缓冲区的容量，默认缓冲区容量为0 (通道两边都必须就绪才可以执行)
// 可以通过channel 来协调goroutine执行顺序

// channel 声明传递数据
func chanTest(){
	ch := make(chan int)
	go func() {
		fmt.Println("gorountine  send int msg")
		ch <- 1
	}()
	// 读不到数据会阻塞
	i :=<- ch

	fmt.Println(i)
}

// 遍历通道数据， 通过for range 获取通道中的数据
// 只发送通道 var sendOnly chan <- int
// 只接受通道 var readOnly <- chan int

func forRangeTest(){
	ch1 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			rand.Seed(time.Now().UnixNano())
			// n := rand.Intn(10)
			fmt.Println("Put:" , i)
			ch1 <- i
		}
		// 内部方法栈未结束， channel 未回收
		close(ch1)
	}()

	for v := range ch1 {
		fmt.Println("receiving: ", v)
	}

	
	// time.Sleep(time.Second)
}


func main(){
	// 启动goroutine
	go fmt.Println("a")
	go fmt.Println("b")
	go fmt.Println("c")
	time.Sleep(time.Second)

	// chanTest()
	forRangeTest()
	
}