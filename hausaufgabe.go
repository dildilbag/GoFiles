package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(".............Variable.............")

	var d = " Techstarter"
	fmt.Println("Hello", d)

	var Correct = true
	fmt.Println("Go rules?", Correct)

	fmt.Println("............Const.............")
	const mynub1 = 20
	const mynub2 = 34.80

	fmt.Printf("The value of const is : %d\n",
		mynub1)

	fmt.Printf("The type of const is : %T\n",
		mynub1)

	fmt.Printf("The value of const is : %f\n",
		mynub2)

	fmt.Printf("The type of const is : %T\n",
		mynub2)

	fmt.Println(".............Aufgabe 2 Bool.............")
	var b1 bool = true // typed declaration with initial value
	var b2 = true      // untyped declaration with initial value
	var b3 bool        // typed declaration without initial value
	b4 := true         // untyped declaration with initial value

	fmt.Println("Typed declaration with initial value : ", b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)

	fmt.Println(".............Aufgabe 2 Int.............")
	var x int = 600
	var y int = -400
	fmt.Printf("Type: %T, value: %v", x, x)
	fmt.Printf("   Type: %T, value: %v", y, y)

	var X uint8 = 225
	fmt.Println(X, X-3)

	// Using 16-bit signed int
	var Y int16 = 32767
	fmt.Println(Y+2, Y-2)

	fmt.Println("   ...............Float..............")
	var m float32 = 123.78
	var n float32 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", m, m)
	fmt.Printf("Type: %T, value: %v", n, n)

	var a float64 = 1.7e+308
	fmt.Printf("   Type: %T, value: %v", a, a)

	fmt.Println("   ...............String..............")
	str := "Dilbag Singh"

	// Display the length of the string
	fmt.Printf("Length of the string is:%d",
		len(str))

	// Display the string
	fmt.Printf("\nString is: %s", str)

	// Display the type of str variable
	fmt.Printf("\nType of str is: %T", str)

	fmt.Println("   ...............Operators..............")
	var p, s = 35, 7

	fmt.Printf("x + y = %d\n", p+s)
	fmt.Printf("x - y = %d\n", p-s)
	fmt.Printf("x * y = %d\n", p*s)
	fmt.Printf("x / y = %d\n", p/s)
	fmt.Printf("x mod y = %d\n", p%s)

	p++
	fmt.Printf("p++ = %d\n", p)

	s--
	fmt.Printf("s-- = %d\n", s)

	var str1 = "Hello GO"

	var str2 = "I didnt hear about you "

	var output = strings.Join([]string{str1, str2}, " ")

	fmt.Println(output)

	var w = []string{"US", "Canada", "Europe", "Australia"}
	var sep = ","
	var out = strings.Join(w, sep)
	fmt.Println(out)

}
