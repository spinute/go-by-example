// 文字列から数値へのパースは、多くのプログラムに現れるパターンである。
// ここでは、Go でのやり方を紹介する。

package main

// `strconv` パッケージを使って数をパースできる。
import "strconv"
import "fmt"

func main() {

	// `ParseFloat` の `64` は何ビットの精度でパースするかを指定する。
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// `ParseInt` の `0` は基数を文字列から推測させることを表す。
	// `64` は、結果が64ビットに収まることを要求する。
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// `ParseInt` は8進記数法で書かれた文字列をパースできる。
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// `ParseUint` もある。
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// `Atoi` は基数が10の `int` 型の値をパースするときに便利である。
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// パースする関数は、入力の形式が正しくないとエラーを返す。
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
