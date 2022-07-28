package main

import "fmt"

func hallo4() {

	// Today we create MOnth plan with String

	var month = [12]string{"Jan", "Feb", "Mar", "Apr", "Mai", "Jun",
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	y1 := month[0:12]
	fmt.Println(y1)
	fmt.Println("Long: ", len(y1), ":", "Capacity", cap(y1))

	h1 := month[0:6]
	fmt.Println(h1)
	fmt.Println("Long: ", len(h1), ":", "Capacity", cap(h1))

	h2 := month[6:12]
	fmt.Println(h2)
	fmt.Println("Long: ", len(h2), ":", "Capacity", cap(h2))

	q1 := month[0:3]
	fmt.Println(q1)
	fmt.Println("Long: ", len(q1), ":", "Capacity", cap(q1))

	q2 := month[3:6]
	fmt.Println(q2)
	fmt.Println("Long: ", len(q2), ":", "Capacity", cap(q2))

	q3 := month[6:9]
	fmt.Println(q3)
	fmt.Println("Long: ", len(q3), ":", "Capacity", cap(q3))

	q4 := month[9:12]
	fmt.Println(q4)
	fmt.Println("Long: ", len(q4), ":", "Capacity", cap(q4))

	// Two array  each with a length of 10.
	var Days = [10]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(Days)
	fmt.Println("Capacity is :", cap(Days))
	fmt.Println("Lenght is :", len(Days))
	fmt.Println("Array 3 is :", Days[2])

	var countries = [10]string{"India,", "Germany,", "USA,", "Canada,",
		"Greece,", "Uk,", "UAE,",
		"NZ,", "AUS,", "France"}
	fmt.Println(countries)
	fmt.Println("Capacity is :", cap(countries))
	fmt.Println("Lenght is :", len(countries))

	constant := countries[9]
	fmt.Println("Array 10 is :", constant)

}
