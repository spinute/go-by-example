# このプログラムを実行すると、ブロックする呼び出しの出力がまず表示され、
# その後2つのゴルーチンの出力が入り混じって表示される。
# こうなるのは Go のランタイムがゴルーチンを平行に実行しているためである。
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done

# 次は、Go における並行プログラミングをする際にゴルーチンとしばしば一緒に使われるチャネルを紹介する。
