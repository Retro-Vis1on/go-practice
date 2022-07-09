package main

import (
	"fmt"
	"log"
	"net/http"
)

//w=response r=request
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse form err: %v", err)
		return
	}
	fmt.Fprint(w, "Post Request Successfull\n")
	name := r.PostForm.Get("name")
	address := r.PostForm.Get("address")
	fmt.Fprintf(w, "Name: %s\nAddress: %s\n", name, address)
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
