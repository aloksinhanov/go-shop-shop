package main

import (
	"context"
	"encoding/json"
	"os"

	"net/http"

	"github.com/aloksinhanov/shopshop/internal/server"
)

type Product struct {
	ProductID   int    `json:"id"`
	Description string `json:"description"`
	RatingID    int    `json:"ratingId"`
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	p := Product{
		RatingID:    5,
		Description: "A matt black kitchen sink",
		ProductID:   2,
	}

	pb, err := json.Marshal(p)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(pb)
	w.WriteHeader(http.StatusOK)
}

func main() {
	ctx := context.Background()
	srv := server.New("ProductService")
	srv.Router.HandleFunc("/products/{id}", getProduct).Methods(http.MethodGet)
	port := os.Getenv("PORT")
	srv.Start(ctx, port)
	//srv.Start(ctx, "8001")
}
