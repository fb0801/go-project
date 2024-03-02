package main

import (
	"fmt"
	"strings"
)

const originalLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func encrypt(key int, plainText string) (result string) {

}

func decrypt(key int, encrypttedText string) (result string) {

}

func main() {
	plainText := "HELLOWORLD"
	fmt.Println("Plain Text", plainText)
	encrypted := encrypt(5, plainText)
	fmt.Println("Encrypted", encrypted)
	decrypted := decrypt(5, encrypted)
	fmt.Println("Decrypted", decrypted)
}