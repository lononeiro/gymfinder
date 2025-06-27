package middleware

import (
	"net/http"
	"strings"

	"github.com/lononeiro/gymfinder/backend/src/utils"
)

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Token ausente ou mal formatado", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		if isAdmin, ok := claims["is_admin"].(bool); !ok || !isAdmin {
			http.Error(w, "Acesso restrito a administradores", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
