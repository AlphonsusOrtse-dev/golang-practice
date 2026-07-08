package main

import "fmt"

func main() {

	for i := 0; i <= 9; i++ {
		fmt.Print(i)
	}

}
package main

import "fmt"

func main() {
	s := "Hello"

	// 1. ONLY THE INDEX
	// Use this if you only care about the position (0, 1, 2...)
	fmt.Println("--- Only Index ---")
	for i := range s {
		fmt.Printf("Position: %d\n", i)
	}

	// 2. BOTH INDEX AND VALUE
	// Use this if you need to know WHERE you are and WHAT is there
	fmt.Println("\n--- Index and Value ---")
	for i, r := range s {
		fmt.Printf("At position %d, the character is %c\n", i, r)
	}

	// 3. ONLY THE VALUE (Using the "Trash Can" _)
	// Use this if you only care about the characters themselves
	fmt.Println("\n--- Only Value ---")
	for _, r := range s {
		fmt.Printf("Character: %c\n", r)
	}
}

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// We use _ because we don't care about the index (0, 1, 2...)
	// We only want the actual number (value)
	for _, value := range numbers {
		fmt.Println(value)
	}
}


package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		// Use the % (modulo) operator to check if the remainder is 0
		if i % 2 == 0 {
			fmt.Printf("%d is even\n", i)
		} else {
			fmt.Printf("%d is odd\n", i)
		}
	}
}


package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range numbers {
		// Check if the number is greater than 5
		if value > 5 {
			fmt.Println("Large number:", value)
		}
	}
}


package main

import "fmt"

// 1. Define the function (it takes an 'int' named 'num')
func CheckEvenOdd(num int) {
	if num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}
}

func main() {
	// 2. Loop from 1 to 10
	for i := 1; i <= 10; i++ {
		// 3. Call your function and pass 'i' into it
		CheckEvenOdd(i)
	}
}


func GetLabel(num int) string {
    if num % 2 == 0 {
        return "even"
    }
    return "odd"
}

package main

import "fmt"

// This function takes a slice of integers []int
func ProcessNumbers(nums []int) {
	for _, n := range nums { // The loop is HIDDEN inside here
		if n%2 == 0 {
			fmt.Printf("%d is even\n", n)
		} else {
			fmt.Printf("%d is odd\n", n)
		}
	}
}

func main() {
	myList := []int{1, 2, 3, 4, 5}
	
	// Just one simple call!
	ProcessNumbers(myList) 
}
