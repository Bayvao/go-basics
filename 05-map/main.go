package main

import (
	"fmt"
)

func main() {

	// 3 ways to create a map

	// first way
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	fmt.Println(colors)

	// second way
	var colorsMap map[string]string
	fmt.Println(colorsMap)

	// third way
	colorMap := make(map[string]string)
	colorMap["white"] = "#fffff"
	fmt.Println(colorMap)

	//delete key from a map
	delete(colors, "red")
	printMap(colors)
}


func printMap(c map[string]string) {
	// map iteration
	for color, hex := range c {
		fmt.Println("Hex code for color: " + color + ", is: " + hex)
	}
}