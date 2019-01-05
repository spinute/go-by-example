// ファイルの読み書きは Go のプログラムで頻繁に必要になる。
// まずは、ファイルを読む例を示す。

package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

// ファイルを読むときは多くの関数でエラーチェックが必要になる。
// このヘルパーでエラーチェックを簡単に済ます。
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // 最も基本的なファイル読み出しは、ファイル全体をメモリに読み出すことだろう。
    dat, err := ioutil.ReadFile("/tmp/dat")
    check(err)
    fmt.Print(string(dat))

    // ファイルのどの部分をどのように読むかを細かく制御したいこともあるだろう。
    // その場合、まずはファイルを `Open` して、`os.File` の値を得る。
    f, err := os.Open("/tmp/dat")
    check(err)

    // ファイルの先頭から何バイトか読んでみる。
    // ここでは最大5バイトまで読み出すようにしているが、実際に何バイト読んだかも表示している。
    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1))

    // `Seek` して既知のある箇所から `read` することもできる。
    o2, err := f.Seek(6, 0)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

    // `io` パッケージにはファイルを読むのに便利な関数が入っている。
    // 例えば、上の読み出しの例は、`ReadAtLeast` を使ってより頑強に実装できる。
    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    // 組み込みの巻き戻し（rewind）関数は無いが、`Seek(0, 0)` で代替できる。
    _, err = f.Seek(0, 0)
    check(err)

    // `bufio` パッケージにはバッファ付きの reader があり、小さな読み出しを効率よく行えたり、追加の読み出し方法があったりして役立つことがある。
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))

    // やることが終わったらファイルを閉じること（普通は `Open` の直後に `defer` を使ってファイルを確実に閉じるのがよい）。
    f.Close()

}
