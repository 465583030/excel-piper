package main

import (
	"fmt"
	"strings"
)

func main() {
	printRectangleStars(20, 10)
}

func printRectangleStars(width int, heihgt int) {
	for h := 0; h < heihgt; h++ {
		if h == 0 || h == heihgt-1 {
			fmt.Println(strings.Repeat("*", width))
		} else {
			for w := 0; w < width; w++ {
				if w == 0 || w == width-1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Printf("\n")
		}
	}
}
