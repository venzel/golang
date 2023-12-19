// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/go-chi/chi"
// 	"github.com/go-chi/chi/middleware"
// 	"github.com/go-chi/render"
// 	"venzel.com.br/chi/domain"
// )

// func MyMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("Before: Executing middleware")
// 		next.ServeHTTP(w, r)
// 		fmt.Println("After: Executing middleware")
// 	})
// }

// func main() {
// 	route := chi.NewRouter()

// 	route.Use(middleware.Logger)
// 	route.Use(MyMiddleware)

// 	route.Get("/", func(res http.ResponseWriter, req *http.Request) {
// 		res.Write([]byte("hi"))
// 	})

// 	route.Get("/{id}", func(res http.ResponseWriter, req *http.Request) {
// 		userId := chi.URLParam(req, "id")

// 		res.Write([]byte(userId))
// 	})

// 	route.Get("/", func(res http.ResponseWriter, req *http.Request) {
// 		product := req.URL.Query().Get("product")
// 		id := req.URL.Query().Get("id")

// 		if product == "" {
// 			res.Write([]byte("none!"))
// 		} else {
// 			res.Write([]byte(product + id))
// 		}
// 	})

// 	route.Get("/json", func(res http.ResponseWriter, req *http.Request) {
// 		res.Header().Set("Content-Type", "application/json")

// 		name := "Venzel"
// 		email := "tiago@gmail.com"
// 		password := "x012ds"

// 		account := domain.Account{}
// 		account.NewAccount(name, email, password)

// 		obj := map[string]interface{}{
// 			"name":     name,
// 			"email":    email,
// 			"password": password,
// 		}

// 		render.JSON(res, req, obj)
// 	})

// 	route.Post("/json", func(res http.ResponseWriter, req *http.Request) {
// 		res.Header().Set("Content-Type", "application/json")

// 		var account domain.Account

// 		render.DecodeJSON(req.Body, &account)

// 		account.Name = "Marcos"

// 		render.JSON(res, req, account)
// 	})

// 	port := ":3000"
// 	http.ListenAndServe(port, route)
// }