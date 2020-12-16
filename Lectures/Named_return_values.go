package main

import "fmt"

// 戻り値となる変数に名前をつけることができる。
// 今回は戻り値となる変数にx, yと名前をつけている。
// 戻り値に名前をつけるとreturn で何も書かなくていい。これを"naked return"という。
// naked returnは短い関数でのみ利用するべき(長い関数で利用すると可読性に悪影響)
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
