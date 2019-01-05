# プログラムを起動して2秒経つと1つ目のタイマーは切れる。
# しかし、2つ目のタイマーは2秒立つ前に停止するので、時間切れにならない。
$ go run timers.go
Timer 1 expired
Timer 2 stopped
