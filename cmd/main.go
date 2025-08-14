package main

import (
	"fmt"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server err", err)
	}
	fmt.Println("Server is running")
}
