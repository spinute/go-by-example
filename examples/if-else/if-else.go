// Go における `if` と `else` を使った分岐の書き方は素直なものだ。

package main

import "fmt"

func main() {

	// 基本的な例
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// else がなくても `if` 文を書ける
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 条件の前に文を書いてもよい。
	// この文で宣言された変数は、いずれの分岐からも見える。
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// 注意点として、条件を囲む丸括弧は必要ないが、波括弧は必須である。
