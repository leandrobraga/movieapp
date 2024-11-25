package main

import (
	"log"
	"movieapp/rating/internal/controller/rating"
	httphanlder "movieapp/rating/internal/handler/http"
	"movieapp/rating/internal/repository/memory"
	"net/http"
)

func main() {
	log.Println("Starting the rating service")
	repo := memory.New()
	ctrl := rating.New(repo)
	h := httphanlder.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
}
