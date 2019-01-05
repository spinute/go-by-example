// Go の _select_ を使うと、複数のチャネルで受信を待てる。
// ゴルーチンとチャネル、そして select の組み合わせは Go の強力な機能である。

package main

import "time"
import "fmt"

func main() {

    // この例では2つのチャネルに select を使う。
    c1 := make(chan string)
    c2 := make(chan string)

    // いずれのチャネルも一定の時間が経てば値を受信する。
    // これはゴルーチンが RPC を同期的に実行する状況を模している。
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    // We'll use `select` to await both of these values
    // simultaneously, printing each one as it arrives.
    // `select` を使ってふたつの値をいずれも非同期で受信し、届いた方から表示する。
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
