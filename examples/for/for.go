// `for` は Go においてループを書くための唯一のキーワードである。
// 基本的な `for` ループの書き方をいくつか紹介する。

package main

import "fmt"

func main() {

	// 1つだけ条件を書く最も基本的な書き方
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 初期化、条件、ループ間処理を書く古典的な `for` ループ
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 条件を書かない `for` ループは、`break` でループから抜けるか、`return` で関数から抜けるまで繰り返し続ける。
	for {
		fmt.Println("loop")
		break
	}

	// `coutinue` と書くと、ループ内の次の繰り返しに進む。
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
