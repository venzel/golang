package routes

import "github.com/gorilla/mux"

func GetRoutes() *mux.Router {
	return mux.NewRouter()
}
