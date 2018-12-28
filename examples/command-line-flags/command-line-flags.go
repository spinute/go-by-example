// [_コマンドラインフラグ_](http://en.wikipedia.org/wiki/Command-line_interface#Command-line_option)
// はコマンドラインからオプションを指定するためによく使われる。
// 例えば `wc -l` の `-l` はコマンドラインフラグである。

package main

// Go の `flag` パッケージを使ってコマンドラインフラグをパースできる。
// 実際にコマンドラインプログラムを作って、このパッケージを使ってみよう。
import "flag"
import "fmt"

func main() {

	// 文字列、整数、真偽値のオプションを受け取るフラグを宣言できる。
	// ここではフラグ `word` を宣言し、そのデフォルト値を `"foo"` とし、フラグの短い説明を与えている。
	// 関数 `flag.String` は（文字列そのものではなく）文字列のポインタを返す。
	// このポインタの使い方は追って説明する。
	wordPtr := flag.String("word", "foo", "a string")

	// `word` と同様のやり方で、フラグ `numb`、`fork` を宣言している。
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// プログラム内の別の場所で宣言した変数を使い、オプションを宣言することもできる。
	// なお、この場合変数のポインタを渡すことに注意する。
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// すべてのフラグを宣言したら、`flag.Parse()` を呼んでコマンドラインをパースする。
	flag.Parse()

	// パースしたオプションと、余りの引数を表示する。
	// `*wordPtr` のように参照を剥がしてオプション値を呼んでいることに注意する。
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
