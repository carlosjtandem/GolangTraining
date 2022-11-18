package main

import "fmt"

/*
the main function, when part of the main package, identifies the entry point of the application
*/
func main() { // If I move the bracket to the next line, will show a error because It's necesesary here(the reason is because ; is added internally by GO)
	fmt.Println("Hello World. I'm learning.")
}

/* Tips*/
/*
 When I commented the line 9 , GO show an error because we don't use the package fmt. That its important for mantenability
*/
