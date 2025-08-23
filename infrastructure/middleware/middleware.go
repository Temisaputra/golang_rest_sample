package middleware

import (
	"log"
	"net/http"

	"github.com/Temisaputra/warOnk/config"
	"github.com/Temisaputra/warOnk/pkg/auth"
	"github.com/Temisaputra/warOnk/pkg/helper"
)

type authMiddleware struct {
	cfg config.Config
}

func NewMiddleware(cfg config.Config) *authMiddleware {
	return &authMiddleware{cfg: cfg}
}

func (a *authMiddleware) Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path, r.URL.Scheme)
		user, err := auth.ValidateCurrentUser(a.cfg, r)
		if err != nil {
			helper.WriteResponse(w, err, nil)
			return
		}

		r = auth.SetUserContext(r, user)
		next.ServeHTTP(w, r)
	})
}
