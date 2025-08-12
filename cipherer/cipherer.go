package cipherer

import (
	"encoding/base64"
	"errors"
)

func Cipher(text, secretKey string) (string, error) {
	a, err := processString([]byte(text), []byte(secretKey))

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(a), nil
}

func Decipher(cipheredText, secretKey string) (string, error) {
	cipheredBytes, err := base64.StdEncoding.DecodeString(cipheredText)

	if err != nil {
		return "", err
	}

	decipheredBytes, processErr := processString(cipheredBytes, []byte(secretKey))

	if processErr != nil {
		return "", processErr
	}

	if len(decipheredBytes) == 0 {
		return "", errors.New("deciphered text is empty")
	}

	return string(decipheredBytes), nil
}

func processString(input, secret []byte) ([]byte, error) {
	if len(secret) == 0 {
		return nil, errors.New("secret key must contain at least 1 character")
	}

	for i, b := range input {
		input[i] = b ^ secret[i%len(secret)] // 0..len(secret)
	}
	return input, nil
}
