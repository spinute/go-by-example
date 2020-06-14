// [<em>コマンドライン引数</em>](http://en.wikipedia.org/wiki/Command-line_interface#Arguments)はプログラムに実行時パラメータを与えるためによく使われる。
// 例えば `go run hello.go` はプログラム `go` の引数として、`run` と `hello.go` を与えている。

package main

import (
	"fmt"
	"os"
)

func main() {

	// `os.Args` を使って生のコマンドライン引数にアクセスできる。
	// ただし、このスライスの一番目の値はプログラムのパスである。
	// そのため `os.Args[1:]` がプログラムの引数である。
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// インデックスを指定して個々の引数を取得できる。
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
