// Go で書いたプログラムで [Unix シグナル](http://en.wikipedia.org/wiki/Unix_signal)を受付け、処理したいことがある。
// 例えば、`SIGTERM` を受け取ったらサーバーを安全に停止したり、`SIGINT` を受け取ったらコマンドライン処理を中断したりといった具合である。
// ここでは Go でチャネルを使ってシグナルをどう扱うのかを紹介する。

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Go ではシグナルが通知されたとき、`os.Signal` 型の値がチャネルに届く。
	// これらの値を受け取るためのチャネルを作る
	// （また、プログラムが終了できることを知らせるためのチャネルも作る）。
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// `signal.Notify` に渡したチャネルには、指定したシグナルの通知が届くようになる。
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// このゴルーチンはシグナルを受信するためにブロックする。
	// シグナルを受信すると、それを表示し、プログラムに終了できることを知らせる。
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// プログラムはここでシグナルが届くのを待って、終了する（シグナルが届いたことは、上のゴルーチンが `done` に値を送信したことから判断する）。
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
