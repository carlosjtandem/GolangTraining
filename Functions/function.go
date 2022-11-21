package main

import (
	"errors"
	"fmt"
)

func main() {
	port := 3000
	port, error := startWebServer(port)
	fmt.Println(error)
}

// returning multiple params
func startWebServer(port int) (int, error) {
	fmt.Println("Starting server")
	//todo
	fmt.Println("Server started on port", port)
	// return port,nil
	newError := errors.New("Something went wrong")
	return port, newError
}

//Bool is't the type return
// func startWebServer(port int) bool {
// 	fmt.Println("Starting server")
// 	//todo
// 	fmt.Println("Server started on port", port)
// 	return true
// }
