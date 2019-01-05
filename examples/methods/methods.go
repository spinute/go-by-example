// Go は構造体型に対して定義するメソッドをサポートしている。

package main

import "fmt"

type rect struct {
    width, height int
}

// このメソッド `area` の<em>レシーバ</em>の型は `*rect` である。
func (r *rect) area() int {
    return r.width * r.height
}

// メソッドはレシーバのポインタか値に対して定義できる。
// これはレシーバの値に対してメソッドを定義する例である。
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    // 構造体型に定義した2つのメソッドを呼んでみる。
    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    // Go はメソッドを呼ぶとき、メソッド定義に従って、値とポインタの変換を自動で行う。
    // レシーバのポインタを使ってメソッドを呼ぶと、メソッド呼び出し時のコピーを避けたり、メソッドに値を書き換えさせることができる。
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}
