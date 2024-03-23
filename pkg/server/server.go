package server

import (
	"fmt"
	"log"
	"net/http"
)

func RunServer() {
	http.HandleFunc("/", dummyHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func dummyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hey from hypervector")
}
