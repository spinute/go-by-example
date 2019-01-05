// Go でファイルに書き込むには、前に見たファイルを読み出すのと似たパターンを使う。

package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // まずは文字列（またはバイト列）を単にファイルに書き込む。
    d1 := []byte("hello\ngo\n")
    err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
    check(err)

    // より細かな制御をするために、ファイルを書き込み用に開く。
    f, err := os.Create("/tmp/dat2")
    check(err)

    // ファイルを開いた直後に `Close` を `defer` するのはイディオムである。
    defer f.Close()

    // `Write` を使ってバイトのスライスを書き込める。
    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    // `WriteString` という関数もあり、こちらは文字列を書き込める。
    n3, err := f.WriteString("writes\n")
    fmt.Printf("wrote %d bytes\n", n3)

    // `Sync` を使って書き込みをストレージにフラッシュする。
    f.Sync()

    // `bufio` を使ってバッファ付きのリーダーを作る例を前に紹介した。
    // 同様に、バッファ付きのライターも作れる。
    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    fmt.Printf("wrote %d bytes\n", n4)

    // `Flush` を使って、ライターのバッファされている操作をすべてフラッシュする。
    w.Flush()

}
