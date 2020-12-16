package main

import "fmt"

func main() {
	i, j := 42, 2701

	// メモリのアドレスを保持する変数をポインタ
	p := &i         //iのメモリアドレスをpに
	fmt.Println(*p) // ポインタpを通してiから値を読み出す
	*p = 21         // ポインタpを通したiを値へ代入する
	fmt.Println(i)

	p = &j       // jのメモリアドレスをpに代入
	*p = *p / 37 // ポインタpを通してjの値を読み出す。計算結果をポインタpを通してjに代入
	fmt.Println(j)
}
