package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed cncgkottayam.png
var static embed.FS

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		http.ServeFileFS(w, r, static, "cncgkottayam.png")
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

