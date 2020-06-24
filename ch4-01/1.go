package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 :=
		sha256.Sum256([]byte("x"))
	c2 :=
		sha256.Sum256([]byte("X"))

	count := compareSha256(c1, c2)
	fmt.Println(count)
}

func compareSha256(a [32]byte, b [32]byte) int {
	var count int
	for i := 0; i < 32; i++ {
		for j := 0; j < 8; j++ {
			if (a[i] >> j & 1) != (b[i] >> j & 1) {
				count++
			}
		}
	}
	return count
}
