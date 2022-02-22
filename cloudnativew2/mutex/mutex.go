package main

import (
	"fmt"
	"sync"
)

func MutexLock() {
	l := sync.Mutex{}

	for i := 0; i< 3; i++ {
		l.Lock()
		defer l.Unlock()
		fmt.Println("Mutex: " , i)
	}

	fmt.Println("Mutex Procss finish !")
}

func RwMutexLockR(){
	l := sync.RWMutex{}

	for i := 0; i< 3; i++ {
		l.RLock()
		defer l.RUnlock()
		fmt.Println("RWMutex R: " , i)
	}

	fmt.Println("RWMutex R Procss finish !")
}

func RwMutexLockW(){
	l := sync.RWMutex{}

	for i := 0; i< 3; i++ {
		l.Lock()
		defer l.Unlock()
		fmt.Println("RWMutex: " , i)
	}

	fmt.Println("RWMutex Procss finish !")
}

func WgMutexLock(){
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i< 100; i++{
		go func(){
			fmt.Println("exec " , i)
			defer wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("wait group prcess finish !")
}

func main() {
	go MutexLock()
	go RwMutexLockR()
	go RwMutexLockW()

	WgMutexLock()

}