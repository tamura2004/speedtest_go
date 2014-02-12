// client
package main

import (
	"fmt"
	//"io/ioutil"
	"net"
	"time"
)

// クライアント計測処理
func (p *parm) client() {

	// 指定回数繰り返し
	for i := 0; i < p.repeat; i++ {

		// サーバ接続処理
		p.connect()

		// 指定時間待ち
		time.Sleep(time.Duration(p.wait) * time.Second)
	}
}

//// サーバ接続および転送速度計測
func (p *parm) connect() {

	// 接続アドレス生成
	addr, err := net.ResolveUDPAddr("udp", p.host+":"+p.port)
	p.handle(err, "接続アドレス生成")

	// ソケット取得
	conn, err := net.DialUDP("udp", nil, addr)
	p.handle(err, "ソケット取得")
	defer conn.Close()

	// 開始ログ
	p.log("start connecting")

	// 転送サイズ
	kbyte := float64(p.length)

	// 転送レートを計測[秒]
	buf := make([]byte, 1024)
	sec := p.benchmark(func() {
		// 転送サイズ指定電文送信
		_, err := conn.Write([]byte(fmt.Sprintf("%d", p.length)))
		p.handle(err, "転送サイズ指定電文送信")

		// 終了電文受信まで1KBづつ転送
		for {
			n, err := conn.Read(buf)
			p.handle(err, "1KB受信")
			if string(buf[:n]) == "end" {
				break
			}
		}
	})

	// 転送レート計算
	kbps := kbyte * 8.0 / sec

	// 転送サイズ、レート記録
	p.log(fmt.Sprintf("transrate: %.2fKbyte, time: %.2fsec, speed: %.2fKbps", kbyte, sec, kbps))

}
