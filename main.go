package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadFile("public/404.html")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write(bytes)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	http.HandleFunc("/", notFoundHandler)

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
