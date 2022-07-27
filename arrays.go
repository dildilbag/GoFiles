package main

import "fmt"

func main() {
	fmt.Print("")

	// möglichkeit wie man eine Array unter Go anlegen Kann!
	var array1 [3]int // diese array(int) halt nur 3 folders
	array1[0] = 293
	array1[1] = 456
	array1[2] = 123

	// inhalt des Arrays ausgeben auf der konsole
	fmt.Println(array1)

	// 2, möglichkeit wie man eine Array unter Go anlegen Kann!
	var array2 = [5]string{"text", "hallo", "hey", "du", "ich"}
	fmt.Println(array2)

	//Float Möglichkeit
	var array3 = [4]float32{23.2, 24.1, 58.8, 68.1}
	//Länge eines Arrays ausgeben!
	fmt.Println("Die Laenge von array3 ist: ", +len(array3))
	fmt.Println(array3)

	// array int
	var array4 = [5]int{30, 99, 12, 145} // when int 5 und zahl 4 ist 5 come 0 aus
	fmt.Println(array4)

	//for Loops(scliefe)
	for i := 0; i <= 5; i++ {
		fmt.Println("For Loops", +i)
	}

	// for Loops mit array
	summe := 0
	for i := 0; i < len(array4); i++ {
		summe = summe + array4[i]
		fmt.Println("Die Summe ist:", summe)
	}

	var array5 [10]int // array mit int
	array5[0] = 30
	array5[1] = 44
	array5[2] = 20
	array5[3] = 49
	array5[4] = 38
	array5[5] = 94

	fmt.Println(array5)

	var array6 [10]int // array mit float32
	array6[4] = 44
	array6[9] = 22
	fmt.Println(array6)

	copy(array5[:], array6[:])
	fmt.Println(array5[:], array6[:])

}
