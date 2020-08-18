package main

import (
	"log"
	"net/http"
	"os"
)

func staticFileHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat("public/" + r.URL.Path); os.IsNotExist(err) {
		http.ServeFile(w, r, "public/404.html")
	} else {
		http.ServeFile(w, r, "public/"+r.URL.Path)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	http.HandleFunc("/", staticFileHandler)

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
