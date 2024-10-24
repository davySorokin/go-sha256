package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	cryptoHash()
}

func cryptoHash() {
	sum := sha256.Sum256([]byte("Kirill David Sorokin\n"))
	fmt.Printf("%x", sum)
}
