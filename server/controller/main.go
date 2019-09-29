package controller

import (
	"net/http"

	"github.com/techbot/server/config"
)

type Endpoint struct {
	Path    string
	Execute func(w http.ResponseWriter, r *http.Request)
}

var Endpoints = map[string]*Endpoint{
	getEndpointKey(postAnswer): postAnswer,
	getEndpointKey(sendAnswer): sendAnswer,
}

func getEndpointKey(endpoint *Endpoint) string {
	return endpoint.Path
}

// Authenticated verifies if provided request is performed by a logged-in Mattermost user.
func Authenticated(w http.ResponseWriter, r *http.Request) bool {
	userID := r.Header.Get(config.HeaderMattermostUserID)
	if userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}

	return true
}
