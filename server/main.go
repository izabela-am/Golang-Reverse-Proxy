package main

import(
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Server)
	fmt.Println("serving in port 8080")
	http.ListenAndServe(":8080", nil)
}

func Server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, %s!\n", r.URL.Path[1:])
	log.Println(r.URL)
}