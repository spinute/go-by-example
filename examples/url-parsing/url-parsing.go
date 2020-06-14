// URL は[リソースを特定するための統一されたやり方](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/)である。
// この例では、Go で URL をパースするやり方を紹介する。

package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// 以下の URL をパースしてみる。
	// この URL には、スキーマ、認証情報、ホスト、ポート、パス、クエリパラメータ、クエリフラグメントが含まれている。
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// URL をパースし、エラーが発生していないことを確認する。
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// スキーマは直接的に読み出せる。
	fmt.Println(u.Scheme)

	// `User` は認証に必要なすべての情報を含む。
	// `UserName`、`Password` を呼べば個々の値を読み出せる。
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// `Host` はホスト名と、（あれば）ポートとを含む。
	// `SplitHostPort` を使えばそれぞれを抽出できる。
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// ここでは、パスと `#` に続くフラグメントとを読み出している。
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// `RawQuery` を使うと `k=v` という形式の文字列としてクエリパラメータを取得できる。
	// クエリパラメータをマップとしてパースすることもできる。
	// パースして得られるクエリパラメータは、文字列から文字列のスライスへのマップである。
	// そのため、ひとつめの値を読み出したいだけなら `[0]` を使う。
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
