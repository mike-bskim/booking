package main

import "fmt"

func main() {
	fmt.Println("Test")

	var whatToSay string
	whatToSay = "hi kim"
	fmt.Println(whatToSay)

	whatSaid := saySomething()
	fmt.Println(whatSaid)

}

func saySomething() string {
	return "something"
}
