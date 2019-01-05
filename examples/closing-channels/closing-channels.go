// チャネルを<em>閉じる</em>とそのチャネルにはもう値を送らないことを示せる。
// これはチャネルの受信者に終わりを伝えるのに使える。

package main

import "fmt"

// この例ではチャネル `jobs` を使って `main()` ゴルーチンからワーカーゴルーチンに仕事を送る。
// ワーカーに実行してもらう仕事を送り終えると、チャネル `jobs` を閉じる。
func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    // ワーカーを起動する。
    // このゴルーチンは `j, more := <-jobs` の行で繰り返し `jobs` から値を受信する。
    // チャネルから2つ値を読み出しているのがキモだ。
    // チャネルが閉じられており、そのチャネルの値をすべて受信し終えているとき、`more` は false になる。
    // こうして最後の仕事を終えたことがわかるので、`done` を送り返せる。
    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    // チャネル `jobs` を通じて3つの仕事をワーカーに送り、その後チャネルを閉じる。
    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    // 既に紹介した[同期](channel-synchronization)のやり方で、ワーカーを待つ。
    <-done
}
