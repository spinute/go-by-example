// Go は文字、文字列、真偽値、数値の<em>定数（constant）</em>をサポートしている。

package main

import (
	"fmt"
	"math"
)

// `const` を使って定数を宣言できる。
const s string = "constant"

func main() {
	fmt.Println(s)

	// `const` 文は `var` 文が書けるところならどこでも書ける。
	const n = 500000000

	// 定数式では任意精度で算術を実行する。
	const d = 3e20 / n
	fmt.Println(d)

	// 数値の定数は、明示的にキャストするなどして型が決まるまでは、型を持たない。
	fmt.Println(int64(d))

	// 数値のは型を必要とするコンテキストによって与えられる。
	// 変数の代入や、関数呼び出しが型を必要とする場面の例である。
	// 例えばここでは、`math.Sin` は `float64` 型の引数を受け取る。
	fmt.Println(math.Sin(n))
}
