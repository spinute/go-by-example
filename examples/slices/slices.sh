# スライスの型と配列の型は違うが、`fmt.Println` はいずれの型の値も同様に表示する。
$ go run slices.go
emp: [  ]
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
2d:  [[0] [1 2] [2 3 4]]

# スライスの設計と実装に関する詳細が知りたければ、Go の開発チームが書いた[素晴らしいブログポスト](http://blog.golang.org/2011/01/go-slices-usage-and-internals.html)を見てほしい。

# これまで配列とスライスを見てきたが、Go の重要な組み込みデータ型として次はマップを紹介する。
