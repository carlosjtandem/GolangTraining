package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//Open a resource
	f, err := os.Open("myapp.log")
	// Check if exist any error
	if err != nil {
		log.Fatal(err)
	}
	//close file opened
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, "ERROR") {
			fmt.Println(s)
		}
	}
}
