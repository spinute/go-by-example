// チャネルを使ってゴルーチン間を同期できる。
// この例では、ゴルーチンが処理を終えるのを待つために、受信がブロックされることを利用する。
// 複数のゴルーチンが処理を終えるのを待つときは [WaitGroup](waitgroups) を使う方が好ましいだろう。

package main

import (
	"fmt"
	"time"
)

// この関数をゴルーチンで実行する。
// `done` チャネルはこの関数の仕事が終わったことを、別のゴルーチンに知らせるために使う。
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 終わったことを知らせるために値を送信する
	done <- true
}

func main() {
	// 通知用のチャネルを渡してゴルーチンを起動する。
	done := make(chan bool, 1)
	go worker(done)

	// ワーカーからこのチャネルに通知が届くまでブロックされる。
	<-done
}
