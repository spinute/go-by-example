// Go は[<em>無名関数</em>](http://en.wikipedia.org/wiki/Anonymous_function)をサポートしている。
// 無名関数を使って<a href="http://en.wikipedia.org/wiki/Closure_(computer_science)"><em>クロージャ</em></a>を作れる。
// 無名関数は名前を付けずにインラインで関数を定義でき、便利である。

package main

import "fmt"

// この関数`intSeq` は、`intSeq` の中で定義した無名関数を返す。
// 返される関数に変数 `i` を<em>閉じ込めて</em>おり、クロージャを作っている。
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {

    // `intSeq` を呼んで、その返り値（関数）を `nextInt` に代入している。
    // この関数値は独自の `i` を持っていて、`nextInt` を呼ぶたびにその `i` の値が更新される。
    nextInt := intSeq()

    // `nextInt` を何度か呼んでみてクロージャの効果を確認しよう。
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    // 関数ごとに状態が独立であることを確認するため、新しいクロージャを作ってみよう。
    newInts := intSeq()
    fmt.Println(newInts())
}
