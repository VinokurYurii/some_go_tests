package game

import (
	"bufio"
	"fmt"
	"mult_table/questions"
	"os"
	"strings"
)

func Run(questions []questions.Question) (correctAnswers uint) {
	fmt.Println("Welcome to the Country Quiz Game!")

	for _, question := range questions {
		if askQuestion(question) {
			correctAnswers++
		}
	}

	return correctAnswers
}

func askQuestion(question questions.Question) bool {
	fmt.Printf("\nWhat is the capital of %s? ", question.Country)

	if getUserInput() == strings.ToLower(question.Capital) {
		fmt.Println("Correct!")
		return true
	} else {
		fmt.Printf("Incorrect! The capital of %s is %s.\n", question.Country, question.Capital)
		return false
	}
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Your answer: ")
		result, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}

		return strings.ToLower(strings.TrimRight(result, "\r\n"))
	}

}
