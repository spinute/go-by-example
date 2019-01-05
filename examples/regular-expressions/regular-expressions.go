// Go は組み込みの[正規表現](http://en.wikipedia.org/wiki/Regular_expression)処理機能を持っている。
// Go でよく使う正規表現関係の処理を紹介する。

package main

import "bytes"
import "fmt"
import "regexp"

func main() {

    // 文字列がパターンにマッチするかどうかを判定する。
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)

    // 上では文字列としてパターンを直接使った。
    // しかし、他のことをするには `Compile` して最適化された `Regexp` 構造体にしておく必要がある。
    r, _ := regexp.Compile("p([a-z]+)ch")

    // この構造体は色々なメソッドを持っている。
    // まずは先程の、マッチするかどうかの判定をしてみる。
    fmt.Println(r.MatchString("peach"))

    // マッチした文字列を見つける。
    fmt.Println(r.FindString("peach punch"))

    // 同様にマッチした文字列を見つけるが、マッチした文字列自体ではなく、その最初と最後のインデックスを返す。
    fmt.Println(r.FindStringIndex("peach punch"))

    // `Submatch` はパターンにマッチした文字列全体と、その中の部分マッチの情報を共に返す。
    // 例えば、この場合は `p([a-z]+)ch` と `([a-z]+)` のマッチ結果が返ってくる。
    fmt.Println(r.FindStringSubmatch("peach punch"))

    // 同様にこちらはマッチした文字列全体と部分マッチの、インデックス情報を返す。
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))

    // `All` が付いているものははじめのマッチひとつだけではなく、入力のうちマッチした箇所すべてを処理するものだ。
    // 例えば、すべてのマッチを見つけるには次のようにする。
    fmt.Println(r.FindAllString("peach punch pinch", -1))

    // これまで見てきた他の関数にも `All` がついたものがある。
    fmt.Println(r.FindAllStringSubmatchIndex(
        "peach punch pinch", -1))

    // 第二引数に非負整数を渡すと、マッチの数に上限を設ける。
    fmt.Println(r.FindAllString("peach punch pinch", 2))

    // これまで見た例は文字列を引数に取り、`MatchString` のような名前であった。
    // `[]byte` を引数に取るものもある。
    // 関数名から `String` を外せば、対応する `[]byte` を引数に取る関数名になる。
    fmt.Println(r.Match([]byte("peach")))

    // 正規表現の定数を作るときは、`Compile` ではなく `MustCompile` を使う。
    // `Compile` は値を2つ返すので、定数には使えない。
    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println(r)

    // `repexp` パッケージを使って、部分文字列を別の値に置き換えることもできる。
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

    // `Func` が名前に付いたものは、マッチした文字列を引数に渡した関数で変換する。
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}
