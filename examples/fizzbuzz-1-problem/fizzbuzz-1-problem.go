//We’re going to implement a go program that does this.  In order to
//do this, you’ll need to be introduced to two more operators.
package main

import "fmt"

func main() {

	//The first is the boolean and operation, which is &&.  This
	//takes two boolean values, and will return true if and only
	//if they are both true.  Here a few quick examples.
	fmt.Println(true && true)
	fmt.Println(false && true)
	fmt.Println(false && false)

	//You can see that only the line with both inputs being true is true.
	//The second tool you will need to do your job is the modulus
	//operator, which is the “%” symbol.  This operator will return
	//the remainder of two numbers after they are divided.  For example,
	//6 % 3 will return zero because 6 is divisible by 3, and 5 % 2 will
	//return 1 because 5 has a remainder of 1 when divided by 2.
	fmt.Println(6 % 3)

	fmt.Println(5 % 2)
}
