package main

import "fmt"

func main() {
	const (
		modelName   = "chatgpt"
		maxToken    = 20
		temperature = 0.667889
	)
	fmt.Printf("%s, has a max of %d, tokens at %f\n", modelName, maxToken, temperature)
	// %.3f  will print the float value to 3 decimal places
}
