# 文字列をエンコードした結果は、標準の形式と、URL 文字列として扱える形式とで微妙に違う（`+`と`-`など）。
# しかし、いずれの形式を使っても、デコードするともとの文字列に戻る。
$ go run base64-encoding.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
