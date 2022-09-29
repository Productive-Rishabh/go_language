/*
package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello World!")
}
*/
// =====================================================================================================================

// Write a program to print first 10 natural number.
/*
package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 11; i++ {
		fmt.Printf("%v is Natural Number\n", i)
	}
}
*/
// =====================================================================================================================

// Write a program to print first 10 even numbers.
/*
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 20; {
		fmt.Printf("%v is Even Number\n", i)
		i += 2
	}
}
*/
// =====================================================================================================================
// Write a program to print first 10 odd numbers.

/*
package main

import (
	"fmt"
)

	// func main() {
	// 	for i := 1; i < 20; {
	// 		if i < 10 {
	// 			fmt.Printf("0%v\n", i)
	// 		} else {
	// 			fmt.Println(i)
	// 		}
	// 		i += 2
	// 	}
	// }
	func main() {
		for i := 1; i < 20; {
			fmt.Println(i)
			i += 2
		}
	}

*/
// =====================================================================================================================

// Write a program to print first 10 even numbers in reverse order
/*
package main

import (
	"fmt"
)

func main() {
	for i := 18; i >= 0; {
		fmt.Println(i)
		i -= 2
	}
}
*/
// =====================================================================================================================

// Write a program to find the factorial of a number
/*
package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Print("Enter Number to get factorial: ")
	fmt.Scan(&x)
	var z int = 1
	for i := x; i >= 1; i-- {
		z = z * i
	}
	fmt.Print(z)
}
*/
// =====================================================================================================================
// Write a program to find the sum of the digits of a number accepted from user
/*
package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Print("How Many Numbers")
	fmt.Scan(&x)
	for i := x; i >= 1; i-- {
		var catch int
		fmt.Print("Enter Number: ")
		fmt.Scan(&catch)
	}
}
*/
// =====================================================================================================================

// Table of a Given Digit
/*
package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Print("Enter Number: ")
	fmt.Scan(&x)
	var i int
	for i = 1; i < 11; i++ {
		// fmt.Printf("%d * %d: %d\n", x, i, x*i)
		if i < 10 && x*i < 10 {
			fmt.Printf("%d * 0%d : 0%d\n", x, i, x*i)
		} else if i < 10 && x*i >= 10 {
			fmt.Printf("%d * 0%d : %d\n", x, i, x*i)
		} else if i == 10 {
			fmt.Printf("%d * %d : %d\n", x, i, x*i)
		}
	}
}
*/
// =====================================================================================================================

// Write a program to check if the given number is a natural number.
/*
package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello!\n")
	var x int
	fmt.Print("Enter Value: ")
	fmt.Scan(&x)
	if x > 0 {
		fmt.Print("Yes it is a Natural Number")
	} else if x <= 0 {
		fmt.Print("No Not a Natural Number")
	}
}
*/
// ==========================================================================

// star from loops

/*
package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello Triangle\n")
	var star string = "*"
	// var space string = " "
		// for i := 1; i < 6; i++ {
			// fmt.Print("\n")
			// for j := 6; j > i; j-- {
				// fmt.Print(star)
			// }
		// }
		for i := 1; i < 6; i++ {
			fmt.Print("\n")
		for j := 0; j < i; j++ {
			fmt.Print(star)
		}
	}
	for i := 1; i < 6; i++ {
		fmt.Print("\n")
		for j := 5; j > i; j-- {
			fmt.Print(star)
		}
	}

}
*/

// =+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+

// Unsolved
// Get ASCII Value of a Alphabet /Digit
// Armstrong Number
/*
package main

import (
	"fmt"
)

func main() {
	var x string
	var y int
	fmt.Print("Enter Value of X: ")
	fmt.Scan(&x)
	fmt.Print("Re-enter Value of X: ")
	fmt.Scan(&y)
	var l int = len(x)
	// fmt.Printf("%c", x[0])
	fmt.Print(l)
	// var j int = 0
	var i int
	for i = 0; i < l; i++ {
		// fmt.Print("Yes")
		fmt.Printf("\n%c", x[i])

		// j += int("%c",x[i]) * int("%c",x[i]) * int("%c",x[i])
		// fmt.Print("\n", j)
	}
}
*/

// print pattern |_ triangle
/*
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hi Brother\n")
	const s string = "*"
	for i := 1; i < 6; i++ {
		fmt.Print("\n")
		for j := 0; j < i; j++ {
			fmt.Print(s)
		}
	}
}
*/

// print pattern opp |/ triangle
/*
package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 6; i++ {
		fmt.Print("\n")
		for j := 6; j > i; j-- {
			fmt.Print("*")
		}
	}
}
*/

// half Diamond
/*
package main

import (
	"fmt"
)

func main() {
	for n := 1; n < 6; n++ {
		fmt.Print("\n")
		for o := 0; o < n; o++ {
			fmt.Print("*")
		}
	}
	for n := 1; n < 6; n++ {
		fmt.Print("\n")
		for o := 6; o > n; o-- {
			fmt.Print("*")
		}
	}
}
*/

/*
package main

import (

	"fmt"

)

	func main() {
		for i := 1; i < 6; i++ {
			fmt.Print("\n")
			for j := 6; j > i; j-- {
				fmt.Print(" ")
			}
			for k := 0; k < i; k++ {
				fmt.Print("*")
			}
		}
	}
*/

/*
package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 6; i++ {
		fmt.Print("\n")
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for k := 6; k > i; k-- {
			fmt.Print("*")
		}

	}
}
*/
// full diamond

package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 6; i++ {
		fmt.Print("\n")
		for j := 6; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 0; k < i; k++ {
			fmt.Print("*")
		}
		for l := 1; l < i; l++ {
			fmt.Print("*")
		}
	}
	for i := 1; i < 6; i++ {
		fmt.Print("\n")
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for k := 5; k > i; k-- {
			fmt.Print("*")
		}
		for l := 5; l >= i; l-- {
			fmt.Print("*")
		}
	}
}
