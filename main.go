package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")
	
	// Some code with potential security issues for gosec to catch
	handleUserInput()
	startServer()
}

func handleUserInput() {
	// G204: Subprocess launched with variable
	userCmd := os.Getenv("USER_COMMAND")
	if userCmd != "" {
		fmt.Printf("Executing command: %s\n", userCmd)
	}
	
	// G401: Use of weak cryptographic primitive
	password := "hardcoded_password"
	fmt.Printf("Using password: %s\n", password)
}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// G107: Potential HTTP request made with variable url
		userID := r.URL.Query().Get("id")
		if userID != "" {
			id, err := strconv.Atoi(userID)
			if err != nil {
				http.Error(w, "Invalid ID", http.StatusBadRequest)
				return
			}
			fmt.Fprintf(w, "User ID: %d", id)
		} else {
			fmt.Fprintf(w, "Hello from Go server!")
		}
	})
	
	// G114: Use of net/http serve function that has no support for setting timeouts
	log.Fatal(http.ListenAndServe(":8080", nil))
}