// Middleware for authentication in
package main

import (
	"context"
	"net/http"
	"strings"
)

func TokenAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var clientId = r.URL.Query().Get("clientId")
		clientProfile, ok := database[clientId]
		if !ok || clientId == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		token := r.Header.Get("Authorization")
		if !inValidToken(clientProfile, token) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		ctx := context.WithValue(r.Context(), "clientProfile", clientProfile)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
}

// Checks if token matches
func inValidToken(clientProfile ClientProfile, token string) bool {
	if strings.HasPrefix(token, "Bearer ") {
		return strings.TrimPrefix(token, "Bearer ") == clientProfile.Token
	}
	return false
}
