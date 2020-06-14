// チャネルを関数に渡すとき、そのチャネルが送信専用・受信専用であることを明記できる。
// こうするとプログラムの型安全性が増す。

package main

import "fmt"

// この関数 `ping` の引数であるチャネル `pings` は送信専用である。
// このチャネルから受信しようとすれば、コンパイルエラーになる。
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// この関数 `pong` の引数であるチャネル `pings` は受信専用、`pongs` は送信専用である。`
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
