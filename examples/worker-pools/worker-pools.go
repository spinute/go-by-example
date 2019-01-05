// この例ではゴルーチンとチャネルを使って<em>ワーカープール</em>を実装するやり方を紹介する。

package main

import "fmt"
import "time"

// このワーカーをいくつか平行に動かす。
// 各ワーカーは `jobs` チャネルに仕事を受信し、その結果を `results` チャネルに送信する。
// 実行に時間のかかる処理をシミュレートするために、仕事をするたびに1秒スリープする。
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func main() {

    // ワーカープールを使うために、仕事を送り、結果を収集する必要がある。
    // そのために使う2つのチャネルを作る。
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // ここで3つワーカーを起動する。
    // まだジョブがないので、いずれのワーカーもブロックする。
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // `jobs` に値を5つ送り、その後チャネルを閉じて仕事を送り終えたことを伝える。
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    // 最後に仕事の結果を集める。
    for a := 1; a <= 5; a++ {
        <-results
    }
}
