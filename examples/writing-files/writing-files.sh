# ファイルに書き込みをするプログラムを実行してみる。
$ go run writing-files.go 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

# そして、書き込み対象のファイルの内容を確認する。
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered

# 続いて、ここまで見てきたファイル IO のアイデアを、ストリームである `stdin` と `stdout` に適用する例を紹介する。
