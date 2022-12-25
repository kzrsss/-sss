package main

import (
	"fmt"
	"sync"
)

// 可以通过chan顺序输出
var count = 1

func main() {
	wg := sync.WaitGroup{}
	chanA := make(chan struct{}, 1)
	chanB := make(chan struct{}, 1)
	chanC := make(chan struct{}, 1)
	chanA <- struct{}{}
	wg.Add(3)
	go printA(&wg, chanA, chanB)
	go printB(&wg, chanB, chanC)
	go printC(&wg, chanC, chanA)
	wg.Wait()
}
func printA(wg *sync.WaitGroup, chanA, chanB chan struct{}) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		<-chanA
		fmt.Println("goroutine1")
		chanB <- struct{}{}
	}
}
func printB(wg *sync.WaitGroup, chanB, chanC chan struct{}) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		<-chanB
		fmt.Println("goroutine2")
		chanC <- struct{}{}
	}
}
func printC(wg *sync.WaitGroup, chanC, chanA chan struct{}) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		<-chanC
		fmt.Println("goroutine3")
		chanA <- struct{}{}
	}
	fmt.Println("successful")
}
