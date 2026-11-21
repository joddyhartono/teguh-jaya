package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/joddyhartono/teguh-jaya/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type mockUserRepo struct {
	GetByUsernameFn func(ctx context.Context, username string) (*model.User, error)
}

func (m *mockUserRepo) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	return m.GetByUsernameFn(ctx, username)
}

func Test_Login_InvalidBody(t *testing.T) {
	// Arrange
	handler := &AuthHandler{}
	
	req := httptest.NewRequest(http.MethodPost, "/api/auth/login", strings.NewReader("bukan json"))
	rec := httptest.NewRecorder()

	// Act
	handler.Login(rec, req)
	
	// Assert
	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func Test_Login_EmptyFields(t *testing.T) {
	// Arrange
	handler := &AuthHandler{}
	loginReq := loginRequest{
		Username: "",
		Password: "",
	}

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(loginReq)
	if err != nil {
		t.Fatalf("failed to encode login request: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/auth/login", &body)

	rec := httptest.NewRecorder()

	// Act
	handler.Login(rec, req)

	// Assert
	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func Test_Login_UserNotFound(t *testing.T) {
	// Arrange
	mockRepo := &mockUserRepo{
		GetByUsernameFn: func(ctx context.Context, username string) (*model.User, error) {
			return nil, model.ErrNotFound
		},
	}

	handler := &AuthHandler{
		repo: mockRepo,
		jwtSecret: "jwt-secret",
	}

	loginReq := loginRequest{
		Username: "admin",
		Password: "admin",
	}

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(loginReq)
	if err != nil {
		t.Fatalf("failed to encode login request: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/auth/login", &body)

	rec := httptest.NewRecorder()

	// Act
	handler.Login(rec, req)
	
	// Assert
	if rec.Code != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusUnauthorized, rec.Code)
	}
}

func Test_Login_RepoError(t *testing.T) {
	// Arrange
	mockRepo := &mockUserRepo{
		GetByUsernameFn: func(ctx context.Context, username string) (*model.User, error) {
			return nil, errors.New("db down")
		},
	}

	handler := &AuthHandler{
		repo: mockRepo,
		jwtSecret: "jwt-secret",
	}

	loginReq := loginRequest{
		Username: "admin",
		Password: "admin",
	}

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(loginReq)
	if err != nil {
		t.Fatalf("failed to encode login request: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/auth/login", &body)

	rec := httptest.NewRecorder()

	// Act
	handler.Login(rec, req)
	
	// Assert
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, rec.Code)
	}
}

func Test_Login_WrongPassword(t *testing.T) {
	// Arrange
	hash, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.MinCost)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	mockRepo := &mockUserRepo{
		GetByUsernameFn: func(ctx context.Context, username string) (*model.User, error) {
			return &model.User{
				Username: "admin",
				PasswordHash: string(hash),
			}, nil
		},
	}

	handler := &AuthHandler{
		repo: mockRepo,
		jwtSecret: "jwt-secret",
	}

	loginReq := loginRequest{
		Username: "admin",
		Password: "salah",
	}

	var body bytes.Buffer
	err = json.NewEncoder(&body).Encode(loginReq)
	if err != nil {
		t.Fatalf("failed to encode login request: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/auth/login", &body)

	rec := httptest.NewRecorder()

	// Act
	handler.Login(rec, req)
	
	// Assert
	if rec.Code != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusUnauthorized, rec.Code)
	}
}

func Test_Login_Success(t *testing.T) {
	// Arrange
	hash, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.MinCost)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	mockRepo := &mockUserRepo{
		GetByUsernameFn: func(ctx context.Context, username string) (*model.User, error) {
			return &model.User{
				Username: "admin",
				PasswordHash: string(hash),
			}, nil
		},
	}

	handler := &AuthHandler{
		repo: mockRepo,
		jwtSecret: "jwt-secret",
	}

	loginReq := loginRequest{
		Username: "admin",
		Password: "admin",
	}

	var body bytes.Buffer
	err = json.NewEncoder(&body).Encode(loginReq)
	if err != nil {
		t.Fatalf("failed to encode login request: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/auth/login", &body)

	rec := httptest.NewRecorder()

	// Act
	handler.Login(rec, req)
	
	// Assert
	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, rec.Code)
	}
}