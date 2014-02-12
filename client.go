// client
package main

import (
	"fmt"
	//"io/ioutil"
	"net"
	//"time"
)

// クライアント計測処理
func (p *parm) client() {

	//// 指定回数繰り返し[デフォルト3回]
	//for i := 0; i < p.repeat; i++ {

	//	// サーバ接続処理
	//	p.connect()

	//	// 指定時間待ち[デフォルト10秒]
	//	time.Sleep(time.Duration(p.wait) * time.Second)
	//}
	//}

	//// サーバ接続および転送速度計測
	//func (p *parm) connect() {

	// 接続urlを編集
	//url := "http://" + p.host + ":" + p.port + "/speed"
	addr, err := net.ResolveUDPAddr("udp", p.host+":"+p.port)
	handle(err)

	conn, err := net.DialUDP("udp", nil, addr)
	handle(err)
	defer conn.Close()

	// 開始ログ
	p.log("start connecting")

	// 接続
	//res, err := http.Get(url)
	//handle(err)
	//defer res.Body.Close()

	// 転送サイズ
	kbyte := float64(p.length)

	// 転送レートを計測[秒]
	buf := make([]byte, 1024)
	sec := p.benchmark(func() {
		conn.Write([]byte(fmt.Sprintf("%d", p.length)))
		for {
			n, err := conn.Read(buf)
			handle(err)
			if string(buf[:n]) == "end" {
				break
			}
		}
		//body, err := ioutil.ReadAll(res.Body)
		//handle(err)
		//kbyte = float64(len(body)) / 1024
	})

	// 転送レート計算
	kbps := kbyte * 8.0 / sec

	// 転送サイズ、レート記録
	p.log(fmt.Sprintf("transrate: %.2fKbyte, time: %.2fsec, speed: %.2fKbps", kbyte, sec, kbps))

}
