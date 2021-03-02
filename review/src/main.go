package main

import (
	"context"
	"encoding/json"
	"os"

	"net/http"

	"github.com/aloksinhanov/shopshop/internal/server"
)

type Review struct {
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
	ProductID int    `json:"productId"`
}

func getReview(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	rv := Review{
		Rating:    5,
		Comment:   "A very durable product",
		ProductID: 100,
	}

	rb, err := json.Marshal(rv)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(rb)
	w.WriteHeader(http.StatusOK)
}

func main() {
	ctx := context.Background()
	srv := server.New("ReviewService")
	srv.Router.HandleFunc("/reviews/{id}", getReview)

	port := os.Getenv("PORT")
	srv.Start(ctx, port)
}
