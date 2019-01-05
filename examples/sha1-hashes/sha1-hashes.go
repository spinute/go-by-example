// [<em>SHA1 ハッシュ関数</em>](http://en.wikipedia.org/wiki/SHA-1)はバイナリや文字列の短い識別子を計算するためによく使う。
// 例えばバージョン管理システムである [Git](http://git-scm.com/) ではファイルやディレクトリを識別するために SHA1 を使っている。
// それでは、Go で SHA1 を計算する例を見ていこう。

package main

// `crypto/*` パッケージには何種類かのハッシュ関数が含まれている。
import "crypto/sha1"
import "fmt"

func main() {
    s := "sha1 this string"

    // ハッシュ値を生成するには `sha1.New()`、`sha1.Write(bytes)`、`sha1.Sum([]byte{})` の順で関数を呼ぶ。
    // まずは新たなハッシュ関数を生成する。
    h := sha1.New()

    // `Write` の入力はバイト列である。
    // 文字列 `s` のハッシュ値を計算したいなら、`[]byte(s)` と書いてハッシュ値に変換してやらなければならない。
    h.Write([]byte(s))

    // バイトのスライスとして、最終的なハッシュ値を得る。
    // `Sum` の引数を、これまで入力したバイト列に追記できるが、普通はこれは使わない。
    bs := h.Sum(nil)

    // SHA1 のハッシュ値は、Git がそうしているように、16進記数法で表示することが多い。
    // フォーマット文字列に `%x` と書けば、ハッシュ計算の結果を16進文字列に変換できる。
    fmt.Println(s)
    fmt.Printf("%x\n", bs)
}
