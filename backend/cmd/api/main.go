package main

import (
	"log"
	"net/http"

	apphttp "your-project/backend/internal/http"
)

func main() {
	mux := apphttp.NewRouter()

	log.Println("backend started on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}