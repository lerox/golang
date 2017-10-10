package main

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
    jey.print() // this print zero values
    jey.firstName = "hey"
    jey.lastName = "anderson"
    jey.print()

    pointer := &jey // jey holds the data... and the pointer holds the address

    pointer.updateFirstName("Jeeey")
    pointer.print()
    jey.print()

    types()
}
