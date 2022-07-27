package main

import "fmt"

func main1() {

	// Das ist Aufgabe I

	var array1 = [5]float32{2.0, 4.5, 78.0, 68.0, 12.0}

	var summe float32 = 0

	for i := 0; i < len(array1); i++ {
		summe = summe + float32(array1[i])
	}
	fmt.Println("Die Summe des Arrays ist: ", summe)
	fmt.Println("Der Durchschnittswert ist: ", summe/5)

	//Das ist Aufgabe II
	var array2 = [3]string{"Dilbag Singh", "Thind", "Punjab"}
	fmt.Println("Ich bin ", array2[0], array2[1], "und wurde in", array2[2], "geboren.")

	//Das ist die Aufgabe III
	var length int = 0

	var array = [10]int{10, 20, 3, 4, 5, 6, 7, 8, 30, 40}
	fmt.Println(array)

	for i := 0; i < len(array); i++ {

		length++

	}
	var summe1 int = 0

	for i := 0; i < len(array); i++ {

		summe1 = summe1 + (array[i])

	}
	var teilwert float32 = float32(summe1 / 2)

	fmt.Println("Länge des Arrays: ", length)
	fmt.Println("Wert an der Stelle 5: ", array[5])
	fmt.Println("Gesamtwert des Arrays: ", summe1)

	fmt.Println("Gesamtwert durche 2 geteilt: ", teilwert)

}
