// Go での慣習として、エラーを独立した返り値として明示的に返す。
// これは例外を使う Java や Ruby、結果とエラーを一つの値で返す C などの言語とは対照的である。
// Go のやり方だとどの関数がエラーを返すかがわかりやすく、またエラーのために特別な言語機能を増やす必要がない。

package main

import "errors"
import "fmt"

// 慣習的に、最後の返り値を組み込みインターフェースである `error` 型の値にすることで、エラーの有無を返す。
func f1(arg int) (int, error) {
	if arg == 42 {

		// `errors.New` は基本的な `error` 型の値を作る関数である。引数はエラーメッセージを表す。
		return -1, errors.New("can't work with 42")

	}

	// error の値 `nil` はエラーが無かったことを表す。
	return arg + 3, nil
}

// `Error()` メソッドを実装すれば、自分で作った型を `error` として使える。
// エラーが引数エラーであることを明示する型を作り、先の例と同様の例を書いてみる。
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		// `&argError` と書いて、フィールド `arg` と `prob` の値を渡しながら構造体を作っている。
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// 以下の2つのループでは、上で実装したエラーを返す関数を試している。
	// `if` の行でエラーの有無を確認するのは、Go の頻出パターンである。
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// 自作したエラーのデータをプログラム中で使うときには、型アサーションを使って自作したエラー型のインスタンスを作る必要がある。
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
