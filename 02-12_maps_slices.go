package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// maps
	myMap1 := make(map[string]interface{})
	myMap := make(map[string]User)

	me := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	myMap["me"] = me
	myMap1["me"] = me

	log.Println(myMap["me"].FirstName)
	log.Println("myMap:", myMap["me"])
	log.Println("myMap1:", myMap1["me"])

	third := myMap1["me"].(User)
	log.Println(third.FirstName, third.LastName)

	you := User{
		FirstName: "BS",
		LastName:  "KIM",
	}
	myMap["you"] = you
	log.Println(you)

	you.FirstName = "bskim"
	log.Println(you)

	log.Println("---------------------------")

	// slices
	var mySlice []int

	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	// sort.Ints(mySlice)

	log.Println(mySlice)

	// shorthand for slices

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)

	// print first two elements of slice, starting at first position
	log.Println(numbers[0:2])

	// create a slice of strings
	names := []string{"one", "seven", "fish", "cat"}
	log.Println(names)
	sort.Strings(names)
	log.Println(names)
}
