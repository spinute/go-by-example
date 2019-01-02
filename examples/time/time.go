// Go には時刻や経過時間を扱うための機能が多くある。
// ここではその例を示す。

package main

import "fmt"
import "time"

func main() {
	p := fmt.Println

	// まずは現在時刻を取得する。
	now := time.Now()
	p(now)

	// `time` 構造体を、年・月・日などを指定して作ることもできる。
	// 時刻は常に位置（`Location`）、すなわちタイムゾーンと結びついている。
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// 時刻の構成要素の取り出しかたは直感的だ。
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// 曜日（Monday-Sunday）を表す `Weekday` もある。
	p(then.Weekday())

	// 以下のメソッドは2つの時刻を比べて、1つ目が2つ目より前か、後か、それとも同時刻かをそれぞれテストする。
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// メソッド `Sub` は2つの時刻の間隔を表す経過時間 `Duration` を返す。
	diff := now.Sub(then)
	p(diff)

	// 経過時間の長さの単位を変換できる。
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// `Add` を使って指定した期間だけ時刻を進めることもできる。`-` を使えば、時刻を戻すこともできる。
	p(then.Add(diff))
	p(then.Add(-diff))
}
