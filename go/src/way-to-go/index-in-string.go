package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "Hi, I'm Marc, Hi."
	fmt.Printf("The position of \"Marc\" is %d\n", strings.Index(str, "Marc"))
	fmt.Printf("First index of \"Hi\" is %d\n", strings.Index(str, "Hi"))
	fmt.Printf("Last index of \"Hi\" is %d\n", strings.LastIndex(str, "Hi"))
	fmt.Printf("The position of \"Bingo\" is %d\n", strings.Index(str, "Bingo"))
	fmt.Printf("\"i\" occurs %d times in str\n", strings.Count(str, "i"))
	newStr := strings.Replace(str, "Hi", "Hello", 2)
	fmt.Println(newStr)
	repeatStr := strings.Repeat(str, 3)
	fmt.Println(repeatStr)
	fmt.Printf("The length of str is %d\n", len(str))
	strASCII := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("The length of strAscii is %d\n", len(strASCII))
	fmt.Printf("The rune count of strASCII is %d\n", utf8.RuneCountInString(strASCII))
	strUTF8 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("The length of strUTF8 is %d\n", len(strUTF8))
	fmt.Printf("The rune count of strUTF8 is %d\n", utf8.RuneCountInString(strUTF8))
	strCN := "贺利华30岁了"
	fmt.Println("strCN length:", len(strCN), "rune count:", utf8.RuneCountInString(strCN))
	fmt.Println(len("牛"))
	fmt.Println(len("🐂"))
}
