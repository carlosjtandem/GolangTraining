package main

import "fmt"

func main() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u.ID = 1
	u.FirstName = "Carlos"
	u.LastName = "Aucancela"
	fmt.Println(u)

	// Form2
	u2 := user{
		ID: 1, FirstName: "Leonel", LastName: "Messi",
	}
	fmt.Println(u2)

}
