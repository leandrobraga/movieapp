package main

import (
	"log"
	"movieapp/metadata/internal/controller/metadata"
	httphanlder "movieapp/metadata/internal/handler/http"
	"movieapp/metadata/internal/repository/memory"
	"net/http"
)

func main() {
	log.Println("Starting the movie metadata service")
	repo := memory.New()
	ctr := metadata.New(repo)
	h := httphanlder.New(ctr)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
