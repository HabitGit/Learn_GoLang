package main

import "fmt"

func generations(year int) string {
	switch {
	case year >= 1945 && year <= 1964:
		fmt.Println("Привет, бумер!")
	case year >= 1965 && year <= 1980:
		fmt.Println("Привет, представитель Х!")
	case year >= 1981 && year <= 1996:
		fmt.Println("Привет, миллениал!")
	case year >= 1997 && year <= 2012:
		fmt.Println("Привет, зумер!")
	case year >= 2013:
		fmt.Println("Привет, альфа!")
	}
}
