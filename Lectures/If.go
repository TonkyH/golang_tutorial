package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// 条件式の前に、評価するための変数を用意できる(ここでいうと変数vのこと)
	// ここで宣言された変数は、if分の内だけで使える(ここでいうと変数vのこと)
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// ifの外では変数vは使えない
	return lim
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-4))
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))
}
