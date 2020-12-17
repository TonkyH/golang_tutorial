package main

import "fmt"

func main() {
	// 組み込み関数の make でスライスを作成
	// 動的サイズの配列を作成できる
	// make はゼロ化された配列を割り当て、その配列を指すスライスを返す
	a := make([]int, 5) //len(a)=5
	printSlice("a", a)

	// make の３番目の引数に、スライスの容量(capacity)を指定できる
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSlice("b", b)

	c := b[:3]
	printSlice("c", c)

	d := c[:5]
	printSlice("d", d)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
