// Go は<em><a href="http://en.wikipedia.org/wiki/Pointer_(computer_programming)">ポインタ</a></em>をサポートしている。
// ポインタを使うと、プログラム中で値やレコードへの参照を渡すことができる。

package main

import "fmt"

// ここではポインタと値の振る舞いの違いを、`zeroval` と `zeroptr` の2つの関数を使って見てもらう。
// `zeroval` は `int` 型の引数を取るので、引数は値として渡される。
// `zeroval` は `ival` のコピーを渡されるが、この実体は関数を呼び出す側で引数に渡した変数とは別のものである。
func zeroval(ival int) {
	ival = 0
}

// 一方、`zeroptr` の引数の型は `*int` である。
// すなわち、引数は `int` 型のポインタとして渡される。
// 関数の中の `*iptr` と書いた箇所では、ポインタの<em>参照を剥がし</em>ている。
// これは、メモリ上のアドレスから、その番地にある現在の値にアクセスする操作だ。
// 剥がされたポインタに値を代入すると、ポインタが参照していたアドレスに書かれた値が書き換わる。
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// `&i` という記法で、`i` のメモリアドレスを取得できる。
	// このメモリアドレスは、`i` のポインタである。
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// ポインタも表示できる。
	fmt.Println("pointer:", &i)
}
