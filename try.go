package main

import (
	"fmt"
	"strings"
	// "strings"
)

func checkdec(s string) string {
	if strings.ContainsAny(s, "!") {
		return "punct"
	}
	r := rune(s[0])

	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
		return "alpha"
	}
	if r >= '0' && r <= '9' {
		return "numb"
	}
	return "unknown"
}
func main() {
	fmt.Println(checkdec("!"))
	fmt.Println(checkdec("a"))
	fmt.Println(checkdec("2"))
	fmt.Println(checkdec("%"))

}
