// Go は組み込みで [base64](http://en.wikipedia.org/wiki/Base64) という
// エンコード形式をサポートしている。

package main

// このように書くと、`encoding/base64` パッケージを（デフォルトの base64 の代わりに）`b64` という名前でインポートする。
// これはスペースを節約するのに役立つ。
import b64 "encoding/base64"
import "fmt"

func main() {

    // 今からエンコード・デコードする文字列
    data := "abc123!?$*&()'-=@~"

    // Go は標準の base64 と、URL 文字列として使える base64 をいずれもサポートしている。
    // ここでは標準の方を使う。
    // エンコーダには `[]byte` 型の値を渡すので、ここでは `string` 型からキャストしている。
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

    // 入力によってはデコードは失敗するかもしれない。
    // その可能性があるなら、返り値からエラーの有無を確認できる。
    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

    // 根では URL として使える base64 フォーマットを使ってみる。
    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}
