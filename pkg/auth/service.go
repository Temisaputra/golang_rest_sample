package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Temisaputra/warOnk/config"
	"github.com/Temisaputra/warOnk/pkg/helper"
)

type UserContextKey string
type JWTContextKey string

const (
	UserContext UserContextKey = "USER_CONTEXT_KEY"
	JWTContext  JWTContextKey  = "USER_CONTEXT_KEY"
)

func ValidateCurrentUser(cfg config.Config, r *http.Request) (*User, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return nil, helper.NewErrUnauthorized("token not found")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/auth/current-user", nil), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response GetCurrentUserResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, helper.NewErrUnauthorized(response.Message)
	}

	return &response.Data, nil
}

func SetUserContext(r *http.Request, user *User) *http.Request {
	ctx := r.Context()
	ctx = context.WithValue(ctx, UserContext, user)
	ctx = context.WithValue(ctx, JWTContext, r.Header.Get("Authorization"))
	return r.WithContext(ctx)
}

func GetUserContext(ctx context.Context) *User {
	return ctx.Value(UserContext).(*User)
}

func GetJWTContext(ctx context.Context) string {
	return ctx.Value(JWTContext).(string)
}
