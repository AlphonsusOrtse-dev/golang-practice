***If / Else — Making Decisions
The basic structure in Go:


if condition {
    // runs when condition is true
} else if otherCondition {
    // runs when otherCondition is true
} else {
    // runs when nothing above is true 
}


Comparison & Logical Operators
// Comparison
==   // equal to
!=   // not equal to
>    // greater than
<    // less than
>=   // greater than or equal
<=   // less than or equal

// Logical
&&   // AND — both must be true
||   // OR  — at least one must be true
!    // NOT — flips true to false


********EXERCISE*******
// Write a program that declares a `temperature` variable set to `0.85`.
//  Check if it's valid for an AI model (between `0.0` and `1.0` inclusive). Print either:
```
package main

import "fmt"

func main() {
    temp := 0.85
    if temp >= 0.0 && temp <= 1.0 {
        fmt.Printf("Temperature %.2f is valid\n", temp)
    } else {
        fmt.Printf("Temperature %.2f is out of range\n", temp)
    }
}
```