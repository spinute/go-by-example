// Go の `sort` パッケージは、組み込み型もユーザー定義型もソートできる。
// まずは、組み込み型のソートを見てみよう。

package main

import "fmt"
import "sort"

func main() {

	// ソートメソッドは組み込み型ごとに別々である。
	// ここでは文字列をソートするメソッドを使っている。
	// ソートは in-place である。つまり、引数に渡したスライスを変更し、新たなスライスを返すわけではない。
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// `int` 型のスライスをソートする例
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// `sort` パッケージを使って、スライスがソート済みかどうかを確認することもできる。
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
