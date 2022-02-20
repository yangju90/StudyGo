package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
	"strings"
)

// 必须打包运行
func main() {
	queue := make(chan int, 10)
	baseCtx := context.Background()

	ctx, cancel := context.WithCancel(baseCtx)

	go prod(queue, ctx)

	go consumer(queue, ctx)

	for {
		fmt.Println("请输入：")
		read := bufio.NewReader(os.Stdin)
		s, _ := read.ReadString('\n')
		if strings.Compare(strings.TrimSpace(s) , "test") == 0 {
			fmt.Println("ready !!! :", s)
			cancel()
			break
		}
	}

}

func prod(queue chan<- int, ctx context.Context) {
	ticker := time.NewTicker(time.Second / 2)
	for t := range ticker.C {
		queue <- t.Second()
		select {
		case f := <-ctx.Done():
			fmt.Println("prod process interrupt: ", f)
			return
		default:
			fmt.Println("send time : ", t.Second())
		}
	}
}

func consumer(queue <-chan int, ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	for t := range ticker.C {
		select {
		case f := <-ctx.Done():
			fmt.Println("consumer process interrupt: ", f)
			return
		default:
			fmt.Println("received time : ", <-queue, " now: ", t.Second())
		}
	}
}
