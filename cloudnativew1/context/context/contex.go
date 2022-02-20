package main

import(
	"fmt"
	"context"
	"time"
)

// context 传递上下文参数
func simpleContext(c context.Context) {
	go func(c context.Context) {
		fmt.Println("simpleContext", c.Value("key"))
	}(c)
}


func timeoutContext(c context.Context, cancel context.CancelFunc){
	ticker := time.NewTicker(time.Second /2)
	defer cancel()

	for _ = range ticker.C {
		select {
		case <- c.Done():
			fmt.Println("child process interrupt ... ")
			return
		default:
			fmt.Println("enter default")
		}
	}

}

func main(){
	baseCtx := context.Background()
	parameterCtx := context.WithValue(baseCtx, "key", "value")
	go simpleContext(parameterCtx)

	// 定时处理，select使用
	timeoutCtx, cancel := context.WithTimeout(baseCtx, 2*time.Second)
	go timeoutContext(timeoutCtx, cancel)
	
	time.Sleep(time.Second * 3)

	select{
	case <- timeoutCtx.Done():
		fmt.Println("Main process interrupt ... ")
	default:
		fmt.Println("Main default")
	}

	fmt.Println("Error message : ",  timeoutCtx.Err())


}