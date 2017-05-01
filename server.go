package main

import (
	"fmt"
	"net"
	"strconv"
)

// udpサーバ処理
func (p *parm) server() {

	// 開始メッセージ
	p.log(fmt.Sprintf("start server. listening udp port %s.", p.port))

	// アドレス設定
	addr, err := net.ResolveUDPAddr("udp", p.host+":"+p.port)
	p.handle(err, "アドレス設定")

	// 接続待ち設定
	conn, err := net.ListenUDP("udp", addr)
	p.handle(err, "接続待ち設定")
	defer conn.Close()

	// サーバ読み込み
	buf := make([]byte, 1024)
	for {
		// 読み込み
		n, remote, err := conn.ReadFromUDP(buf)
		p.handle(err, "読み込み")

		// 受信データを送信データサイズ[Kbyte]に変換
		kbyte, err := strconv.Atoi(string(buf[:n]))
		p.handle(err, "受信データを送信データサイズ[Kbyte]に変換")

		// 返信用1Kbyteデータ
		s := []byte(random(1024))

		// データ転送時間[秒 float64]を計測
		sec := p.benchmark(func() {
			for i := 0; i < kbyte; i++ {
				_, err := conn.WriteToUDP(s, remote)
				p.handle(err, "")
			}
			_, err := conn.WriteToUDP([]byte("end"), remote)
			p.handle(err, "end送信")
		})

		// 転送速度を計算
		kbps := float64(kbyte*1024*8) / sec / 1000

		// 転送速度計測結果
		p.log(fmt.Sprintf("transrate: %dKbyte, time: %.2fsec, speed: %.2fKbps", kbyte, sec, kbps))
		p.record(kbps)
	}
}
