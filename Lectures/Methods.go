package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// メソッドはレシーバ引数を関数にとる
// func レシーバ引数 メソッド名 の順番
// 以下の例では、Absメソッドがvという名前のVertex型のレシーバを持つことを意味する
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
