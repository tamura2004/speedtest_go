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
	fh, err := os.OpenFile(p.file, os.O_APPEND|os.O_CREATE, 0666)
	p.handle(err, "")
	defer fh.Close()

	// 時刻
	t := time.Now().Format("2006/01/02 15:04:05 ")

	// ログファイルとコンソールにメッセージ書き込み
	fmt.Fprintln(fh, t+msg)
	fmt.Println(t + msg)
}

// 速度記録
func (p *parm) record(kbps float64) {
	fh, err := os.OpenFile(p.csv, os.O_APPEND|os.O_CREATE, 0666)
	p.handle(err, "")
	defer fh.Close()

	// 時刻
	t := time.Now().Format("2006/01/02,15:04:05,")

	//　記録
	fmt.Fprintf(fh, "%s%.2f\n", t, kbps)
}
