package handlers

import (
	"net/http"
	"taskmanagerapi/internal/utils"
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
	if r.Method != "POST" {
		utils.SendErrorResponse(w, http.StatusMethodNotAllowed, "POST Method only allowed")
		return
	}
	var request SignupRequest
	err := utils.ReadJSONBody(r, &request)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid JSON Body")
		return
	}

	// Sending confirmation to the User
	data := map[string]string{"message": "User signed up", "username": request.Username}
	utils.SendJSONResponse(w, http.StatusOK, true, data)
}

// Handles Create Task
func HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.SendErrorResponse(w, http.StatusMethodNotAllowed, "POST Method only allowed")
		return
	}
	params := utils.ParseURLParams(r)
	taskID := ""
	if len(params) > 2 {
		taskID = params[2]
	}

	data := map[string]string{"message": "Task created", "task_id": taskID}
	utils.SendJSONResponse(w, http.StatusOK, true, data)
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
