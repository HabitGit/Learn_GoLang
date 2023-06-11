package main

import "fmt"

func main() {
	const (
		one = 1 + iota*2
		three
		five
		seven
		nine
		eleven
	)

	fmt.Println(one, three, five, seven, nine, eleven)
}
