package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // sumをcに送信
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	// channel型は、<- を用いて値の送受信ができる通り道
	// 通常、片方が準備できるまで送受信がブロックされる

	// 以下のようにチャネルを作成できる
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // cから値を受信

	fmt.Println(x, y, x+y)
}
