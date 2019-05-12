package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../bruker"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request")
	ocj := &bruker.Bruker{
		Navn:  "Ole Christian",
		Epost: "olecj12@uia.no"}

	switch r.Method {
	case "GET":
		j, _ := json.Marshal(ocj)
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	case "POST":
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "POST not suppoerted")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Status not allowed")
	}
}

func main() {
	//Serveren kan nåes på "127.0.0.1:8080/bruker"
	http.HandleFunc("/bruker", handler)

	fmt.Println("HTTP server started")
	http.ListenAndServe(":8080", nil)
}
