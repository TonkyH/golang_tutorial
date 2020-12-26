package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// goroutineは、Goのランタイムに管理される軽量なスレッド
	// 先頭に go をつけるだけ（ここでは go say("world"))
	// sayの評価は実行元のgoroutineで実行され
	// sayの実行は新しいgoroutineで実行される
	go say("world")
	say("hello")
}
