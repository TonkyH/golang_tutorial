package main

import "fmt"

func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	// このfibonacciはメモ化してないためforを回しすぎると計算が終わらない
	f := fibonacci()
	for i := 0; i < 5; i++ {
		fmt.Println(f())
	}
}
