// <em>ゴルーチン</em>は軽量なスレッドである。

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// `f(s)` という関数を呼ぶとしよう。
	// 普通にこのように書けばよく、関数は同期的に実行される。
	f("direct")

	// `go f(s)` と書くと、この関数をゴルーチンの中で実行する。
	// こうすると、新たなゴルーチンが呼び出し側と平行に実行される。
	go f("goroutine")

	// 無名関数のゴルーチンを起動することもできる。
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 2つの関数呼び出しは、別々のゴルーチンにおいて非同期的に実行されている。
	// そのため処理はここまで抜けてくる。
	// ここで処理が終わるのを待つ（[WaitGroup](waitgroups) を使うとより確実だ）。
	time.Sleep(time.Second)
	fmt.Println("done")
}
