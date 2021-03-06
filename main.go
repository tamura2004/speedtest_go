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
	csv      string //速度記録ファイル名
}

// 起動オプションの読み込み
func (p *parm) init() {
	flag.BoolVar(&p.isServer, "s", false, "サーバとして起動")
	flag.BoolVar(&p.isClient, "c", false, "クライアントとして起動")
	flag.StringVar(&p.host, "h", "", "ホスト名")
	flag.StringVar(&p.port, "p", "7777", "ポート番号")
	flag.IntVar(&p.length, "l", 500, "転送データサイズ[Kbyte]")
	flag.IntVar(&p.repeat, "r", 3, "繰り返し回数")
	flag.IntVar(&p.wait, "w", 3, "測定間隔[秒]")
	flag.StringVar(&p.file, "f", "text.log", "ログファイル名")
	flag.StringVar(&p.csv, "v", "record.csv", "ログファイル名")
	flag.Parse()
}

func main() {
	// 起動オプションの初期化
	var p parm
	p.init()

	// 起動オプションによりサーバまたはクライアントとして起動
	if p.isServer {
		p.server()
	} else if p.isClient {
		p.client()
	} else {
		panic("-s オプションか -c オプションを指定して下さい")
	}
}
