package main

import "fmt"

// 定数はConstを使って宣言する
// 定数は文字、文字列、真偽値、数値を定義する時のみ使える
// 定数は :=　を使って宣言できない

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
