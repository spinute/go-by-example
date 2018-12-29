// _マップ_は Go に組み込みの[連想データ型](http://en.wikipedia.org/wiki/Associative_array)である。
// これは他の言語では_ハッシュ_や_辞書_と呼ばれることがあるものだ。

package main

import "fmt"

func main() {

	// 空のマップを作るには、組み込みの `make` を使って、`make(map[key-type]val-type)` のように書く。
	m := make(map[string]int)

	// キーと値の組みを登録するには、`name[key] = val` のように書く。
	m["k1"] = 7
	m["k2"] = 13

	// マップを例えば `fmt.Println` などを使って表示すると、キーと値の組みをすべて表示する。
	fmt.Println("map:", m)

	// マップから値を読み出すには、`name[key]` のように書く。
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// 組み込みの `len` は、マップに含まれるキーと値の組みの個数を返す。
	fmt.Println("len:", len(m))

	// 組み込みの `delete` はマップからキーと値の組みを削除する。
	delete(m, "k2")
	fmt.Println("map:", m)

	// オプションの2つ目の返り値は、マップにキーが含まれるかどうかを表す真偽値である。
	// これはキーが入っていない場合と、そのキーの値としてゼロ値（`0` や `""` など）が入っている場合とを区別するために使える。
	// ここでは、値はいらないので、_ブランク識別子_ `_` を使って、無視している。
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 以下のように書くと、マップの宣言と初期化を同じ行で済ませられる。
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
