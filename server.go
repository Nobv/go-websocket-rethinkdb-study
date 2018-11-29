package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewServer(addr string) *http.Server {

	router = initRouting()

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}

func StartServer(server *http.Server) {
	log.Println("stating server....")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func initRouting() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/all", indexHandler)
	r.HandleFunc("/new", newHandler)

	r.Handle("/ws/all", newChangesHandler(allChanges))

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	return r
}
