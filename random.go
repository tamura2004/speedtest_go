// random
package main

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

// 引数で与えられた文字数(バイト数)のランダムな英数文字列を生成する
func random(length int) string {
	const base = 36
	size := big.NewInt(base)
	n := make([]byte, length)
	for i, _ := range n {
		c, _ := rand.Int(rand.Reader, size)
		n[i] = strconv.FormatInt(c.Int64(), base)[0]
	}
	return string(n)
}
