package main

import(
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	queue []string
	cond *sync.Cond
}

func (q *Queue) Enqueue(item string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.queue = append(q.queue, item)
	fmt.Println("Add item : " , item)
	q.cond.Broadcast()
} 

func (q *Queue) Dequeue() {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if len(q.queue) == 0 {
		fmt.Println("queue is empty")
		q.cond.Wait()
	}

	fmt.Println("receive msg:" , q.queue[0])
	q.queue = q.queue[1:]
}

func test(q *Queue){
	q.queue = append(q.queue, "test")
	time.Sleep(time.Second *1)
}


func main(){

	queue := Queue{
		queue : []string{},
		cond : sync.NewCond(&sync.Mutex{}),
	}

	test(&queue)

	fmt.Println(len(queue.queue))

	// go func (){
	// 	for{
	// 		queue.Enqueue("msg")
	// 		time.Sleep(time.Second *2)
	// 	}
	// }()


	// for{
	// 	queue.Dequeue()
	// }
}