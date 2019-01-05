// <em>スライス</em>は Go のデータ型の中でも特に重要だ。
// スライスは値の列を表す、配列よりも強力なインターフェースである。

package main

import "fmt"

func main() {

    // 配列とは違い、スライスの型はは要素の型だけを含む（つまり、要素数は含まない）。
    // 空でない（長さ0でない）スライスを作るには、組み込みの `make` を使う。
    // ここでは、文字列が3つ入るスライスを作っている（初期値はゼロ値である）。
    s := make([]string, 3)
    fmt.Println("emp:", s)

    // 値の読み書きは配列と同様に行える。
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    // `len` はスライスの長さを返す。
    fmt.Println("len:", len(s))

    // 配列にあった基本的な操作に加えて、スライスにはより豊富な操作が可能である。
    // 例えば組み込みの `append` は一つかそれ以上の新たな値を含むスライスを返す。
    // ここで、新たな値を得るには、`make` の返り値を受け取る必要があることに注意する。
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    // スライスはコピー（`copy`）することもできる。
    // ここでは、`s` と同じ長さの空のスライス `c` を作り、`s` の内容を `c` にコピーしている。
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    // スライスを「スライス」する `slice[low:high]` という記法もある。
    // 例えばこの例では、`s[2]`、`s[3]`、`s[4]` からなるスライスを取得している。
    l := s[2:5]
    fmt.Println("sl1:", l)

    // これは `s[5]` までの要素（`s[5]` を含まない）をスライスしている。
    l = s[:5]
    fmt.Println("sl2:", l)

    // これは `s[2]` から先の要素（`s[2]` を含む）をスライスしている。
    l = s[2:]
    fmt.Println("sl3:", l)

    // スライスの宣言と定義も、一行で済ませられる。
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    // スライスを組み合わせて多次元のデータ構造を作れる。
    // 内側のスライスの長さは同じでなくてもよく、これは配列の場合とは異なる。
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
