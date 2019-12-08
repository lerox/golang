// value types & reference types
package old_basic_course

import "fmt"

func types() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	updateSlice(mySlice)

	fmt.Println(mySlice)
	// It prints: [Bye There How Are You]
	//
	// Because a slice is a reference type and not a value type
	// In short words a slice will contain the length, max capacity & a pointer to an array
	// This will not happen with a struct for example.
}

func updateSlice(s []string) {
	s[0] = "Bye"
	// mySlice will be passed by value
	// s will be a copy of mySlice
	// both of them will point to the same array
}
