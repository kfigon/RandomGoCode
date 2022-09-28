package main

import (
	"encoding/hex"
	"bytes"
	"crypto/cipher"
	"log"
	"encoding/base64"
	"crypto/aes"
)

func main() {
	base64Encoded := encodeBase64("hello world!")
	log.Println("base64 encoded:", base64Encoded)
	log.Println("base64 decoded:", decodeBase64(base64Encoded))

	aesEncrypted := encryptAes("hello world!")
	log.Println("aes encrypted:", hex.EncodeToString([]byte(aesEncrypted)))
	log.Println("aes decrypted:", decryptAes(aesEncrypted))

}

func encodeBase64(msg string) string {
	// url encoding != std encoding
	return base64.URLEncoding.EncodeToString([]byte(msg))
}
func decodeBase64(msg string) string {
	res, err := base64.URLEncoding.DecodeString(msg)
	if err != nil {
		log.Println("Error during decoding base64, err:", err)
		return ""
	}
	return string(res)
}

var aesKey = "mySecretKey12345" // 16 bytes
func encryptAes(msg string) string {
	b, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		log.Fatalln("Got error during aes encruption")
		return ""
	}
	iv := make([]byte, aes.BlockSize)
	s := cipher.NewOFB(b, iv)

	buffer := &bytes.Buffer{}
	streamWriter := cipher.StreamWriter {
		S: s,
		W: buffer,
	}
	streamWriter.Write([]byte(msg))
	return buffer.String()
}

func decryptAes(msg string) string {
	return encryptAes(msg)
}