package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/joddyhartono/teguh-jaya/internal/model"
	"github.com/joddyhartono/teguh-jaya/internal/response"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (*model.User, error)
}

type AuthHandler struct {
	repo 		UserRepository
	jwtSecret 	string

}

func NewAuthHandler(repo UserRepository, jwtSecret string) *AuthHandler {
	return &AuthHandler{
		repo: repo,
		jwtSecret: jwtSecret,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WriteJSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Username == "" || req.Password == "" {
		response.WriteJSONError(w, http.StatusBadRequest, "username and password are required")
		return
	}

	user, err := h.repo.GetByUsername(r.Context(), req.Username)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			response.WriteJSONError(w, http.StatusUnauthorized, "invalid credentials")
			return
		}
		log.Printf("failed to get user: %v", err)
		response.WriteJSONError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		response.WriteJSONError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	token, err := Generate(user.ID, h.jwtSecret)
	if err != nil {
		log.Printf("failed to generate token: %v", err)
		response.WriteJSONError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	response.WriteJSON(w, http.StatusOK, loginResponse{Token: token})
}