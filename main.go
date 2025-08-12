package main

import (
	"bufio"
	"cipherer/cipherer"
	"flag"
	"fmt"
	"os"
	"strings"
)

var mode = flag.String("mode", "cipher", "Set 'cipher' or 'decipher'. Default is 'cipher'")
var secretKey = flag.String("secret", "", "Your secret key. Must ccontain at least 1 character.")

func main() {
	flag.Parse()

	switch *mode {
	case "cipher":
		userInput := getUserInput("Enter text to cipher: ")
		cipheredText, err := cipherer.Cipher(userInput, *secretKey)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error during encrypting text: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("You encrypted text:\n", cipheredText)
	case "decipher":
		ciperedInput := getUserInput("Enter your cipher data to decipher: ")
		decipheredText, err := cipherer.Decipher(ciperedInput, *secretKey)

		if err != nil {
			fmt.Println("Error during deciphering:", err)
			os.Exit(1)
		}

		fmt.Println("Your decrypted text:\n", decipheredText)
	default:
		fmt.Println("Invalid mode. Use 'cipher' or 'decipher'.")
		os.Exit(1)
	}

}

func getUserInput(msg string) string {
	fmt.Print(msg)
	reader := bufio.NewReader(os.Stdin)

	for {
		userInput, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		return strings.TrimRight(userInput, "\n")
	}
}
