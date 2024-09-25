package middleware

import (
	"encoding/json"
	"fmt"
	"http-basic/model/web"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type AuthMiddleware struct {
	handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{handler: handler}
}

func (a *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		sendUnauthorizedResponse(w, "Authorization header missing")
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		sendUnauthorizedResponse(w, "Invalid token format")
		return
	}

	userToken := strings.TrimPrefix(authHeader, "Bearer ")

	secretKey := []byte("this is very secret")

	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		fmt.Println("Error parsing token:", err)
		sendUnauthorizedResponse(w, "Invalid or expired token")
		return
	}
	a.handler.ServeHTTP(w, r)
}

func sendUnauthorizedResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	webResponse := web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
		Data:   message,
	}
	json.NewEncoder(w).Encode(webResponse)
}
