// <em>インターフェース</em>とは、メソッドのシグネチャの集まりに名前を付けたものである。

package main

import "fmt"
import "math"

// これは幾何学図形のインターフェースである
type geometry interface {
    area() float64
    perim() float64
}

// `rect` 型と `circle` 型で、このインターフェースを実装してみせる。
type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

// Go のインターフェースを実装するには、インターフェースに含まれるメソッドをすべて実装すればよい。
// ここでは `rect` 型で `geometry` インターフェースを実装してみせる。
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

// `circle` 型についても `geometry` を実装する。
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// 変数がインターフェース型を持つなら、そのインターフェースのメソッドを呼ぶことができる。
// この `measure` 関数は汎用である。すなわち `geometry` 型の変数であればこの関数を使える。
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    // 構造体型 `circle` と `rect` はいずれも `geometry` インターフェースを実装している。
    // そのためこれらの型のインスタンスを `measure` の引数にすることができる。
    measure(r)
    measure(c)
}
