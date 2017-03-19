package main

import "fmt"

func main() {
	var num1 int = 1000
	switch num1 {
	case 88, 98:
		fmt.Println("It's equal to 98")
	case 1000:
		fmt.Println("It's equal to 1000")
		fallthrough
	default:
		fmt.Println("It's not equal to 98 or 1000")
	}
}
