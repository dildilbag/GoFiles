package main

import (
	"fmt"
	"math/rand"
	"time"
)

// main function

func main() {

	fmt.Println("Give Number One Here: ")

	// var type is a int

	var first int

	fmt.Scanln(&first)

	fmt.Println("Give Number Two Here: ")

	var second int

	fmt.Scanln(&second)

	fmt.Print("Total Number is = ")

	// Here i gave result one + second = Answer
	result := first + second
	fmt.Println(result)

	fmt.Println("*******************Aufgabe 3******************")

	fmt.Println("Write your First Name Here: ")
	var one string

	fmt.Scanln(&one)

	fmt.Println("Write your Second Name Here: ")

	var two string

	fmt.Scanln(&two)

	fmt.Println("Where are You from: ")

	var three string

	fmt.Scanln(&three)

	fmt.Println("What is your Profession: ")

	var four string

	fmt.Scanln(&four)

	var array [4]string
	array[0] = one
	array[1] = two
	array[2] = three
	array[3] = four

	fmt.Println("Full Info :", array)

	rand.Seed(time.Now().Unix())
	// random produce a different output on your machine
	rand.Shuffle(len(array), func(a, b int) {
		array[a], array[b] = array[b], array[a]
	})

	fmt.Println("(shuffle Example) :", array)

}
