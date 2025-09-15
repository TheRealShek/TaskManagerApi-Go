package handlers

import "net/http"

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Router struct {
	routes map[string]HandlerFunc
}

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
