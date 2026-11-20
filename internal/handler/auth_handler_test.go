package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joddyhartono/teguh-jaya/internal/model"
)

type mockUserRepo struct {
	GetByUsernameFn func(ctx context.Context, username string) (*model.User, error)
}

func (m *mockUserRepo) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	return m.GetByUsernameFn(ctx, username)
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