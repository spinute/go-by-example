// Go の値には様々な型がある。
// 例えば、文字列、整数、実数、真偽値などである。
// いくつか基本的な例を紹介しよう。

package main

import "fmt"

func main() {

	// 文字列である。`+` を使って連結できる。
	fmt.Println("go" + "lang")

	// 整数と実数である。
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// 真偽値である。ブール演算子を使えば、期待した振る舞いをするはずだ。
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
