package handler

import (
	"encoding/json"
	"net/http"

	"stakeholders.xws.com/model"
	"stakeholders.xws.com/service"
)

type AppRatingHandler struct {
	RatingService *service.AppRatingService
}

func (handler *AppRatingHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	ratings, err := handler.RatingService.GetAll()
	if err != nil {
		http.Error(writer, "Failed to fetch ratings", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(ratings)
}

func (handler *AppRatingHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var rating model.AppRating
	if err := json.NewDecoder(req.Body).Decode(&rating); err != nil {
		http.Error(writer, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Check if the user has already rated
	hasRated, err := handler.RatingService.HasUserRated(rating.UserId)
	if err != nil {
		http.Error(writer, "Failed to check user rating", http.StatusInternalServerError)
		return
	}
	if hasRated {
		http.Error(writer, "User has already rated", http.StatusBadRequest)
		return
	}

	// Create the rating
	createdRating, err := handler.RatingService.Create(&rating)
	if err != nil {
		http.Error(writer, "Failed to create rating", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(createdRating)
}
