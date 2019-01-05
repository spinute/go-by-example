// <em>行フィルタ</em>はよくある種類のプログラムで、標準入力を読んで、それを処理し、結果を標準出力に書き出す。
// `grep` や `sed` はよく使われる行フィルタの例である。

// ここでは Go で行フィルタを実装してみる。
// このプログラムは入力テキストをいずれも大文字にしたものを出力する。
// このパターンを真似れば、自分で好きな行フィルタを実装できるだろう。
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {

    // バッファの付いていない `os.Stdin` をバッファ付きのスキャナでラップすると、便利な `Scan` メソッドを使ってトークンごとに入力を読み進められる。
    // デフォルトのスキャナは Stdin を行ごとに読み出すのだ。
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        // `Text` は現在のトークンを返す。
        // 今の場合は、入力から読み出した次の行である。
        ucl := strings.ToUpper(scanner.Text())

        // 大文字にした行を書き出す。
        fmt.Println(ucl)
    }

    // `Scan` でエラーが無かったか確認する。
    // なお、終端記号（EOF）が見つかっても、`Scan` はエラーとして扱わない。
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
