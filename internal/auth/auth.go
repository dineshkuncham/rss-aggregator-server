package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts the API key from the
// Http Request Header
// Ex: Authorization: Bearer {api key}
func FetchApiKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no authorization header found")
	}
	authHeaderValues := strings.Split(authHeader, " ")
	if len(authHeaderValues) != 2 || authHeaderValues[0] != "Bearer" {
		return "", errors.New("authorization header is malfunctioned")
	}
	return authHeaderValues[1], nil
}
