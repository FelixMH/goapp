package main

import (
	"fmt"
	"log"
	"net/http"
)

func messageHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola este es un handleFunc</h1>")
}

func messageHFCustom(message string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, message)
		},
	)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", messageHandleFunc)
	mux.Handle("/saludar", messageHFCustom("<h1>Hola Amigos!!!!</h1>"))

	log.Println("Ejecutando server en http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
}
