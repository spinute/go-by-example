// 前の例で、単純なカウンタの状態を[アトミックな操作](atomic-counters.html)で管理する方法を紹介した。
// より複雑な状態を管理したいときには、[<em>ミューテックス（mutex）</em>](http://en.wikipedia.org/wiki/Mutual_exclusion)を使って複数のゴルーチンから安全にデータを読み書きすることもできる。

package main

import (
    "fmt"
    "math/rand"
    "sync"
    "sync/atomic"
    "time"
)

func main() {

    // 状態の例としてマップを使う。
    var state = make(map[int]int)

    // この `mutex` を使って、`state` へのアクセスを同期する。
    var mutex = &sync.Mutex{}

    // 読み操作・書き操作がそれぞれ何度行われたかをカウントする。
    var readOps uint64
    var writeOps uint64

    // 100 個のゴルーチンを作り、それぞれが1ミリ秒ごとに state を繰り返し読み出す。
    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {

                // 読み出しのたびに、アクセスするキーを選ぶ。
                // `Lock()` で `mutex` をロックし、`state` への排他的なアクセスを獲得し、値を読み出し、`Unlock()` で `mutex` のロックを解除し、最後に `readOps` の値を1増やす。
                key := rand.Intn(5)
                mutex.Lock()
                total += state[key]
                mutex.Unlock()
                atomic.AddUint64(&readOps, 1)

                // 読み出しのたびに少し待つ。
                time.Sleep(time.Millisecond)
            }
        }()
    }

    // 10個のゴルーチンを起動し、読み出しと同様のパターンで書き込みを行う。
    for w := 0; w < 10; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    // 10個のゴルイーチンが `state` や `mutex` で仕事をできるよう、一秒待つ。
    time.Sleep(time.Second)

    // 結果を読み出して、読み書き操作のカウントを報告する
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)

    // `state` をロックして、最終的にどうなったかを表示する。
    mutex.Lock()
    fmt.Println("state:", state)
    mutex.Unlock()
}
