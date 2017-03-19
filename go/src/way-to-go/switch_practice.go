package main

import (
	"fmt"
)

func main() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

	fmt.Printf("season(1) is %s\n", season(1))
	fmt.Printf("season(7) is %s\n", season(7))
	fmt.Printf("season(-1) is %s\n", season(-1))
}

func season(month int) string {
	switch month {
	case 1, 2, 3:
		return "春"
	case 4, 5, 6:
		return "夏"
	case 7, 8, 9:
		return "秋"
	case 10, 11, 12:
		return "冬"
	default:
		return "无"
	}
}
