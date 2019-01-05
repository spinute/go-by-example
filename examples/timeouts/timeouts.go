// 外部のリソースに接続したり、実行時間に上限を設けたりするプログラムには<em>タイムアウト</em>が必要になる。
// Go ではチャネルと `select` を使ってタイムアウトをうまく実現できる。

package main

import "time"
import "fmt"

func main() {

    // チャネル `c1` に2秒後に結果を返す外部機能を呼び出したいとする。
    c1 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
    }()

    // `select` を使ってタイムアウトを実装する。
    // `res := <-c1` は結果を待ち受け、<-Time.After` は1秒のタイムアウト後に送られる値を受け取る。
    // `select` は最初に準備できた値を受信するので、操作が許容時間1秒よりも長い時間を使うと、タイムアウトのケースに進む。
    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout 1")
    }

    // もっと長いタイムアウトとして3秒を使えば、`c2` からの受信は成功し、結果を表示する。
    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2")
    }
}
