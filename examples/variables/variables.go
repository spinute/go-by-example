// Go の<em>変数</em>は明示的に宣言され、コンパイラによってチェックされる（例えば、関数呼び出しの型安全性など）。

package main

import "fmt"

func main() {

    // `var` を使って、1つまたはそれ以上の変数を宣言できる。
    var a = "initial"
    fmt.Println(a)

    // 複数の変数を一度に宣言することもできる。
    var b, c int = 1, 2
    fmt.Println(b, c)

    // 変数を初期化すれば Go は変数の型を推測する。
    var d = true
    fmt.Println(d)

    // 明示的に初期化していない変数の値は<em>ゼロ値</em>である。
    // 例えば、`int` のゼロ値は `0` である。
    var e int
    fmt.Println(e)

    // `:=` は宣言と初期化を行うための省略記法である。
    // ここの例は `var f string = "short"` と書くのと同じ意味である。
    f := "short"
    fmt.Println(f)
}
