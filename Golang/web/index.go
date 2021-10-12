package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", nil)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Connect")
	})

	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"product request")
	})

	
	
}