// netgo
package main

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

// 共通エラー処理(手抜き)
func handle(err interface{}) {
	if err != nil {
		panic(err)
	}
}
