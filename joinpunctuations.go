package main

import (
	"fmt"
	"strings"
)

func joinPunctuations(s []string) string {
	result := strings.Join(s, "")

	result = strings.ReplaceAll(result, " ,", ",")
	result = strings.ReplaceAll(result, " !", "!")
	return result
	// you can also return like this without putting result,
	// if you have only one punctuation to replace
	// 	return strings.ReplaceAll(result, " ,", ",")

}

func main() {
	fmt.Println(joinPunctuations([]string{"hello ,,, !"}))
}
