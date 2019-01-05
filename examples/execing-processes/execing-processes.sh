# プログラムを実行すると、そのプログラムは `ls` に置き換わる。
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# Go は Unix の古典的な `fork` 関数を提供していないことに注意する。
# しかし、普通はそれで問題はない。
# なぜなら、ゴルーチンを起動したり、プロセスを spawn、exec したりできれば、
# `fork` のユースケースの大部分をカバーできるからである。
