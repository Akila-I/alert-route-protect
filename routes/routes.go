// File: routes/routes.go
package routes

import (
	"akila-i/github-oauth-go/config"
	"akila-i/github-oauth-go/handlers"
	"akila-i/github-oauth-go/middleware"
	"net/http"
)

type Route struct {
	Path        string
	Methods     []string
	HandlerFunc http.HandlerFunc
}

func SetupRoutes(cfg *config.Config) http.Handler {
	mux := http.NewServeMux()

	routes := []Route{
		{
			Path:        "/image/deploy-alert",
			Methods:     []string{http.MethodPost},
			HandlerFunc: middleware.ValidateGithubToken(cfg.ClientID, cfg.ClientSecret)(handlers.TriggerAutoBuildAlert),
		},
	}

	// Register routes
	for _, route := range routes {
		handler := route.HandlerFunc
		mux.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			// Method validation
			methodAllowed := false
			for _, method := range route.Methods {
				if r.Method == method {
					methodAllowed = true
					break
				}
			}

			if !methodAllowed {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			handler(w, r)
		})
	}

	return mux
}
