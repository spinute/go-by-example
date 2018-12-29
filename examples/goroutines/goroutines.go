// _ゴルーチン_は軽量なスレッドである。

package main

import "fmt"

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
	// ここにある `Scanln` はキーを押すまでブロックするため、プログラムが終了するのを妨げてくれる。
	fmt.Scanln()
	fmt.Println("done")
}
