package cipherer

import (
	"encoding/base64"
	"fmt"
	"os"
)

func Cipher(text, secretKey string) string {
	return base64.StdEncoding.EncodeToString(
		processString(
			[]byte(text),
			[]byte(secretKey),
		),
	)
}

func Decipher(cipheredText, secretKey string) string {
	cipheredBytes, err := base64.StdEncoding.DecodeString(cipheredText)

	if err != nil {
		fmt.Println("Error decoding base64 string:", err)
		os.Exit(1)
	}

	decipheredBytes := processString(cipheredBytes, []byte(secretKey))

	if len(decipheredBytes) == 0 {
		fmt.Println("Deciphered text is empty. Please check your input and secret key.")
		os.Exit(1)
	}

	return string(decipheredBytes)
}

func processString(input, secret []byte) []byte {
	if len(secret) == 0 {
		fmt.Println("Secret key must contain at least 1 character.")
		os.Exit(1)
	}

	for i, b := range input {
		input[i] = b ^ secret[i%len(secret)]
	}
	return input
}
