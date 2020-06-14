// [タイマー](timers.html)を使えば、将来のあるタイミングで何かを実行できる。
// 一方、_ticker_ を使えば、定期的に何かを繰り返し実行できる。
// この例では、ticker を使って定期的にメッセージを送る方法を紹介する。

package main

import (
	"fmt"
	"time"
)

func main() {
	// タイマーと同様に ticker も値を送るチャネルを使う。
	// ここでは組み込みの `range` をチャネルに使い、500ミリ秒ごとに届く値を繰り返し受け取っている。
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// ticker の止め方はタイマーと同様だ。
	// ticker を止めると、それ以降そのチャネルは値を受信しなくなる。
	// ここでは1600ミリ秒待って ticker を停止している。
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
