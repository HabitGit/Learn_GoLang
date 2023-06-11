package main

import "fmt"

func main() {
	type String string
	var str string
	str = "Hello, world!"
	fmt.Println(String(str[0]))
}
