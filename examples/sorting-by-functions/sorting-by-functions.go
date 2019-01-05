// コレクションを特別な順序でソートしたいことがある。
// 例えば、文字列を辞書順ではなく、その長さの順にソートしたいとする。
// ここではそのようにソートをカスタムする方法を紹介する。

package main

import "sort"
import "fmt"

// Go で独自の関数を使ってソートするには、そのための型を定義しなければならない。
// ここでは `byLength` という型を定義するが、これは `[]string` の単なるエイリアスである。
type byLength []string

// `sort.Interface` を実装する。
// これは `Len`、`Less`、`Swap` の3つのメソッドを我々の型に実装するということだ。
// こうすると、`sort` パッケージの汎用な `Sort` 関数を使えるようになる。
// `Len` や `Swap` はどの型でも似たような実装になる。
// `Less` こそが実際にソートのやり方を決める関数である。
// 我々の場合、文字列の長さの昇順にしたいので、`len(s[i])` と `len(s[j])` を使って `Less` を実装する。
func (s byLength) Len() int {
    return len(s)
}
func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

// ここまでの準備をすれば、あとはスライス `fruits` を `byLength` 型にキャストし、`sort.Sort` を呼べば、独自のソートを実装できる。
func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(byLength(fruits))
    fmt.Println(fruits)
}
