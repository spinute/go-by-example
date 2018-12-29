// Go は_多値の返り値_を組み込みでサポートしている。
// この機能は Go のイディオムで使われることが多く、例えば結果とエラーの療法を関数から返す場合に使う。

package main

import "fmt"

// この関数のシグネチャにおける `(int, int)` は、この関数が2つの `int` 型の値を返すことを表している。
func vals() (int, int) {
	return 3, 7
}

func main() {

	// ここでは2つの返り値を_多重代入_を使って受け取っている。
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// もし返り値のうちの一部だけが必要なら、ブランク識別子 `_` を使えばよい。
	_, c := vals()
	fmt.Println(c)
}
