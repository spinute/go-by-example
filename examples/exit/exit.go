// `os.Exit` を使うと引数に渡したステータスで直ちにプログラムを終了する。

package main

import "fmt"
import "os"

func main() {

	// `os.Exit` を使うと `defer` は実行_されない_。
	// そのためこの `fmt.Println` は呼ばれない。
	defer fmt.Println("!")

	// ステータス 3 で終了する。
	os.Exit(3)
}

// C などと違って、Go は `main` 関数が返す整数型の値を使って終了ステータスを表すことはない。
// そのため、もしゼロでないステータスを返したければ、`os.Exit` を使わなければならない。
