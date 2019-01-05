// <em>switch 文</em>は多くの分岐からなる条件を表す。

package main

import "fmt"
import "time"

func main() {

    // `switch` の基本的な使い方
    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    // コンマを使うと一つの `case` 文に複数の式を書ける。
    // この例では `default` のケースも使っている。
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

    // `switch` の直後に式を書かなければ、それは単なる if/else の別の書き方である。
    // また、定数でないものを使って `case` 式を書ける。
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    // `switch` は値だけでなく型について分岐することもできる。
    // これを使って、インターフェースの型を調べることができる。
    // この例では、変数 `t` は対応する節に書かれた型を持つ。
    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}
