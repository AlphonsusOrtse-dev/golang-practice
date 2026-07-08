package main

import (
	"fmt"
	"strings"
)

func lastTwo(word []string, n int) []string {
	for i := len(word) - n; i < len(word); i++ {
		word[i] = strings.ToUpper(word[i])
	}
	return word
}

func main() {
	word := []string{"how", "are", "you", "today"}
	fmt.Println(lastTwo(word, 2))
	//this two mean the samething
	// fmt.Println(lastTwo([]string{"how", "are", "you", "today"}, 2))
}
