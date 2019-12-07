package main

import (
    "fmt"
    "github.com/lerox/golang/old_basic_course"

    // "runtime"
)

// func init() {
//     runtime.GOMAXPROCS(runtime.NumCPU())
//     fmt.Println(runtime.NumCPU())
// }

func main() {
    alex := old_basic_course.person{
        firstName: "alex",
        lastName:  "foo bar",
        contactInfo: old_basic_course.contactInfo{
            email:   "hi@ho.com",
            zipCode: 12333,
        },
    }

    alex.print()

    var jey old_basic_course.person
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

    old_basic_course.types()

    old_basic_course.mapExample()

    eb := old_basic_course.englishBot{}
    sb := old_basic_course.spanishBot{}

    old_basic_course.printGreeting(eb)
    old_basic_course.printGreeting(sb)

    old_basic_course.request()

    old_basic_course.statusChecker()
}
