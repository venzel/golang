package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type product struct {
	Name string
	Qte  int
}

func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := chi.NewRouter()

	router.Use(myMiddleware)

	router.Get("/category", func(w http.ResponseWriter, r *http.Request) {
		m := map[string]string{"category": "gamer"}
		render.JSON(w, r, m)
	})

	router.Get("/product", func(w http.ResponseWriter, r *http.Request) {
		queryName := r.URL.Query().Get("name")

		if queryName == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Write([]byte(queryName))
	})

	router.Get("product/{id}", func(w http.ResponseWriter, r *http.Request) {
		productId := chi.URLParam(r, "id")

		w.Write([]byte(productId))
	})

	router.Post("/product", func(w http.ResponseWriter, r *http.Request) {
		var product product

		render.DecodeJSON(r.Body, &product)

		render.JSON(w, r, product)
	})

	router.Put("/product", func(w http.ResponseWriter, r *http.Request) {
		var product product

		render.DecodeJSON(r.Body, &product)

		render.JSON(w, r, product)
	})

	fmt.Println("Listen in port 3000")
	http.ListenAndServe(":3000", router)
}
