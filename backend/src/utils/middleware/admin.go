package middleware

import (
	"net/http"

	"github.com/lononeiro/gymfinder/backend/src/utils"
)

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value("userClaims").(*utils.Claims)
		if !ok || claims == nil {
			http.Error(w, "Token inválido ou não encontrado", http.StatusUnauthorized)
			return
		}

		if !claims.IsAdmin {
			http.Error(w, "Acesso restrito a administradores", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
