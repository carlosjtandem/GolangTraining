package main

import "fmt"

func main() {
	//variables Form1
	var i int
	i = 42
	fmt.Println(i)

	//variables Form2
	var f float32 = 3.14
	fmt.Println(f)

	//variables Form3
	personName := "Carlos JAvier"
	fmt.Println(personName)

	//Boolean
	b := true
	fmt.Println(b)

	//Pointers  use * like c++
	firstName := "Carlos"
	fmt.Println(firstName)
	ptr := &firstName
	fmt.Println(ptr, *ptr)
	firstName = "Juana"
	fmt.Println(ptr, *ptr)

	//Constants
	const pi = 3.1415 // Automatically asign float type
	fmt.Println(pi)

	const c int = 3 // Implicit int
	//fmt.Println(c + 1.15)  // Erro because not is posible to add a int and float
	fmt.Println(float32(c) + 1.15) // Erro because not is posible to add a int and float
}
