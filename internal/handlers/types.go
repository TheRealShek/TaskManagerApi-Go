package handlers

import "net/http"

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Router struct {
	routes map[string]HandlerFunc
}
