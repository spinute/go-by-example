// [<em>可変長引数関数</em>](http://en.wikipedia.org/wiki/Variadic_function)は、末尾の引数に何個でも変数を受け取れる関数である。
// 例えば、`fmt.Println` はよく使う可変長引数関数である。

package main

import "fmt"

// これは何個でも `int` 型の引数を取る関数である。
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// 可変長引数関数は普通の関数と同様に呼び出せる。
	sum(1, 2)
	sum(1, 2, 3)

	// 複数の引数をスライスとして持っているなら、以下の `func(slice...)` という形式で可変長引数関数にスライスの中身を渡せる。
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
