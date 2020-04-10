package middleware

import (
	"context"
	"net/http"
)

func AddJwtContext(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		jwt := r.Header.Get("JWT")
		ctx := context.WithValue(r.Context(), "jwt", jwt)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
