# プログラムはハッシュ値を計算し、人が読みやすい16進記数法で結果を表示する。
$ go run sha1-hashes.go
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94


# 上と同様のやり方で他のハッシュ関数を計算することもできる。
# 例えば MD5 ハッシュ関数を使いたければ、
# `crypto/md5` をインポートして `md5.New()` を使えばよい。

# ただし、暗号学的に安全なハッシュ関数を使いたい場合には
# [ハッシュ関数の強度](http://en.wikipedia.org/wiki/Cryptographic_hash_function)をよく読むべきだ。
