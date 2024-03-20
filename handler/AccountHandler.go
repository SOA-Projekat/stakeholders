package handler

import (
	"encoding/json"
	"net/http"

	"stakeholders.xws.com/model"
	"stakeholders.xws.com/service"
)

type AccountHandler struct {
	AccountService *service.AccountService
}

func (handler *AccountHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	accounts, err := handler.AccountService.GetAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(accounts)
}

func (handler *AccountHandler) BlockOrUnblock(writer http.ResponseWriter, req *http.Request) {
	var account model.Account
	if err := json.NewDecoder(req.Body).Decode(&account); err != nil {
		http.Error(writer, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	updatedAccount, err := handler.AccountService.BlockOrUnblock(&account)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(updatedAccount)
}
