package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
)

type controller struct {
}

func (c controller) InitController(router *mux.Router) {
	router.HandleFunc("/", c.createOne).Methods("POST")
	router.HandleFunc("/", c.getAll).Methods("GET")
	router.HandleFunc("/", c.deleteAll).Methods("DELETE")
	router.HandleFunc("/many", c.deleteMany).Methods("DELETE")
	router.HandleFunc("/{id}", c.getOne).Methods("GET")
	router.HandleFunc("/{id}", c.updateOne).Methods("PUT")
	router.HandleFunc("/{id}", c.deleteOne).Methods("DELETE")
}

func (c controller) createOne(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (c controller) updateOne(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (c controller) deleteOne(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (c controller) deleteAll(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (c controller) deleteMany(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (c controller) getOne(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (c controller) getAll(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
