package main

import "fmt"

// var宣言では初期化子(initializer)を与えられる
var i, j int = 1, 2

func main() {
	// 初期化子(initializer)が与えられている場合、型の宣言は省略できる
	// 変数は初期化子(initializer)が持つ型になる
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
