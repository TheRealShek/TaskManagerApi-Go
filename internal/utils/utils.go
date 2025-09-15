package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

/*
In this Utils we will be reading JSON body, returning JSON responses, handling errors, and parsing URL params.
*/

// ReadJSONBody() will convert the request from Struct type to JSON format so that it is human readable
// We use destination interface{} as interface can have any type (struct, map, slice, etc.)
// any is an alias for interface{}
func ReadJSONBody(r *http.Request, destination any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	// Here we call Unmarshal(data []byte, v any) -> data(to be Unmarshaled) and location where result needs to be stored
	return json.Unmarshal(body, destination) // This works cause Unmarshal Func returns a error
}

// SendJSONResponse() will convert the request from Struct type to JSON format so that it is consistent
// status is the protocol-level signal (200 OK, 400 Bad Request .. )
// success is a application-level signal inside your JSON body (for self verification)
func SendJSONResponse(w http.ResponseWriter, status int, success bool, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	resp := map[string]any{
		"success": success,
		"data":    data,
	}
	json.NewEncoder(w).Encode(resp) //encode the data from resp and write to w
}

// SendErrorResponse() is the error handling if SendJSONResponse() does not work
func SendErrorResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	resp := map[string]interface{}{
		"success": false,
		"error":   message,
	}
	json.NewEncoder(w).Encode(resp)
}

// ParseURLParams() is used to extract the parameters from the URL
// Trim-> "/users/123/tasks" → "users/123/tasks"
// Split-> ["users", "123", "tasks"]
func ParseURLParams(r *http.Request) []string {
	path := strings.Trim(r.URL.Path, "/")
	return strings.Split(path, "/")
}
