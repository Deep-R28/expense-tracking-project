package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {


	
	//Setting up server
	port := ":8080"
	fmt.Printf("Server starting at port: %v", port)
	//Starting Server - http://localhost:8080/
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
