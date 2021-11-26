package main

import (
	"GO/trevor/bookings-31/pkg/handlers"
	"fmt"
	"net/http"
)

const portNumber = ":3000"

// main is the main function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	tmp := fmt.Sprintf("Staring application on port %s", portNumber)
	fmt.Println(tmp)
	_ = http.ListenAndServe(portNumber, nil)
}
