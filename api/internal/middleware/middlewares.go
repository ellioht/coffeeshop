package middleware

import "net/http"

type IAuthMiddleware interface {
	AuthMiddleware(next http.Handler) http.Handler
}

type Middlewares struct {
	Auth IAuthMiddleware
}
