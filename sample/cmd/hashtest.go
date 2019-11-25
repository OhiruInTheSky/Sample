package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func main() {

	// 公開鍵想定
	key := "public-key"
	// 32bit OS用 バイト化して256
	sha256 := sha256.Sum256([]byte(key))
	// 64bit OS用 バイト化して512
	sha512 := sha512.Sum512([]byte(key))

	fmt.Println("sha256: " + hex.EncodeToString(sha256[:]))
	fmt.Println("sha512: " + hex.EncodeToString(sha512[:]))
}
