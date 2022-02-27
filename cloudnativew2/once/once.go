package main

import(
	"fmt"
	"sync"
)

type SliceNum []int

func NewSlice() SliceNum{
	return make(SliceNum, 0)
}

// 有值传递的方法，需要使用指针
func (s *SliceNum) Add(elem int) *SliceNum {
	*s = append(*s, elem)
	fmt.Println("add", elem)
	fmt.Println("add SliceNum end", s)
	return s
}


func (s SliceNum) AddNoPoint(elem int) SliceNum {
	s = append(s, elem)
	fmt.Println("add", elem)
	fmt.Println("add SliceNum end", s)
	return s
}

func main(){
	s := NewSlice()
	once := sync.Once{}

	once.Do(func(){s.Add(16)})

	once.Do(func(){s.Add(132)})

	fmt.Println(s)

	s.AddNoPoint(32)
	
	fmt.Println(s)


	fmt.Println(s.AddNoPoint(48))
}