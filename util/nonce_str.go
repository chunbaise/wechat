package util

import "github.com/chunbaise/goutil/rand"

func NonceStr() string {
	return string(rand.NewHex())
}
