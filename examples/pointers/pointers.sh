# `zeroval` は `main` の `i` を書き換えないが、
# `zeroptr` は `i` を書き換える。
# これは後者には変数のメモリアドレスへの参照を渡しているからである。
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100
