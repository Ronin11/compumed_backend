package auth

import (
	"net/http"
	"strings"
)

type AuthenticatedHandler func(http.ResponseWriter, *http.Request, *User)

type EnsureAuth struct {
	handler AuthenticatedHandler
}

func getTokenFromRequest(r *http.Request) string {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) > 1 {
		return splitToken[1]
	}
	return ""
}

func (ea *EnsureAuth) Handle(w http.ResponseWriter, r *http.Request) {
	token := getTokenFromRequest(r)
	authHandler := GetAuthHandlerInstance()
	user := authHandler.GetUserFromToken(token)
	if user == nil {
		http.Error(w, "please sign-in", http.StatusUnauthorized)
		return
	}

	ea.handler(w, r, user)
}

func BuildRouteWithUser(handlerToWrap AuthenticatedHandler) *EnsureAuth {
	return &EnsureAuth{handlerToWrap}
}
