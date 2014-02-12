// log
package main

import (
	"fmt"
	"os"
	"time"
)

// 共通エラー処理(手抜き)
func (p *parm) handle(err interface{}, msg string) {
	if err != nil {
		panic(err)
		//} else if msg != "" {
		//	p.log(msg)
	}
}

// ログ出力処理
func (p *parm) log(msg string) {

	// ファイルの準備
	fh := openOrCreate(p.file)
	defer fh.Close()

	// 時刻
	t := time.Now().Format("2006/01/02 15:04:05 ")

	// ログファイルとコンソールにメッセージ書き込み
	fmt.Fprintln(fh, t+msg)
	fmt.Println(t + msg)
}

// ファイルがあれば追記モードで開く。なければ新規作成。
func openOrCreate(file string) *os.File {

	// ファイルの準備
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		fh, err := os.Create(file)
		if err != nil {
			panic(err)
		}
		return fh
	} else {
		fh, err := os.OpenFile(file, os.O_APPEND, 0777)
		if err != nil {
			panic(err)
		}
		return fh
	}
}
