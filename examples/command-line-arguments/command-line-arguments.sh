# コマンドライン引数の実験をするには、`go build` を実行してバイナリを先に作っておくのがいい。
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]       
[a b c d]
c

# 続いては、フラグを使ってより高度なコマンドライン処理の仕方を見ていく。
