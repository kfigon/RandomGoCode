package main

import (
	"log"
	"crypto/hmac"
	"crypto/sha256"
)

// signing demo
func main() {
	msg := "hello world"
	hashed := hashStuff(msg)

	log.Println("Equal check on the same:", hmac.Equal(hashed, hashStuff(msg)))
	log.Println("Equal check on different:", hmac.Equal(hashed, hashStuff("helloworld")))
}

func hashStuff(data string) []byte {
	hash := hmac.New(sha256.New, []byte("secretKey"))
	hash.Write([]byte(data)) // write can be used iteratively to sign a file
	return hash.Sum(nil)
}