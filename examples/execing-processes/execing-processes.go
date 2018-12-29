// 前の例で[外部プロセスを作る](spawning-processes)方法を紹介した。
// 実行中の Go プロセスから外部プロセスにアクセスするときはこの機能を使う。
// しかし、Go のプロセスを別の（Go ではないかもしれない）プロセスに置き換えたい場面もある。
// このようなときは昔からある <a href="http://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a> の機能を、Go で実装したものを使う。

package main

import "syscall"
import "os"
import "os/exec"

func main() {

	// この例では `ls` を exec する。
	// 実行したいバイナリの絶対パスを Go は必要とするので、 // `exec.LookPath` を使う（おそらく `/bin/ls` だろう）。
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// `Exec` の引数はスライスで表現する（ひとつの大きな文字列ではない）。
	// ここでは `ls` によく使う引数を渡してみる。
	// なお、最初の引数はプログラム名であることに注意する。
	args := []string{"ls", "-a", "-l", "-h"}

	// `Exec` には[環境変数](environment-variables)も渡す必要がある。
	// ここでは、現在の環境変数をそのまま渡す。
	env := os.Environ()

	// `syscall.Exec` を呼ぶ。
	// 呼び出しが成功すると、このプロセスはここで終わり、 `/bin/ls -a -l -h` を実行するプロセスに置き換わる。
	// もしエラーがあれば、それを表す値が返ってくる。
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
