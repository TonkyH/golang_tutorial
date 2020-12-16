package main

import "fmt"

// 基本的に、 for ループはセミコロン ; で3つの部分に分かれている
// 1. 初期化ステートメント: 最初のイテレーション(繰り返し)の前に初期化が実行
// 2. 条件式: イテレーション毎に評価
// 3. 後処理ステートメント: イテレーション毎の最後に実行
// ループは条件式の評価がFalseになったら終了

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	// 初期化と後処理ステートメントの記述は任意
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
