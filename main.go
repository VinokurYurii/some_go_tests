package main

import (
	"fmt"
	"mult_table/game"
	"mult_table/questions"
	"mult_table/shuffler"
	"os"
)

func main() {
	questions, err := questions.LoadQuestions()
	if err != nil {
		fmt.Printf("Error loading questions: %v\n", err)
		os.Exit(1)
	}

	shuffler.Shuffle(questions)

	correctAnswers := game.Run(questions)
	fmt.Printf("\nYou answered %d questions correctly!\n", correctAnswers)
}
