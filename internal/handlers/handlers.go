package handlers

import (
	"fmt"
	"net/http"
)

// NewRouter() gives us a pointer to a Router Struct which is mapped to a HandleFunc
func NewRouter() *Router {
	return &Router{
		routes: make(map[string]HandlerFunc),
	}
}

/*
So if you call:

	router.Handle("GET", "/tasks", getTasksHandler)

It stores it in the map named routes:

	"GET-/tasks" → getTasksHandler
*/
// So this method is to add the keys for the predefined routes
func (r *Router) Handle(method, path string, handler HandlerFunc) {
	key := method + "-" + path
	r.routes[key] = handler
}

// Handles Signup
func HandleSignup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Signup handler called!")
}

// Handles Create Task
func HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create Task handler called!")
}

// This method is to verify if the route exists and only then will it redirect
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, exists := r.routes[key]; exists {
		handler(w, req)
	} else {
		http.NotFound(w, req)
	}
}
