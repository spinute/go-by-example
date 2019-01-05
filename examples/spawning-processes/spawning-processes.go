// Go のプログラムから Go ではない外部プロセスを spawn したくなることがある。
// 例えば、このサイトのシンタックスハイライトは Go で書いたプログラムから [`pygments`](http://pygments.org/) を実行するプロセスを spawn [している](https://github.com/mmcgrana/gobyexample/blob/master/tools/generate.go)。
// Go からプロセスを spawn する例を他にも見てみよう。

package main

import "fmt"
import "io/ioutil"
import "os/exec"

func main() {

    // まずは引数も入力も取らず、標準出力に何かを書くだけの単純なコマンドから始めよう。
    // ヘルパーコマンド `exec.Command` は外部プロセスを表すオブジェクトを作る。
    dateCmd := exec.Command("date")

    // 他に `.Output` というヘルパーもあり、これはコマンドを実行し、その終了を待ちながら、出力を集める。
    // エラーが起きなければ、`dateOut` は日時情報を表すバイト列を保持する。
    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> date")
    fmt.Println(string(dateOut))

    // 続いて、パイプを使って外部プロセスの標準入力にデータを渡し、標準出力から結果を収集する方法を見てみる。
    grepCmd := exec.Command("grep", "hello")

    // ここでは明示的に入力・出力パイプを得て、プロセスを開始し、入力を書き込み、結果を読み出し、プロセスが終了するのを待っている。
    grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()
    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, _ := ioutil.ReadAll(grepOut)
    grepCmd.Wait()

    // 上の例ではエラーチェックを省いたが、いつものように `if err != nil` のパターンでエラーチェックをしてもよい。
    // また、`StdoutPipe` だけを見ていたが、`StderrPipe` も全く同じやり方で読み出せる。
    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))

    // コマンドを spawn するときは、コマンドラインに打ち込む文字列を丸ごと渡すのではなく、コマンドとその引数の配列を明示的に並べなければならない。
    // もしコマンド全体を単に文字列で書きたいなら、`bash` の `-c` オプションを使えばよい。
    lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))
}
