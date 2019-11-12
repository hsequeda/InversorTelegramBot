package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", helloworld)
	if err := http.ListenAndServe(":"+port, http.DefaultServeMux); err != nil {
		fmt.Println(err)
	}
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Hello, World!")); err != nil {
		fmt.Println(err)
	}
}
