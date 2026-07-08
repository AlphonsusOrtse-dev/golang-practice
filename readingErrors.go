package main

import "fmt"

func main() {
	age := 25
	age = "twenty five"
	fmt.Println(age)
}
//this is the message you'll recieve when you run this code:
//cannot use "twenty five" (untyped string constant) as int value in assignment


**Three important pieces of information packed in one line:

"twenty five"  → the value you tried to assign
untyped string constant → what Go thinks that value IS.
int  → what the variable actually accepts
Go is not just saying "wrong" — it is  telling you exactly what it expected and exactly what it got. 


In Go, a variables type is fixed at declaration and can never change.




age := 25
// At this moment Go says:
// "age is now and forever an int"
// No negotiation. No exceptions.

age = "twenty five"
// Go: "age is an int. A string is not an int. Rejected."



### In Go, an uppercase first letter is a signal that means 
**"this is public/exported."** Using it on local variables sends the wrong signal to every developer reading your code.