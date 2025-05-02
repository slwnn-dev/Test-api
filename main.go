package main
//test1231122

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go123!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Starting server on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}

# test comment
