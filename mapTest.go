package main

import "fmt"

func main() {
	A := [7]int{1, 5, 6, 7, 4, 3, 10}
	k := 15
	var result []int
	for indexOne := range A {
		for indexTwo := range A {
			if indexOne == indexTwo {
				continue
			}
			if A[indexOne]+A[indexTwo] == k {
				result = append(result, indexOne, indexTwo)
			}
		}
	}
	fmt.Println(result)
}
