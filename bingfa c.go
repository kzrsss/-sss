package main

import (
	"fmt"
	"math"
)

func main() {
	for a := 100; a <= 999; a++ {
		x := math.Pow(float64(a/100), 3)
		y := math.Pow(float64(a%100/10), 3)
		z := math.Pow(float64(a%10), 3)
		if x+y+z == float64(a) {
			fmt.Println(a)
		}
	}

}
