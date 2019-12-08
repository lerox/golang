package old_basic_course

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateFirstName(newOne string) { // getting a pointer
	(*p).firstName = newOne // I have this memory address, just give me the value and then update it
}
