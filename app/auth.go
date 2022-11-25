package app

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"simple-api/models"
	"simple-api/responses"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var userRole string

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/api/user/new", "/api/user/login"} //List of endpoints that doesn't require auth
		requestPath := r.URL.Path                               //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(rw, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			rw.WriteHeader(http.StatusForbidden)
			response := responses.BaseResponse{Status: http.StatusBadRequest, Message: "Missing auth token", Data: map[string]interface{}{}}
			rw.Header().Add("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(response)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			rw.WriteHeader(http.StatusForbidden)
			response := responses.BaseResponse{Status: http.StatusBadRequest, Message: "Invalid/Malformed auth token", Data: map[string]interface{}{}}
			rw.Header().Add("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(response)
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_secret_key")), nil
		})

		userRole = tk.Role

		if err != nil { //Malformed token, returns with http code 403 as usual
			rw.WriteHeader(http.StatusForbidden)
			response := responses.BaseResponse{Status: http.StatusBadRequest, Message: "Malformed authentication token", Data: map[string]interface{}{}}
			rw.Header().Add("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(response)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			rw.WriteHeader(http.StatusForbidden)
			response := responses.BaseResponse{Status: http.StatusBadRequest, Message: "Token is not valid", Data: map[string]interface{}{}}
			rw.Header().Add("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(response)
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		ctx := context.WithValue(r.Context(), "user", tk.Id)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r) //proceed in the middleware chain!
	})
}
