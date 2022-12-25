package main

import (
	"fmt"
	"time"
)

// 这道题我琢磨了半天不知道要我写什么，应该是让我写并发，就写了“出现”的并发
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("出现")
	say("chuxian")
}
