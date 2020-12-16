package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Go runs on ")
	// switch ステートメントは if - else ステートメントのシーケンスを短く書ける
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	// switch caseは、上から下へcaseを評価する
	// caseの条件が一致すれば、そこで停止(自動的にbreak)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	time := time.Now()
	// 条件式のないswitchは switch true　と書くことと同じ
	// switch caseは、上から下へcaseを評価する
	// このswitchの構造は、長くなりがちな "if-then-else" のつながりをシンプルに表現できる
	switch {
	case time.Hour() < 12:
		fmt.Println("Good morning!")
	case time.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
