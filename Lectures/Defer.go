package main

import "fmt"

func main() {
	// defer ステートメントは defer へ渡した関数の実行を、呼び出し元の関数の終わり(return)まで遅延させる
	// defer へ渡した関数の引数は、すぐに評価されるが、その関数自体は呼び出し元の関数がreturnされるまで実行されない
	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("counting")
	// defer へ渡した関数が複数ある場合、その呼び出しはスタック(stack)される
	// 呼び出し元の関数がreturnするとき、deferへ渡した関数はLIFO(last-in-first-out)の順で実行される
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
