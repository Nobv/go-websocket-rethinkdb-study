package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	r "gopkg.in/gorethink/gorethink.v4"
)

var (
	router  *mux.Router
	session *r.Session
)

func init() {
	var err error

	session, err = r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "chat",
		MaxOpen:  40,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	messages := []Message{}

	res, err := r.Table("messages").OrderBy(r.Asc("Created")).Run(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = res.All(&messages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderTemplate(w, "index", map[string]interface{}{
		"Messages": messages,
		"Route":    "all",
	})

}

func newHandler(w http.ResponseWriter, req *http.Request) {
	message := NewMassage(req.PostFormValue("message"))
	message.Created = time.Now()

	_, err := r.Table("messages").Insert(message).RunWrite(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, "/", http.StatusFound)
}

func newChangesHandler(fn func(chan interface{})) http.HandlerFunc {
	o := newObserver()
	go o.start()

	fn(o.broadcast)

	return wsHandler(o)
}
