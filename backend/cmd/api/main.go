package main

import (
	"log"
	"net/http"

	"DropIt/backend/internal/config"
	apphttp "DropIt/backend/internal/http"
)

func main() {
	cfg := config.Load()
	mux := apphttp.NewRouter()

	log.Printf("backend started on %s", cfg.Addr())
	if err := http.ListenAndServe(cfg.Addr(), mux); err != nil {
		log.Fatal(err)
	}
}
