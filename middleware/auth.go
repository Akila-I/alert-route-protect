// File: middleware/auth.go
package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type GithubUser struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
}

func ValidateGithubToken(clientID, clientSecret string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")

			// Validate token with GitHub
			req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
			if err != nil {
				http.Error(w, "Server error", http.StatusInternalServerError)
				return
			}

			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			req.Header.Set("Accept", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				http.Error(w, "Failed to validate token", http.StatusUnauthorized)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				http.Error(w, "Failed to read response", http.StatusInternalServerError)
				return
			}

			var user GithubUser
			if err := json.Unmarshal(body, &user); err != nil {
				http.Error(w, "Failed to parse user data", http.StatusInternalServerError)
				return
			}

			// Store user info in context
			ctx := context.WithValue(r.Context(), "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}
