# `"one"`、`"two"` を順番に受信する。
$ time go run select.go 
received one
received two

# 実行時間の合計は2秒程度である。
# これは1秒間のスリープと2秒間のスリープが平行に実行されるためである。
real	0m2.245s
