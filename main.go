package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("rsaEncrypt", js.NewCallback(encrypt))
	<-c
}

//in Index.html javascript is passing (input, output) which are the js.Values
func encrypt(input []js.Value) {

	PrivateKey, err := rsa.GenerateKey(rand.Reader, 1000)
	if err != nil {
		fmt.Println(err)
	}
	PublicKey := &PrivateKey.PublicKey

	message := input[0].String()

	byteMessage := []byte(message)
	label := []byte("")
	hash := sha256.New()

	//encrypt with Optimal Asymmetric Encryption Padding scheme
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, PublicKey, byteMessage, label)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("OAEP encrypted [%s] to \n[%x]\n", string(message), ciphertext)
	output := ciphertext
	input[1].Invoke(fmt.Sprintf("[%x]",output))
}