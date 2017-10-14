package main

import (
    "fmt"
)

func main() {
    alex := person{
        firstName: "alex",
        lastName:  "foo bar",
        contactInfo: contactInfo{
            email:   "hi@ho.com",
            zipCode: 12333,
        },
    }

    alex.print()

    var jey person
    jey.print() // this print "zero values"
    jey.firstName = "hey"
    jey.lastName = "anderson"
    jey.print()

    pointer := &jey

    fmt.Printf("%p - %p", &jey, &pointer)
    fmt.Println("")
    fmt.Printf("%p", pointer)
    fmt.Println("")
    fmt.Println(*pointer)

    pointer.updateFirstName("Jeeey")
    pointer.print()
    jey.print()

    types()

    mapExample()
}
