// main.go
package main

import (
	"log"
	"net/http"
)

//func homeHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Welcome to my skills showcase!")
//}

//func main() {
//	http.HandleFunc("/", homeHandler)
//	log.Println("Starting server on :8080")
//	log.Fatal(http.ListenAndServe(":8080", nil))
//}

func main() {
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
