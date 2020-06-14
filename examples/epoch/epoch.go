// Go で [Unix エポック（epoch）](http://en.wikipedia.org/wiki/Unix_time) からの経過時間を秒、ミリ秒、ナノ秒単位を取得する方法を紹介する。

package main

import (
	"fmt"
	"time"
)

func main() {

	// `time.Now` と、`Unix` か `UnixNano` を使って、Unix エポックからの経過時間を秒、ナノ秒単位で取得する。
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// `UnixMillis` はない。
	// そのため、エポックからの経過時間をミリ秒単位で知りたいときはナノ秒の結果を割る必要がある。
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// 逆に、エポックからの秒、ナノ秒での経過時間を`time` に変換することもできる。
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
