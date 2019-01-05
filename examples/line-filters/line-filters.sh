# 実装したラインフィルタを使ってみるため、小文字で書いた行をいくつか含むファイルを作る。
$ echo 'hello'   > /tmp/lines
$ echo 'filter' >> /tmp/lines

# ラインフィルタによって、各行は大文字になる。
$ cat /tmp/lines | go run line-filters.go
HELLO
FILTER
