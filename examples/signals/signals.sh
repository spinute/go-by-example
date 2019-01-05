# このプログラムを実行すると、シグナルを送るまで処理がブロックされる。
# `ctrl-C`（ターミナルはこれを `^C`と表示する）を入力すると、シグナル `SIGINT` を送れる。
# このシグナルを受け取るとプログラムは `interrupt` を表示し、終了する。
$ go run signals.go
awaiting signal
^C
interrupt
exiting
