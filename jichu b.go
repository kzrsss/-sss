package main

import "fmt"

// 因为if a{}语句执行后直接return后直接跳过了输出3的指令
func main() {
	var a = true
	defer func() {
		fmt.Println("1")
	}()

	if a {
		fmt.Println("2")
		return
	}

	defer func() {
		fmt.Println("3")
	}()
}
