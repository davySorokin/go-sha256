package main

import (
	"crypto/sha256"
	"fmt"
)

// go run basic.go
// a020ec7a796167d3a0d9a10418639eb952b1aec026a06c9a8f27ccef9148797cKJHL76JYKL
func main() {
	code()
}

// 'New' returns a new hash.Hash computing the SHA256 checksum.
func code() {
	h := sha256.New()
	h.Write([]byte("Kirill David Sorokin\n"))
	fmt.Printf("%x", h.Sum(nil))
}
