package main

import (
	"fmt"
	"strings"
)

func toUpper(s string) string {
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])

}
func main() {
	fmt.Println(toUpper("how are you today"))
}
