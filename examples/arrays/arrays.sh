# `fmt.Println` に配列を渡すと `[v1 v2 v3 ...]`
# と表示される。
$ go run arrays.go
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
2d:  [[0 1 2] [1 2 3]]

# Go では普通、配列より<em>スライス</em>をよく見るだろう。
# スライスについては、次の例で見ていく。
