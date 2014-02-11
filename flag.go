// flag
package main

import (
	"flag"
)

// 起動時オプション
type parm struct {
	isServer bool   //サーバとして起動
	isClient bool   //クライアントとして起動
	host     string //ホスト名
	port     string //ポート番号
	length   int    //転送データサイズ[Kbyte]
	repeat   int    //測定回数[回]
	wait     int    //測定間隔[秒]
	file     string //ログファイル名
}

// 起動オプションの読み込み
func (p *parm) init() {
	flag.BoolVar(&p.isServer, "s", false, "サーバとして起動")
	flag.BoolVar(&p.isClient, "c", false, "クライアントとして起動")
	flag.StringVar(&p.host, "h", "localhost", "ホスト名")
	flag.StringVar(&p.port, "p", "7777", "ポート番号")
	flag.IntVar(&p.length, "l", 500, "転送データサイズ[Kbyte]")
	flag.IntVar(&p.repeat, "r", 3, "繰り返し回数")
	flag.IntVar(&p.wait, "w", 10, "測定間隔[秒]")
	flag.StringVar(&p.file, "f", "text.log", "ログファイル名")
	flag.Parse()
}
