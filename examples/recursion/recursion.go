// Go は<a href="http://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>再帰関数</em></a>をサポートしている。
// この例は古典的な階乗を計算するものだ。

package main

import "fmt"

// この関数 `fact` は、ベースケースである `fact(0)` に到達するまで自身を繰り返し呼び出す。
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
