package main

import (
	"fmt"
)

func main() {
	strUTF8 := "你好啊，golang"
	fmt.Printf("The length of strUTF8 is %d\n", len(strUTF8))
	fmt.Println()

	for pos, char := range strUTF8 {
		fmt.Printf("Character %c starts at byte position %d\n", char, pos)
	}

	fmt.Println()

	for index, rune := range strUTF8 {
		fmt.Printf("%-2d	%d		%U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
	}
}
