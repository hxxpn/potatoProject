package main

import (
	"fmt"
	"log"
	"net/http"
)

func potatoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("aaaaaaaa")
	fmt.Fprintf(w, "potato")
}

func otherPotatoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("potato")
	if r.URL.Path != "/potato" {
		http.NotFound(w, r)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "haiiiiiiii")
}

func main() {
	multiplexer := http.NewServeMux()
	multiplexer.HandleFunc("/potatish", otherPotatoHandler)
	http.HandleFunc("/", potatoHandler)
	http.ListenAndServe(":8080", multiplexer)
	fmt.Println("Hello my potato")
}
