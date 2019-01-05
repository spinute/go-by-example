// Go において<em>関数</em>は中心的な存在だ。
// 例を見ながら関数について学ぼう。

package main

import "fmt"

// これは2つの `int` 型の値を受け取り、その和を表す `int` 型の値を返す関数である。
func plus(a int, b int) int {

    // Go では明示的な return 文を書く必要がある。
    // すなわち、return を書かなければ、最後の式の値を自動で返しはしない。
    return a + b
}

// 同じ型のパラメータを続けて複数取るとき、それらの型名はそのうち最後のものの後にだけ書けばよい。
func plusPlus(a, b, c int) int {
    return a + b + c
}

func main() {

    // `name(args)` と書けば関数を呼べる
    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
}
