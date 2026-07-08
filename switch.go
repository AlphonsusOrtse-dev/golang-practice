What is switch in Go?

👉 A switch is used to choose one action from many options

Instead of writing many if...else statements, you use switch to keep your code clean and readable.


EXAMPLES:

package main

import "fmt"

func main() {
	day := 2

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Unknown day") // output = Tuesday. Because day == 2
	}
}

Switch with DEFAULT value.

package main
import ("fmt")

func main() {
  var day = 4 
  switch day {
  case 1:     
    fmt.Print("Saturday")
  case 2:    
    fmt.Print("Sunday")
  default:    
    fmt.Print("Weekday") // output = weekday
  }
}


How this works:

    * Variable Assignment: The variable day is assigned the value 4.
    * Switch Evaluation: The switch statement checks the value of day against the list of cases.
    * No Match found: Since day is 4, it does not match case 1 or case 2.
    * Fallback: Because no specific case matches, it executes the default block and prints "Weekday".


------------------------------------------------------------
🔹 Multiple Values in One Case

package main
import ("fmt")

func main() {
   day := 5

   switch day {
   case 1,3,5:
    fmt.Println("Odd weekday")
   case 2,4:
     fmt.Println("Even weekday")
   case 6,7:
    fmt.Println("Weekend")
  default:
    fmt.Println("Invalid day of day number") //output = odd weekday
  }
}// this checks if any of the condition is true 
  //if day == 1 || day == 3 || day == 5

//   IMPORTANT RULE:
// Go stops at the first matching

-----------------------------------------------
🔹 Switch Without a Value (Very Powerful)

package main

import "fmt"

func main() {
	score := 75

	switch {
	case score >= 90:
		fmt.Println("Excellent")
	case score >= 60:
		fmt.Println("Good")
	default:
		fmt.Println("Fail")
	}
} // output = Good.

// how Go works: checks each case from top to bottom
//first case: Is 75 >= 90?....no → skip
// second case: Is 75 >= 60?....yes → run this block. 


🔍 So what does this check?

“Is 75 greater than 60 OR equal to 60?”

✅ Break it into two parts

Is 75 > 60? → ✅ YES

Is 75 == 60? → ❌ NO

👉 But here’s the important part:

Only one needs to be true for >= to be true


------------------------------------------
🔥 What is fallthrough in Go?

Normally in Go:

👉 Once a case matches, it stops immediately

BUT…

👉 fallthrough forces Go to continue to the next case, even if it doesn’t match


package main

import "fmt"

func main() {
	x := 1

	switch x {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three") // output = one & two
	}
}

Very Important

fallthrough:

does NOT re-check conditions

just moves to the next block blindly



TAKE A LOOK AT THIS CODE:
x := 3

switch x {
case 1:
	fmt.Println("A")
	fallthrough
case 3:
	fmt.Println("B")
case 4:
	fmt.Println("C") // output B
}

// here, fallthrough fails because,....fallthrough ONLY runs if the case it is inside is executed