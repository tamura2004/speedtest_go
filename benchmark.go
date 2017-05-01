package main

import (
	"time"
)

// 関数の処理時間[秒 float64]を返す
func (p *parm) benchmark(f func()) (sec float64) {

	// 開始時刻
	startTime := time.Now()

	// 転送処理
	f()

	// 終了時刻
	endTime := time.Now()

	//　処理時間[秒 float64]
	sec = endTime.Sub(startTime).Seconds()
	return
}
