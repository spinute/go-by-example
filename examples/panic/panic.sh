# プログラムを実行するとパニックが発生し、
# エラーメッセージとゴルーチンのトレースが表示され、
# 非ゼロのステータスでプログラムは終了する。
$ go run panic.go
panic: a problem

goroutine 1 [running]:
main.main()
	/.../panic.go:12 +0x47
...
exit status 2

# 例外を使って多くのエラーを処理する言語とは違って、
# Go ではエラーの有無を返り値として返すのが普通であることに注意する。
