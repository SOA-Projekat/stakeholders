package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"stakeholders.xws.com/model"
	"stakeholders.xws.com/service"
)

type PersonHandler struct {
	PersonService *service.PersonService
}

func (handler *PersonHandler) Get(writer http.ResponseWriter, req *http.Request) {
	userID := mux.Vars(req)["userId"]
	log.Printf("Getting person with userID: %s", userID)

	// Convert userID to integer
	var userIDInt int
	_, err := fmt.Sscanf(userID, "%d", &userIDInt)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	person, err := handler.PersonService.FindPerson(userIDInt)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(person)
}

func (handler *PersonHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var obj model.Person
	err := json.NewDecoder(req.Body).Decode(&obj)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedPerson, err := handler.PersonService.UpdatePerson(&obj)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(updatedPerson)
}
