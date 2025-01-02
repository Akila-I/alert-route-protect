// File: handlers/deploy.go
package handlers

import (
	"encoding/json"
	"net/http"
)

type DeployAlert struct {
	Message string `json:"message"`
}

func TriggerAutoBuildAlert(w http.ResponseWriter, r *http.Request) {
	var alert DeployAlert
	if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response := map[string]string{
		"status":  "success",
		"message": "Deploy alert received: " + alert.Message,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
