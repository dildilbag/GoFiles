package main

import "fmt"

func main1() {
	fmt.Println("")

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

}
