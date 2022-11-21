package main

import "fmt"

func main() {
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3

	fmt.Println(arr1)

	//FORM2
	arr2 := []int{1, 2, 4}
	// arr2 := [3]int{1, 2, 4}  // is the same
	fmt.Println(arr2)

	//SLICE  // It consists of a pointer to the array. If i change the values in the second array the first array also will change
	fmt.Println("Slice")
	slice := arr1[:]
	arr1[1] = 23
	slice[2] = 34
	fmt.Println(arr1, slice)

	//Append to add elements to array
	slice = append(slice, 100, 200)
	fmt.Println(slice)

	r1 := slice[1:]  // Take all after the 1 place of the array
	r2 := slice[:2]  // Take the 2 first elements
	r3 := slice[1:3] // Start in index 1 and end in index 3 but not including index 3
	fmt.Println(r1, r2, r3)

	//** MAP reserved word to have a structure key value
	m := map[string]int{"foo": 43}
	fmt.Println(m)
	fmt.Println(m["foo"])
	//changing the value
	m["foo"] = 555
	fmt.Println(m["foo"])
	//Delete element from collection
	delete(m, "foo")
	fmt.Println(m)

}
