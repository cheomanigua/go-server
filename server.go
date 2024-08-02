package main

import (
	"net/http"
)

func main() {
	// METODO 1 (más elaborado)
	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fs)
	// http.ListenAndServe(":8080", nil)

	// METODO 2 (intermedio)
	// http.Handle("/", http.FileServer(http.Dir("./static")))
	// http.ListenAndServe(":8080", nil)

	// METODO 3 (más sencillo)
	http.ListenAndServe(":8080", http.FileServer(http.Dir("./static")))

	// log.Print("Listening on :8080...")
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
