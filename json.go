package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name           string    `json:"Имя"`
	Email          string    `json:"Почта"`
	DateOfBirthday time.Time `json:"-"`
}

func main() {
	personData := Person{
		Name:           "Алекс",
		Email:          "alex@yandex.ru",
		DateOfBirthday: time.Now(),
	}
	person := NewPerson(personData)
	jsPerson, _ := json.Marshal(person)
	fmt.Println(string(jsPerson))
}

func NewPerson(data Person) Person {
	return Person{
		Name:           data.Name,
		Email:          data.Email,
		DateOfBirthday: data.DateOfBirthday,
	}
}
