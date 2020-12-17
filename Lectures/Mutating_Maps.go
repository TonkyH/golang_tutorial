package main

import "fmt"

func main() {
	m := make(map[string]int)

	// mへ要素の挿入
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// mの要素の更新
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// mの要素削除
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// keyに対する要素が存在するかどうかは、２つめの引数の値で確認できる
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
