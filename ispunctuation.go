package main

import "fmt"

func isPunctuation(s string) bool {
	for _, r := range s {
		if r == ',' || r == '!' {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(isPunctuation(","))
	fmt.Println(isPunctuation("!"))
	fmt.Println(isPunctuation("x"))

}
