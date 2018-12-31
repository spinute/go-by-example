// [mutex](mutexes) を使い明示的なロックを取って、複数のゴルーチンで共有データへのアクセスを同期する方法は既に紹介した。
// ゴルーチンを同期する組み込みの機能とチャネルを使っても、同じ結果が得られる。
// チャネルを使うこのやり方は、各ゴルーチンが持つデータをやり取りしてメモリを共有する Go のアイデアに合っている。

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// この例では、あるゴルーチンだけが状態を所有する。
// この結果、アクセスの競合によってデータが壊れてしまうことなくなる。
// 状態を読み書きするために、他のゴルーチンは状態を持つゴルーチンにメッセージを送り、返信を受け取る。
// 構造体 `readOp`、`writeOp` はこれらのリクエストをカプセル化したものだ。
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// 前と同じように、操作を実行した回数を数える。
	var readOps uint64
	var writeOps uint64

	// チャネル `reads`、`writes` を使って、他のゴルーチンは読み書きのリクエストする。
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// このゴルーチンが `state` を所有する。
	// `state` は前の例と同様のマップだが、この状態管理用のゴルーチンだけが読み書きをする。
	// このゴルーチンはチャネル `reads`、`writes` を繰り返し `select` し、届いたリクエストに返信する。
	// リクエストを受け取ると、リスエストされた操作を実行し、値を返信用のチャネル `resp` に送信し、リクエストが成功したことを伝える（`reads` の場合は読み出した結果も伝える）。
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// ここで100個のゴルーチンを開始し、チャネル `reads` を通じて状態を管理するゴルーチンに読み出しを発行する。
	// 読み出しリクエストのたびに `readOp` を作り、チャネル `reads` にそれを送り、チャネル `resp` から結果を受け取る。
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 同様に、書き込みを行うゴルーチンも10個開始する。
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 1秒間待って、ゴルーチンに仕事をさせる。
	time.Sleep(time.Second)

	// 最後に操作回数を読み出し、表示する。
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
