package main

import "fmt"

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}

	for indexOne := range input {
		for indexTwo := range input {
			if indexOne == indexTwo {
				continue
			}
			if input[indexOne] == input[indexTwo] {
				input = append(input[:indexTwo], input[indexTwo+1:]...)
				fmt.Println("asd", input)
			}
		}
	}
	fmt.Println("Result: ", input)
}
