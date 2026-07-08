package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello"

	fmt.Printf("%c\n", (str[0]))
	fmt.Println(string(str[1]))
	fmt.Println(str[:1])
	fmt.Println(strings.ToLower(string(str[0])))
	fmt.Println(strings.ToLower((str[:1])))

}
