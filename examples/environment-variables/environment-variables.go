// [環境変数](http://en.wikipedia.org/wiki/Environment_variable)は
// [設定情報を Unix 上で動くプログラムに伝える](http://www.12factor.net/config)一般的な仕組みである。
// 環境変数を読んだり、書いたりしてみよう。

package main

import "os"
import "strings"
import "fmt"

func main() {

    // キーと値のペアをセットするには `os.Setenv` を使う。
    // キーから値を読み出すには `os.Getenv を使う。
    // もしキーが設定されていなければ空の文字列を返す。
    os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR"))

    // 定義されているキーと値の一覧は `os.Environ` で取得する。
    // この関数は `KEY=value` という形式の文字列からなるスライスを返す。
    // `strings.Split` を使って、このスライスからキーと値を取り出せる。
    // ここではすべてのキーを表示している。
    fmt.Println()
    for _, e := range os.Environ() {
        pair := strings.Split(e, "=")
        fmt.Println(pair[0])
    }

}
