package main

var Global = 5

func counter() int {
	Global += 1
	return Global
}

func main() {
	defer func() {
		Global = 5
	}()
	counter()
}
