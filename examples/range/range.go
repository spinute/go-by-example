// _range_ は様々なデータ構造の要素の値を繰り返し取得するのに役立つ。
// これまで学んだデータ構造に `range` を使う方法を紹介する。

package main

import "fmt"

func main() {

	// ここでは `range` を使ってスライスの要素の和を計算している。
	// 配列でも同様に `range` を使える。
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// 配列、スライスに対する `range` から、インデックスと値の両方を読み出せる。
	// 上ではインデックスを使わなかったので、ブランク識別子 `_` を使ってインデックスを無視した。
	// しかし、インデックスを使いたい場面ももちろんある。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// マップに `range` を使うと、キーと値の組みを繰り返し読み出す。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// マップのキーだけを繰り返し読み出すこともできる。
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// 文字列に `range` を使うと、Unicode のコードポイントを繰り返し読み出す。
	// 1つ目の返り値は、その `rune` 型の値のはじめのバイトのインデックスであり、2つ目の返り値は `rune` 型の値そのものである。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
