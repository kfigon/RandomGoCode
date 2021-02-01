package main

import (
	"context"
	"encoding/base64"
	"io"
	"fmt"
	"crypto/hmac"
	"crypto/sha256"
)

func main() {
	hmacDemo()
	base64EncodingDemo()
	contextDemo()
}

// digital signature - verify if data from user is the same we given
// hmac is better hash than md5
func hmacDemo() {
	fmt.Println("=================")
	userInput := "myValueAsd"

	doHash := func(input string) {
		h := hmac.New(sha256.New, []byte("mySecretKey"))
		io.WriteString(h ,input)
		fmt.Printf("%x\n", h.Sum(nil))
	}

	// the same every time
	fmt.Printf("%q\n", userInput)
	doHash(userInput)
	doHash(userInput)

	 // slight change - different hash
	changedValue := " myValueAsd"
	fmt.Printf("%q\n", changedValue)
	doHash(changedValue)
}

// in cookie/url - we can't store everything. Base64 is common
// encoding to have the info fit there
func base64EncodingDemo() {

	fmt.Println("=================")
	myData := `this is data with\ '"123" chars - ^&* :)`

	encoded := base64.StdEncoding.EncodeToString([]byte(myData))

	fmt.Printf("%q\n", myData)
	fmt.Printf("%q\n", encoded)
	decodedBytes, _ := base64.StdEncoding.DecodeString(encoded)
	decoded := string(decodedBytes)
	fmt.Printf("%q\n", decoded)
}

// 
func contextDemo()  {
}