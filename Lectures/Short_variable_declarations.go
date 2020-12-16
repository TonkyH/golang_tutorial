package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// 関数内ではvar宣言の代わりに、:= で暗黙的に型宣言ができる
	// 関数外ではvarやfuncなどをつけないといけない
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
