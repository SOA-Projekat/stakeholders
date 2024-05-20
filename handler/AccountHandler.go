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

func (handler *AccountHandler) Register(writer http.ResponseWriter, req *http.Request) {
	var user model.User

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(writer, "Failed to parse request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	token, err := handler.AccountService.CreateUser(&user)
	if err != nil {
		http.Error(writer, "Error creating user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(token)
}

func (handler *AccountHandler) Login(w http.ResponseWriter, req *http.Request) {
	var credentials model.Credentials
	if err := json.NewDecoder(req.Body).Decode(&credentials); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	token, err := handler.AccountService.Login(&credentials)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Set authentication token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token.AccessToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
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
