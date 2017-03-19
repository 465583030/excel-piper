package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}
	strASCII := "hello golang world"
	for idx := 0; idx < len(strASCII); idx++ {
		fmt.Printf("Character on position %d is %c\n", idx, strASCII[idx])
	}

	strUTF8 := "你好啊，Golang"
	for idx := 0; idx < len(strUTF8); idx++ {
		fmt.Printf("Character on position %d is %c\n", idx, strUTF8[idx])
	}

	for repeatCount := 1; repeatCount <= 25; repeatCount++ {
		fmt.Println(strings.Repeat("G", repeatCount))
	}

	for idx := 1; idx <= 100; idx++ {
		switch {
		case idx%15 == 0:
			fmt.Println("FizzBuzz")
		case idx%3 == 0:
			fmt.Println("Fizz")
		case idx%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(idx)
		}
	}

	var printTimes = 10
	for printTimes >= 0 {
		printTimes--
		fmt.Printf("PrintTimes is %d\n", printTimes)
	}
}
