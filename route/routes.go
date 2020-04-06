package route

import (
	"github.com/gorilla/mux"
)

func registerRoutes(router *mux.Router) {
	a := AuthRoute{AuthRouter: router.PathPrefix("/auth").Subrouter()}
	a.InitRoutes()
}



func GetRouter() *mux.Router {
	router := mux.NewRouter()
	registerRoutes(router)
	return router
}
