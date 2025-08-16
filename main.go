package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%8v", "a\\b")

	for _, number := range numbers {
		fmt.Printf("%8d", number)
	}
	fmt.Print("\n\n")

	for _, a := range numbers {
		fmt.Printf("%8d", a)

		for _, b := range numbers {
			fmt.Printf("%8d", a*b)
		}
		fmt.Print("\n")
	}

	fmt.Print("\n")
}
