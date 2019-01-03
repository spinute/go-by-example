# コマンドラインフラグの実験をするときは、まずプログラムをコンパイルし、その後バイナリを直接実行するとよい。
$ go build command-line-flags.go

# ビルドしたプログラムを、すべてのフラグに値を与えて実行する。
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# フラグを書かなければ、自動的にデフォルト値を使うことに注意する。
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# フラグの後に引数を並べられる。
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# `flag` パッケージでは、すべてのフラグは単なる引数の前に書かなければならない。
# （さもなければ、フラグはただの引数として解釈されてしまう。）
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# `-h` か `--help` を使うと、自動生成したコマンドラインのヘルプを表示する。
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# `flag` パッケージで指定していないフラグをプログラムに渡すと、プログラムはエラーメッセージを表示し、この場合もヘルプを表示する。
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...

# 続いて、プログラムにパラメータを渡す別の方法である、環境変数を紹介する。
