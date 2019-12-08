package old_basic_course

import "fmt"

func mapExample() {
	// var colors map[string]string
	// colors := make(map[int]string)
	// colors[10] = "bla"
	// delete(colors, 10)

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
	}

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for", color, "is", hex)
	}
}
