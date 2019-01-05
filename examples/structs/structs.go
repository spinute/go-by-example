// Go の構造体は型がついたフィールドの集まりにである。
// 構造体を使うとデータをまとめて扱うのに便利である。

package main

import "fmt"

// 構造体型 `person` はフィールド `name`、`age` を持つ。
type person struct {
    name string
    age  int
}

func main() {

    // 新たな構造体はこう作る。
    fmt.Println(person{"Bob", 20})

    // 構造体を初期化するときに、フィールド名を指定してもよい。
    fmt.Println(person{name: "Alice", age: 30})

    // 明記しないフィールドの値はゼロ値になる。
    fmt.Println(person{name: "Fred"})

    // `&` を前に付けると構造体のポインタを作れる。
    fmt.Println(&person{name: "Ann", age: 40})

    // 構造体のフィールドにアクセスするにはドット演算子を使う。
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    // 構造体のポインタにもドット演算子を使える。
    // このとき、ポインタは自動的に参照を剥がされる。
    sp := &s
    fmt.Println(sp.age)

    // 構造体は可変である。
    sp.age = 51
    fmt.Println(sp.age)
}
