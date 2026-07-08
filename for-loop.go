***Issue  — Variable Shadowing 


gocount := 0                        // declares 'count' in main's scope
for count := 512; ...; count++ {  // ← this := creates a BRAND NEW 'count'
                                  // it shadows (hides) the outer count
                                  // the outer count stays 0 forever
}
fmt.Printf("... %d\n", count)     // prints the OUTER count — still 0




*****Key Rule to Remember

a. Inside loop → prints many times

b. Outside loop → prints once


package main

import "fmt"

func main() {
	m := 20
	for n := 1; n <= 5; n++ {
		m += n
		fmt.Println(m)
	}
}
🔍 What Your Code Does Step-by-Step

Start: m = 20

Loop runs from n = 1 to 5

Iteration 1:
m = 20 + 1 = 21
print → 21

Iteration 2:
m = 21 + 2 = 23
print → 23

Iteration 3:
m = 23 + 3 = 26
print → 26

Iteration 4:
m = 26 + 4 = 30
print → 30


Iteration 5:
m = 30 + 5 = 35
print → 35

🖥️ Output You’ll See
21
23
26
30
35

👉 That’s multiple outputs, not a single one.


✅ If You Want a SINGLE Output

Move Println outside the loop:

package main

import "fmt"

func main() {
	m := 20
	for n := 1; n <= 5; n++ {
		m += n
	}

	fmt.Println(m) // only prints once
}




/ Write a `for` loop that starts at `0` tokens and adds `512` each iteration.
/ Stop when tokens reach or exceed `4096`. Print the count at each step.

// Expected output:
// ```
// Tokens: 512
// Tokens: 1024
// Tokens: 1536
// Tokens: 2048
// Tokens: 2560
// Tokens: 3072
// Tokens: 3584
// Tokens: 4096
// package main


package main

import "fmt"

func main() {
	tokens := 0

	for tokens < 4096 {
		tokens += 512
		fmt.Printf("Tokens: %d\n", tokens)
	}
}



Step-by-Step Execution

Step	tokens < 4096?	Action	New tokens	Output
1	✅ true	tokens += 512	512	    512
2	✅ true	tokens += 512	1024	1024
3	✅ true	tokens += 512	1536	1536
4	✅ true	tokens += 512	2048	2048
5	✅ true	tokens += 512	2560	2560
6	✅ true	tokens += 512	3072	3072
7	✅ true	tokens += 512	3584	3584
8	✅ true	tokens += 512	4096	4096
9	❌ false	loop stops	—	—

OTHER WAYS TO SOLVE IT:


package main

import "fmt"

func main() {
	for tokens := 512; tokens <= 4096; tokens += 512 {
		fmt.Println(tokens)
	}
}