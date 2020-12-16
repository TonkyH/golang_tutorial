package main

import "fmt"

func main() {
	// 明示的な型を指定せずに変数を宣言する場合、右側の変数から型推論される
	v := 3.14
	fmt.Printf("v is of type %T\n", v)
}
