package main

import "fmt"

func main() {
	colors := make(map[string]string)
	colors["red"] = "#ff0000"
	colors["green"] = "#00ff00"
	colors["blue"] = "#0000ff"
	printMap(colors)
}

func printMap(colors map[string]string) {
	for color, hex := range colors {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
