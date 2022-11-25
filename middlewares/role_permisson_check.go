package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"simple-api/responses"
)

func contains(array []string, find string) bool {
	for _, item := range array {
		if item == find {
			return true
		}
	}
	return false
}

var RolePermissionCheck = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		isPermitted := false

		permittedEndpoints := map[string]map[string][]string{
			"user": {
				"GET": []string{"/token", "/posts"},
			},
			"admin": {
				"GET":    []string{"/token", "/posts"},
				"POST":   []string{"/users", "/posts"},
				"PUT":    []string{"/posts"},
				"DELETE": []string{"/posts"},
			},
		} //List of endpoints permitted to each role

		requestPath := r.URL.Path //current request path

		if contains(permittedEndpoints[UserRole][r.Method], requestPath) {
			isPermitted = true
		}

		if !isPermitted {
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusForbidden)
			response := responses.BaseResponse{Status: http.StatusForbidden, Message: "You are not permitted to access this endpoint", Data: map[string]interface{}{}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		ctx := context.WithValue(r.Context(), "is-role-permitted", true)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r) //proceed in the middleware chain!
	})
}
