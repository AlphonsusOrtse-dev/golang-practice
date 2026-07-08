package main

import (
	"fmt"
	"strings"
)

func checPunts(s string) string {
	if len(s) == 0 {
		return "empty"
	}

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
	fmt.Println(checPunts("!"))
	fmt.Println(checPunts("a"))
	fmt.Println(checPunts("2"))
	fmt.Println(checPunts("%"))
}
