// Go 言語における _array_ （配列）とは、決まった数だけ要素を並べた列である。

package main

import "fmt"

func main() {

	// `a` は int 型の値5個からなる配列。
	// 配列の型を見ると要素の型と個数がわかる。
	// 配列を宣言すると、要素はゼロ値（int 型の場合は0）になる。
	var a [5]int
	fmt.Println("emp:", a)

	// 配列の index 番目の値を value にするには `array[index] = value` と書く。
	// index 番目の値を読み出すには `array[index]` と書く。
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// 組み込み関数 `len` は配列の長さを返す。
	fmt.Println("len:", len(a))

	// このように書けば、配列の宣言と初期化を一行で済ませられる。
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 配列型は一次元である。
	// しかし、配列型を組み合わせて多次元のデータ構造を作ることもできる。
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
