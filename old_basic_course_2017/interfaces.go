package old_basic_course_2017

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hi"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
