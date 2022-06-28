package main

import (
	"context"
	"log"
	"net/http"
)

type contextKey int

const authenticatedUserKey contextKey = 0

type User struct {
	name string
}

type Logger struct {
	handler http.Handler
}

func (l Logger) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	user := request.Context().Value(authenticatedUserKey).(*User)
	log.Printf("The user is %v\n", user)
}

func NewLogger(handler http.Handler) *Logger {
	return &Logger{handler: handler}
}

type EnsureAuth struct {
	handler http.Handler
}

func (ea *EnsureAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, err := GetAuthenticatedUser(r)
	if err != nil {
		http.Error(w, "please sign-in", http.StatusUnauthorized)
		return
	}

	ctxWithUser := context.WithValue(r.Context(), authenticatedUserKey, user)
	rWithUser := r.WithContext(ctxWithUser)
	ea.handler.ServeHTTP(w, rWithUser)
}

func NewEnsureAuth(handlerToWrap http.Handler) *EnsureAuth {
	return &EnsureAuth{handlerToWrap}
}

func GetAuthenticatedUser(r *http.Request) (*User, error) {
	//validate the session token in the request,
	//fetch the session state from the session store,
	//and return the authenticated user
	//or an error if the user is not authenticated
	return &User{name: "John"}, nil
}

func main() {
	mux := http.NewServeMux()
	wrapped := NewEnsureAuth(NewLogger(mux))

	mux.Handle("/user", wrapped)
	log.Fatal(http.ListenAndServe(":8080", wrapped))
}
