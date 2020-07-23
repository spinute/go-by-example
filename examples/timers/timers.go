// Go のコードを将来のある時点で実行したり、定期的に繰り返し実行したりすることがよくある。
// 組み込みの機能 _timer_ 、 _ticker_ を使ってこれらの機能をシンプルに実現できる。
// まずは timer を紹介し、その後で [ticker](tickers.html) を紹介する。

package main

import (
	"fmt"
	"time"
)

func main() {

	// タイマーは将来のあるイベントを表す。
	// タイマーは待つ期間を指定して作り、そのときに得られるチャネルには、指定した期間が経つと値が届く。
	// ここで作っているタイマーは2秒待つものだ。
	timer1 := time.NewTimer(2 * time.Second)

	// `<-timer1.C` は、タイマーのチャネル `C` にタイマーが切れたことを示す値が届くまでブロックする。
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// ただ待ちたいだけなら、`time.Sleep` を使ってもよかったかもしれない。
	// タイマーを使った方が便利な点としては、タイマーは切れる前にキャンセルすることもできる。
	// ここではその例を紹介する。
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// タイマーがキャンセルされたことを確認するため、`timer2` が（キャンセルされていなければ）発火するのに十分な時間待つ。
	time.Sleep(2 * time.Second)
}
