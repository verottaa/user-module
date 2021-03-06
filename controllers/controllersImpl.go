package controllers

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/verottaa/user-module/entity"
	"github.com/verottaa/user-module/helpers"
	"github.com/verottaa/user-module/service"
	"net/http"
)

type controller struct {
	service        service.Service
	variableReader helpers.VariableReader
	jsonWorker helpers.JsonWorker
}

func CreateController() Controller {
	var controllerInst = new(controller)
	controllerInst.service = service.GetService()
	controllerInst.variableReader = helpers.GetVariableReader()
	controllerInst.jsonWorker = helpers.GetJsonWorker()
	return controllerInst
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
	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	err := c.jsonWorker.Decode(r, user)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	id, err := c.service.Store(&user)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
		response := entity.ObjectCreatedDto{Id: id.String()}
		err = c.jsonWorker.Encode(w, response)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusCreated)
		}
	}
}

func (c controller) updateOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := c.variableReader.GetObjectIdFromQueryByName(r, "id")
	var user entity.User
	err := c.jsonWorker.Decode(r, user)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if id.String() != user.Id.String() {
		var err = errors.New("validation didn't pass because user.id and /:id not equal")
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.service.Update(&user)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (c controller) deleteOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := c.variableReader.GetObjectIdFromQueryByName(r, "id")
	err := c.service.Delete(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (c controller) deleteAll(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := c.service.DeleteAll()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (c controller) deleteMany(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (c controller) getOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := c.variableReader.GetObjectIdFromQueryByName(r, "id")
	user, err := c.service.Find(id)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	err = c.jsonWorker.Encode(w, user)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (c controller) getAll(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := c.service.FindAll()
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusNotFound)
	}

	err = c.jsonWorker.Encode(w, users)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusOK)
	}
}
