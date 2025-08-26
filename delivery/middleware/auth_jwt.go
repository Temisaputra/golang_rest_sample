package middleware

import (
	"net/http"

	"github.com/Temisaputra/warOnk/pkg/auth"
	"github.com/Temisaputra/warOnk/pkg/helper"
)

type AuthMiddleware struct {
	jwtSvc auth.JwtService
}

func NewAuthMiddleware(jwtSvc auth.JwtService) *AuthMiddleware {
	return &AuthMiddleware{jwtSvc: jwtSvc}
}

func (a *AuthMiddleware) Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user, err := a.jwtSvc.ValidateCurrentUser(r)
		if err != nil {
			switch err {
			case auth.ErrTokenExpired:
				helper.WriteResponse(w, auth.ErrTokenExpired, nil)
			case auth.ErrTokenInvalid:
				helper.WriteResponse(w, auth.ErrTokenInvalid, nil)
			case auth.ErrUnauthorized:
				helper.WriteResponse(w, auth.ErrUnauthorized, nil)
			default:
				helper.WriteResponse(w, auth.ErrMissingToken, nil)
			}
			return
		}

		// simpan user ke context
		r = auth.SetUserContext(r, user)
		next.ServeHTTP(w, r)
	})
}
