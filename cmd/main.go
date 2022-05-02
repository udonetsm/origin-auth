package main

import (
	"net/http"
	"origin-auth/auth"
	"origin-auth/getconf"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/authorize", auth.Mdlwr(http.HandlerFunc(auth.Authorize))).Methods(http.MethodPost)
	router.Handle("/newuser", auth.Mdlwr(http.HandlerFunc(auth.NewUser))).Methods(http.MethodPost)
	return router
}

func Server(router *mux.Router) *http.Server {
	return &http.Server{
		Handler: router,
		Addr:    getconf.Server.Addr,
	}
}

func main() {
	Server(Router()).ListenAndServe()
}
