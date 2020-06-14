// データのコレクションを操作したいことがよくある。
// これは例えば、ある条件を満たすアイテムだけを抽出したり、すべてのアイテムに関数を適用したりといった場合である。

// 汎用データ構造や汎用アルゴリズムを使う慣習がある言語もある（参考：http://en.wikipedia.org/wiki/Generic_programming）。
// Go はそうではない。
// プログラムやデータ型に特に必要なときに限って、コレクション操作のための関数を提供するのが Go のやり方だ。

// `strings` のスライスを操作するためのコレクション操作関数を定義してみる。
// この例を元に自分で関数を作るのもいいだろう。
// なお、ヘルパー関数を作らず、コレクションを操作するコードをインラインで書く方がわかりやすいこともある。

package main

import (
	"fmt"
	"strings"
)

// Index は文字列 `t` を始めに見つけたインデックスを返す。
// もし見つからなければ -1 を返す。
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include はスライスの中に文字列 `t` が含まれるなら `true` を返す。
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any はスライスの中に述語が真になる文字列があれば `true` を返す。
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// Any はスライスに含まれる文字列がいずれも述語を真にするなら `true` を返す。
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter は、元のスライスのうち述語 `f` が真になる文字列だけを含む、新たなスライスを返す。
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map は、元のスライスの要素である各文字列に、関数 `f` を適用した結果を含む新たなスライスを返す。
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {

	// コレクション関数を色々使ってみる。
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	// 上の例ではいずれも無名関数を使っているが、型が適合すれば名前付き関数を使ってもよい。
	fmt.Println(Map(strs, strings.ToUpper))

}
