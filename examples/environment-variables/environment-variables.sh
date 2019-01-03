# プログラムを実行すると、プログラム内で設定した `FOO` の値を読み出せたことがわかる。
# 一方、`BAR` の値が空であることもわかる。
$ go run environment-variables.go
FOO: 1
BAR: 

# 環境におけるキーの一覧は、実行するマシンによって違うはずだ。
TERM_PROGRAM
PATH
SHELL
...

# `BAR` の値を設定してからプログラムを実行すれば、プログラムはその値を読み出せる。
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...
