// server
package main

import (
	"fmt"
	"net/http"
)

// ｈｔｔｐサーバ処理
func (p *parm) server() {

	// 開始メッセージ
	p.log(fmt.Sprintf("start server. listening port %s.", p.port))

	// topページのリクエスト処理
	http.HandleFunc("/speed", func(w http.ResponseWriter, r *http.Request) {

		// 接続情報を出力
		p.log(fmt.Sprintf("%s %s from %s", r.Method, r.RequestURI, r.RemoteAddr))

		// length[Kbyte]の文字列を生成
		msg := random(p.length * 1024)

		// データ転送時間[秒 float64]を計測
		sec := p.benchmark(func() {
			fmt.Fprint(w, msg)
		})

		// 転送速度を計算
		kbps := float64(p.length*1024*8) / sec / 1000

		// 転送速度計測結果
		p.log(fmt.Sprintf("transrate: %dKbyte, time: %.2fsec, speed: %.2fKbps", p.length, sec, kbps))

	})

	// 指定ポートでListen
	err := http.ListenAndServe(":"+p.port, nil)

	// エラー処理
	handle(err)
}
