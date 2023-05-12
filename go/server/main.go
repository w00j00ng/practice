package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := 5000

	fmt.Printf("start server on port %d\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello world"))
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Println("ERROR: "+ err.Error())
		os.Exit(1)
	}
}
