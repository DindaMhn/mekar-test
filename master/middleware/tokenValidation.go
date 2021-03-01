package middleware

import (
	"fmt"
	"mekar/master/tools/jwt"
	"net/http"
	"strings"
)

//TokenValidationMiddleware app
func TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if len(token) == 0 {
			http.Error(w, "You are not allowed to access this service", http.StatusUnauthorized)
		} else {
			tokenVal := strings.Split(token, "Bearer ")
			_, err := jwt.JwtDecoder(tokenVal[1])
			if err != nil {
				http.Error(w, "You are not allowed to access this service", http.StatusUnauthorized)
				fmt.Println(err)
			} else {
				next.ServeHTTP(w, r)
			}
		}
	})
}
