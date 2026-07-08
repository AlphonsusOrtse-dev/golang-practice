package main

import (
	"fmt"
	"strconv"
)

func main() {
	myNum := "41"
	myInt, err := strconv.Atoi(myNum)

	fmt.Println(myInt, err) // 41 <nil>

}

//you can do sth like this if you dont want the <nil>

// myNum := "41"
// myInt, err := strconv.Atoi(myNum)
// if err == nil {
// 	fmt.Println(myInt)
// }

// Why the err is there (The "Why")
// Functions like strconv.Atoi return two values because conversion can fail.

//     If myNum is "41", err will be nil (nothing went wrong).
//     If myNum is "ABC", the computer can not turn that into a number.
// 	In this case, myInt will be 0 and err will contain a message explaining the failure.

// If you do not care about the error and just want the number,
// you can use the blank identifier (_):

// myInt, _ := strconv.Atoi(myNum) // Use "_" to tell Go:
// // "I know there's an error, but ignore it."
