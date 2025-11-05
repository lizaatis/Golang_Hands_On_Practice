package main

import (
	"encoding/json"
	"net/http"
)

func handleClientProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetClientProfile(w, r)
	case http.MethodPost:
		UpdateClientProfile(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetClientProfile(w http.ResponseWriter, r *http.Request) {
	clientProfile := r.Context().Value("clientProfile").(ClientProfile)
	response := ClientProfile{
		Email: clientProfile.Email,
		Name:  clientProfile.Name,
		Id:    clientProfile.Id,
	}
	w.Header().Set("Content-Type", "application/json")

	//http.ResponseWriter(w,"StatusOk", http.StatusO
	json.NewEncoder(w).Encode(response)

}
func UpdateClientProfile(w http.ResponseWriter, r *http.Request) {
	clientProfile := r.Context().Value("clientProfile").(ClientProfile)
	var payloadData ClientProfile
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
		// var clientId = r.URL.Query().Get("clientId")
		// clientProfile, ok := database[clientId]
		// if !ok || clientId == "" {
		// 	http.Error(w, "Forbidden", http.StatusForbidden)
	}
	defer r.Body.Close()
	//Update the profile
	clientProfile.Email = payloadData.Email
	clientProfile.Name = payloadData.Name
	clientProfile.Id = payloadData.Id

	w.WriteHeader(http.StatusOK)
}
