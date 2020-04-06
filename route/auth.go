package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

type AuthRoute struct {
	AuthRouter *mux.Router
}

func (ar *AuthRoute) InitRoutes() {

	ar.AuthRouter.HandleFunc("/login", HandleLogin)
	ar.AuthRouter.HandleFunc("/register", HandleRegister)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	bytes := []byte("hello world")
	w.Write(bytes)
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {

}
