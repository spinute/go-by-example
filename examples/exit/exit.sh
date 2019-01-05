# `exit.go` を `go run` で実行すると、プログラム終了は `go ` に補足され、画面に表示される。
$ go run exit.go
exit status 3

# バイナリをビルドしてから実行する場合は、ターミナル上でステータスを確認できる。
$ go build exit.go
$ ./exit
$ echo $?
3

# なお、プログラムによって `!` が表示されることはない。
