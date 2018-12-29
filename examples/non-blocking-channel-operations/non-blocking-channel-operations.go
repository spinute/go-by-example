// 基本的なチャネルの送受信の使い方はブロックする。
// しかし、`select` と `default` 節を使えば、_ノンブロッキング_な送受信や、選択的な受信を実装できる。

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// これはノンブロッキングな受信である。
	// もし `messages` から値を受け取れるなら、`select` は `<-messages` の `case` 節で、その値を受け取る。
	// そうでなければ、`default` 節に直ちに進む。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// ノンブロッキングな送信も同様だ。
	// ここでは `msg` を `messages` チャネルに送信できない。
	// なぜならチャネルにバッファは付いておらず、受信者もいないからである。
	// そのため、`default` 節に進む。
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// `default` 節の前に複数の `case` を書けば、複数のチャネルから選択的かつノンブロッキングに受信できる。
	// ここでは `messages` と`signals` からノンブロッキング受信を試みている。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
