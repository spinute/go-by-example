// Go で状態を管理するとき、まず検討すべきはチャネル通信を使う方法である。
// このような例は既に<a href="./worker-pools" type="text/html">ワーカープール</a>で紹介した。
//
// 別のやり方もある。
// ここでは、`sync/atomic` パッケージにある、複数のゴルーチンからアクセスされる<em>アトミックなカウンタ</em>を使う方法を紹介する。

package main

import "fmt"
import "time"
import "sync/atomic"

func main() {

	// 符号無し整数型の変数で（負の値を取らない）カウンタを表現する
	var ops uint64

	// 更新を並行に実行するため、50個のゴルーチンを起動する。
	// 各ゴルーチンは約1ミリ秒ごとにカウンタの値を1増やす。
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// カウンタを増やすには`AddUint64` を使う。
				// この関数にはカウンタ `ops` のメモリアドレスを渡す。
				// 演算子 `&` を使って変数のアドレスを取得できる。
				atomic.AddUint64(&ops, 1)

				// カウンタを増やすたびに少し待つ。
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 1秒待って、ゴルーチンに仕事をさせる。
	time.Sleep(time.Second)

	// ゴルーチンはまだカウンタを更新し続けている。
	// 安全にカウンタから値を読み出すために `LoadUint64` を使う。
	// 先程と同じように、この関数にはカウンタのメモリアドレス `&ops` を渡す。
	// この関数は安全にカウンタの値を読み出す。
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
