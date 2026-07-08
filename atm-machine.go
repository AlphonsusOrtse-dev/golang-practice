package main

import "fmt"

func main() {
	mak()
	choice := 5
	
	switch choice {
	case 1:
		fmt.Println("Check Balance")
	case 2:
		fmt.Println("Withdraw")
	case 3:
		fmt.Println("Deposit") 
	case 4:
		fmt.Println("Exit")  
	default:
		fmt.Println("Invalid Input") // any number apart from (1-4) will print (Invalid Input)
		                             // that's what the default option does
	}
}
func mak() {
    model := "gpt-3"

    switch model {
    case "gemini-2.0":
        fmt.Println("Google")
    case "gpt-4":
		fmt.Println("OpenAI")
    case "claude-3":
        fmt.Println("Anthropic")
    default:
        fmt.Println("Unknown provider")
    }
}