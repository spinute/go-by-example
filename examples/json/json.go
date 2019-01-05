// Go には組み込みの JSON サポートがあり、組み込みの型でも自作の型でも、エンコード・デコードが可能である。

package main

import "encoding/json"
import "fmt"
import "os"

// この2つの構造体を使って、自作した型をエンコード・デコードしてみせる。
type response1 struct {
    Page   int
    Fruits []string
}
type response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {

    // まず、基本的なデータ型を JSON 文字列にしてみる。
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))

    // スライス、マップの例である。
    // これらの型の値は、想像に難くないように JSON の配列、オブジェクトにエンコードされる。
    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    // JSON パッケージは自作のデータ型も自動でエンコードしてくれる。
    // エンコード結果には、エクスポートされているフィールドだけを含む。
    // また、デフォルトではフィールド名を JSON のキーにする。
    res1D := &response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    // 構造体のフィールド定義にタグを含めれば、JSON キーの名前を設定することもできる。
    // 上の方にある `response2` の定義を見て、どのようにタグを書いているか確認しよう。
    res2D := &response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

    // 今度は JSON データを Go の値にデコードする方法を見ていく。
    // これは一般的なデータ構造の例である。
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

    // JSON パッケージがデコードしたデータを入れる変数が必要だ。
    // この `map[string]interface{}` 型は文字列型から任意の型へのマップである。
    var dat map[string]interface{}

    // ここでデコードを試み、エラーをチェックしている。
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

    // デコードされたマップの値を使うためには、適切な値にキャストする必要がある。
    // 例えば `num` の値は `float64` 型にキャストすれば使える。
    num := dat["num"].(float64)
    fmt.Println(num)

    // 入れ子になったデータにアクセスするには、キャストを繰り返す必要がある。
    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1)

    // JSON を自作のデータ型としてデコードすることもできる。
    // このやり方には型安全性を高め、型アサーションを書かずにデータにアクセスできる利点がある。
    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

    // ここまでの例では、データを JSON の表現にして標準出力に出すまでの中間表現として、バイト型と文字列型を使ってきた。
    // JSON エンコードされたストリームを、`os.Writer' のように `os.Stdout' や HTTP レスポンスのボディ部として直接書き込むこともできる。
    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}
